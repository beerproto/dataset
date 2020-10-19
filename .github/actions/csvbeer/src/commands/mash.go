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

func ParseMash(mashPath, mashStepPath string, output Output, file string) {
	fmt.Println(fmt.Sprintf("Step 1/2 : Loading %s", mashStepPath))

	indexFile, err := loadFile(mashStepPath)
	if err != nil {
		panic(fmt.Errorf("failed to load file %v: %w", mashStepPath, err))
	}
	defer indexFile.Close()

	var mashes []*csv.Mash

	if err := gocsv.UnmarshalFile(indexFile, &mashes); err != nil {
		panic(fmt.Errorf("%v does not match Index %v format: %w", mashStepPath, CsvExt, err))
	}

	fmt.Println(fmt.Sprintf("Step 2/2 : Loading %s", mashPath))

	mashStepsFile, err := loadFile(mashPath)
	if err != nil {
		panic(fmt.Errorf("failed to load file %v: %w", mashPath, err))
	}
	defer mashStepsFile.Close()

	var mashSteps []*csv.MashStep

	if err := gocsv.UnmarshalFile(mashStepsFile, &mashSteps); err != nil {
		panic(fmt.Errorf("%v does not match MashSteps %v format: %w", mashPath, CsvExt, err))
	}

	fmt.Println("Successfully parased")

	items := map[string][]*beerproto.MashStepType{}
	for _, item := range mashSteps {
		arr := items[item.EquipmentID]
		arr = append(arr, item.ToMashStepType())
		items[item.EquipmentID] = arr
	}

	var arr []*beerproto.MashProcedureType
	for _, e := range mashes {
		if equipmentItems, ok := items[e.ID]; ok {
			arr = append(arr, e.ToMashProcedureType(equipmentItems))
		}
	}

	recipe := &beerproto.Recipe{
		Mashes: arr,
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
			jsonPath := mashPath[0:len(mashPath)-len(CsvExt)] + JsonExt
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
