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

func ParseHops(hopItemsPath string, output Output, file string) {
	hopsItemsFile, err := loadFile(hopItemsPath)
	if err != nil {
		panic(err)
	}
	defer hopsItemsFile.Close()

	var hops []*csv.Hop

	if err := gocsv.UnmarshalFile(hopsItemsFile, &hops); err != nil {
		panic(fmt.Errorf("%v does not match Hop %v format: %w", hopItemsPath, CsvExt, err))
	}

	fmt.Println("Successfully parased")

	var arr []*beerproto.VarietyInformation
	for _, e := range hops {
		arr = append(arr, e.ToVarietyInformation())
	}

	recipe := &beerproto.Recipe{
		HopVarieties: arr,
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
			jsonPath := hopItemsPath[0:len(hopItemsPath)-len(CsvExt)] + JsonExt
			f, err := os.Create(jsonPath)
			if err != nil {
				panic(fmt.Errorf("failed to create file %s: %w", jsonPath, err))
			}
			fmt.Println(fmt.Sprintf("Makine file %s", jsonPath))
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
