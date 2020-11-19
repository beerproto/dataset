package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	beerproto "github.com/beerproto/beerproto_go"
	"github.com/beerproto/dataset/src/csv"
	"github.com/gocarina/gocsv"
	"github.com/golang/protobuf/jsonpb"
)

func ParseCulture(cultureItemsPath string, output Output, file string) {
	fmt.Println(fmt.Sprintf("Step 1/1 : Loading %s", cultureItemsPath))

	cultureItemsFile, err := loadFile(cultureItemsPath)
	if err != nil {
		panic(fmt.Errorf("failed to load file %v: %w", cultureItemsPath, err))
	}
	defer cultureItemsFile.Close()

	var cultures []*csv.Culture

	if err := gocsv.UnmarshalFile(cultureItemsFile, &cultures); err != nil {
		panic(fmt.Errorf("%v does not match Culture %v format: %w", cultureItemsPath, CsvExt, err))
	}

	fmt.Println("Successfully parsed")

	var arr []*beerproto.CultureInformation
	for _, e := range cultures {
		arr = append(arr, e.ToCultureInformation())
	}

	recipe := &beerproto.Recipe{
		Cultures: arr,
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
			jsonPath := cultureItemsPath[0:len(cultureItemsPath)-len(CsvExt)] + JsonExt
			f, err := os.Create(jsonPath)
			if err != nil {
				panic(fmt.Errorf("failed to create file %s: %w", jsonPath, err))
			}
			fmt.Println(fmt.Sprintf("Making file %s", jsonPath))
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
