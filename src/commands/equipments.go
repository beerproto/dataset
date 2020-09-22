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

func ParseEquipments(equipmentItemsPath, indexPath string, output Output) {
	fmt.Println(fmt.Sprintf("Step 1/2 : Loading %s", indexPath))

	indexFile, err := loadFile(indexPath)
	if err != nil {
		panic(err)
	}
	defer indexFile.Close()

	var equipments []*csv.Equipment

	if err := gocsv.UnmarshalFile(indexFile, &equipments); err != nil {
		panic(fmt.Errorf("%v does not match Index %v format: %w", indexPath, CsvExt, err))
	}

	fmt.Println(fmt.Sprintf("Step 2/2 : Loading %s", equipmentItemsPath))

	equipmentItemsFile, err := loadFile(equipmentItemsPath)
	if err != nil {
		panic(err)
	}
	defer equipmentItemsFile.Close()

	var equipmentItems []*csv.EquipmentItems

	if err := gocsv.UnmarshalFile(equipmentItemsFile, &equipmentItems); err != nil {
		panic(fmt.Errorf("%v does not match EquipmentItems %v format: %w", equipmentItemsPath, CsvExt, err))
	}

	fmt.Println("Successfully parased")

	items := map[int][]*beerproto.EquipmentItemType{}
	for _, item := range equipmentItems {
		arr := items[item.EquipmentID]
		arr = append(arr, item.ToEquipmentItems())
		items[item.EquipmentID] = arr
	}

	var arr []*beerproto.EquipmentType
	for _, e := range equipments {
		if equipmentItems, ok := items[e.ID]; ok {
			arr = append(arr, e.ToEquipmentType(equipmentItems))
		}
	}

	recipe := &beerproto.Recipe{
		Equipments: arr,
	}

	var w io.Writer

	switch output {
	case TTY:
		w = os.Stdout
		break
	case FILE:
		jsonPath := equipmentItemsPath[0:len(equipmentItemsPath)-len(CsvExt)] + JsonExt
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
