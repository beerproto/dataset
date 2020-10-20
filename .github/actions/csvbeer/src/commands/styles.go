package commands

import (
	"encoding/json"
	"fmt"
	beerproto "github.com/beerproto/beerproto_go"
	"github.com/beerproto/dataset/src/csv"
	"github.com/gocarina/gocsv"
	"github.com/golang/protobuf/jsonpb"
	"io"
	"os"
)

func ParseStyle(stylePath string, output Output, file string) {
	fmt.Println(fmt.Sprintf("Step 1/1 : Loading %s", stylePath))

	stylesFile, err := loadFile(stylePath)
	if err != nil {
		panic(fmt.Errorf("failed to load file %v: %w", stylePath, err))
	}
	defer stylesFile.Close()

	var styles []*csv.Style

	if err := gocsv.UnmarshalFile(stylesFile, &styles); err != nil {
		panic(fmt.Errorf("%v does not match Style %v format: %w", stylePath, CsvExt, err))
	}

	fmt.Println("Successfully parased")


	var arr []*beerproto.StyleType
	for _, style := range styles {
		arr = append(arr, style.ToStyleType())
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
		if file != "" {
			f, err := os.Create(file)
			if err != nil {
				panic(fmt.Errorf("failed to create file %s: %w", file, err))
			}
			fmt.Println(fmt.Sprintf("Makine file %s", file))
			defer f.Close()
			w = f

		} else {
			jsonPath := stylePath[0:len(stylePath)-len(CsvExt)] + JsonExt
			f, err := os.Create(jsonPath)
			if err != nil {
				panic(fmt.Errorf("failed to create file %s: %w", jsonPath, err))
			}
			defer f.Close()
			w = f
		}
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
