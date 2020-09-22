package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	beerproto "github.com/beerproto/beerproto_go"
	"github.com/beerproto/dataset/src/csv"
	"github.com/gocarina/gocsv"
	"github.com/golang/protobuf/jsonpb"
)

type Output string

const (
	CsvExt  = ".csv"
	JsonExt = ".json"

	TTY  = Output("tty")
	FILE = Output("file")
)

func ParseStyle(stylePath, indexPath string, output Output) {

	fmt.Println(fmt.Sprintf("Step 1/2 : Loading %s", indexPath))

	indexFile, err := loadFile(indexPath)
	if err != nil {
		panic(err)
	}
	defer indexFile.Close()

	var indexes []*csv.Index

	if err := gocsv.UnmarshalFile(indexFile, &indexes); err != nil {
		panic(fmt.Errorf("%v does not match Index %v format: %w", indexPath, CsvExt, err))
	}

	fmt.Println(fmt.Sprintf("Step 2/2 : Loading %s", stylePath))

	stylesFile, err := loadFile(stylePath)
	if err != nil {
		panic(err)
	}
	defer stylesFile.Close()

	var styles []*csv.Style

	if err := gocsv.UnmarshalFile(stylesFile, &styles); err != nil {
		panic(fmt.Errorf("%v does not match Style %v format: %w", stylePath, CsvExt, err))
	}

	fmt.Println("Successfully parased")

	indexMap := indexToMap(indexes)

	var arr []*beerproto.StyleType
	for _, style := range styles {
		category := indexMap[style.StyleID]
		arr = append(arr, style.ToStyleType(category))
	}

	recipe := &beerproto.Recipe{
		Styles: arr,
	}

	var w io.Writer

	switch output {
	case TTY:
		w = os.Stdout
		break
	case FILE:
		jsonPath := stylePath[0:len(stylePath)-len(CsvExt)] + JsonExt
		f, err := os.Create(jsonPath)
		if err != nil {
			panic(fmt.Errorf("failed to create file %s: %w", jsonPath, err))
		}
		defer f.Close()
		w = f
	}

	// First marshal through the protobuf jsonpb.Marshaler, standard encoding/json package when called on protobuf message types does not operate correctly.
	m := jsonpb.Marshaler{}
	j, err := m.MarshalToString(recipe)
	if err != nil {
		panic(fmt.Errorf("failed to marshal json: %w", err))
	}
	raw := map[string]json.RawMessage{}

	// Second marshal to get indented json
	json.Unmarshal([]byte(j), &raw)
	data, err := json.MarshalIndent(raw, "", "\t")
	if err != nil {
		panic(fmt.Errorf("failed to marshal json: %w", err))
	}

	_, err = w.Write(data)
	if err != nil {
		panic(fmt.Errorf("failed to write json: %w", err))
	}

	fmt.Println("\nDone")
}

func indexToMap(indexs []*csv.Index) map[int]string {
	indexMap := map[int]string{}
	for _, i := range indexs {
		indexMap[i.ID] = i.Style
	}

	return indexMap
}

func loadFile(filePath string) (*os.File, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		return nil, fmt.Errorf("file not found %v: %w", filePath, err)
	}
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("file not found %v", filePath)
	}
	ext := filepath.Ext(filePath)
	if ext != CsvExt {
		return nil, fmt.Errorf("%v not a %v file", filePath, CsvExt)
	}
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to open %v: %w", filePath, err)
	}
	return file, nil
}
