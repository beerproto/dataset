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

func ParseFermentation(fermentationPath, fermentationStepPath string, output Output, file string) {
	fmt.Println(fmt.Sprintf("Step 1/2 : Loading %s", fermentationStepPath))

	indexFile, err := loadFile(fermentationStepPath)
	if err != nil {
		panic(fmt.Errorf("failed to load file %v: %w", fermentationStepPath, err))
	}
	defer indexFile.Close()

	var fermentations []*csv.Fermentation

	if err := gocsv.UnmarshalFile(indexFile, &fermentations); err != nil {
		panic(fmt.Errorf("%v does not match Index %v format: %w", fermentationStepPath, CsvExt, err))
	}

	fmt.Println(fmt.Sprintf("Step 2/2 : Loading %s", fermentationPath))

	mashStepsFile, err := loadFile(fermentationPath)
	if err != nil {
		panic(fmt.Errorf("failed to load file %v: %w", fermentationPath, err))
	}
	defer mashStepsFile.Close()

	var fermentationSteps []*csv.FermentationStep

	if err := gocsv.UnmarshalFile(mashStepsFile, &fermentationSteps); err != nil {
		panic(fmt.Errorf("%v does not match FermentationStep %v format: %w", fermentationPath, CsvExt, err))
	}

	fmt.Println("Successfully parased")

	items := map[string][]*beerproto.FermentationStepType{}
	for _, item := range fermentationSteps {
		arr := items[item.FermentationID]
		arr = append(arr, item.ToFermentationStepType())
		items[item.FermentationID] = arr
	}

	var arr []*beerproto.FermentationProcedureType
	for _, e := range fermentations {
		if equipmentItems, ok := items[e.ID]; ok {
			arr = append(arr, e.ToFermentationProcedureType(equipmentItems))
		}
	}

	recipe := &beerproto.Recipe{
		Fermentations: arr,
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
			jsonPath := fermentationPath[0:len(fermentationPath)-len(CsvExt)] + JsonExt
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
