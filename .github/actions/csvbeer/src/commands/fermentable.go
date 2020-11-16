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

func ParseFermentables(maltItemsPath string, output Output, file string) {
	fmt.Println(fmt.Sprintf("Step 1/1 : Loading %s", maltItemsPath))

	maltsItemsFile, err := loadFile(maltItemsPath)
	if err != nil {
		panic(fmt.Errorf("failed to load file %v: %w", maltItemsPath, err))
	}
	defer maltsItemsFile.Close()

	var malts []*csv.Fermentable

	if err := gocsv.UnmarshalFile(maltsItemsFile, &malts); err != nil {
		panic(fmt.Errorf("%v does not match Fermentables %v format: %w", maltItemsPath, CsvExt, err))
	}

	fmt.Println("Successfully parased")

	var arr []*beerproto.FermentableType
	for _, e := range malts {
		arr = append(arr, e.ToFermentableType())
	}

	recipe := &beerproto.Recipe{
		Fermentables: arr,
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
			fmt.Println(fmt.Sprintf("Making file %s", file))
			defer f.Close()
			w = f

		} else {
			jsonPath := maltItemsPath[0:len(maltItemsPath)-len(CsvExt)] + JsonExt
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
