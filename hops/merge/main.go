package main

import (
	"fmt"
	"github.com/beerproto/dataset/hops/barthhaas"
	"github.com/beerproto/dataset/hops/hopslist"
	"github.com/gocarina/gocsv"
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	clientsFile, err := loadFile("../hopslist/cmd/hopslist.csv")
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()

	hops := []*hopslist.Hop{}

	if err := gocsv.UnmarshalFile(clientsFile, &hops); err != nil { // Load clients from file
		panic(err)
	}

	clientsFileMore, err := loadFile("../barthhaas/cmd/barthhaas.csv")
	if err != nil {
		panic(err)
	}
	defer clientsFileMore.Close()

	hopsMore := []*barthhaas.Hop{}

	if err := gocsv.UnmarshalFile(clientsFileMore, &hopsMore); err != nil { // Load clients from file
		panic(err)
	}

	hopsMoreMap := map[string]*barthhaas.Hop{}
	for _, hop := range hopsMore {
		hopsMoreMap[hop.Name] = hop
	}

	use := func(s string) string {
		arr := strings.Split(s, "&")
		if len(arr) > 1 {
			t := ""
			for _, item := range arr {
				t += strings.TrimSpace(item) + ", "
			}

			return strings.TrimRight(t, ", ")
		}

		return strings.TrimSpace(arr[0])
	}
	hopsOuput := []*HopOuput{}
	for _, hop := range hops {
		if moreHop, ok := hopsMoreMap[hop.Name]; ok {
			hopOutput := &HopOuput{
				Hop : hop,
			}

			hopOutput.Characteristics = moreHop.Flavors


			hopOutput.Characteristics = strings.ReplaceAll(moreHop.Flavors, "\"", "'")
			hopOutput.Characteristics = strings.ReplaceAll(hopOutput.Characteristics, "\r", " ")
			hopOutput.Characteristics = strings.ReplaceAll(hopOutput.Characteristics, "\n", " ")

			if hopOutput.AlphaAcidHigh == "" {
				hopOutput.AlphaAcidLow = moreHop.AlphaAcidsLow
				hopOutput.AlphaAcidHigh = moreHop.AlphaAcidsHigh
			}

			if hopOutput.BetaAcidHigh == "" {
				hopOutput.BetaAcidLow = moreHop.BetaAcidsLow
				hopOutput.BetaAcidHigh = moreHop.BetaAcidsHigh
			}

			if hopOutput.BetaAcidHigh == "" {
				hopOutput.BetaAcidLow = moreHop.BetaAcidsLow
				hopOutput.BetaAcidHigh = moreHop.BetaAcidsHigh
			}

			if hopOutput.TotalOilHigh == "" {
				hopOutput.TotalOilLow = moreHop.TotalOilLow
				hopOutput.TotalOilHigh = moreHop.TotalOilHigh

			}

			if hopOutput.MyrceneOilHigh == "" {
				hopOutput.MyrceneOilLow = moreHop.MyrceneLow
				hopOutput.MyrceneOilHigh = moreHop.MyrceneHigh
			}

			if hopOutput.LinaloolOilHigh == "" {
				hopOutput.LinaloolOilLow = moreHop.LinaloolLow
				hopOutput.LinaloolOilHigh = moreHop.LinaloolHigh
			}

			if hopOutput.PolyphenolsOilHigh == "" {
				hopOutput.PolyphenolsOilLow = moreHop.TotalPolyphenolsLow
				hopOutput.PolyphenolsOilHigh = moreHop.TotalPolyphenolsHigh
			}

			hopOutput.Purpose = use(hopOutput.Purpose)
			hopOutput.ID = uuid.New().String()
			hopsOuput = append(hopsOuput, hopOutput)
			delete(hopsMoreMap, hop.Name)
		} else {
			hopOutput := &HopOuput{
				Hop : hop,
			}
			hopOutput.ID = uuid.New().String()
			hopOutput.Purpose = use(hopOutput.Purpose)
			hopsOuput = append(hopsOuput, hopOutput)
		}

	}


	for _, hop := range hopsMoreMap {
		hopOutput := &HopOuput{
			Hop: &hopslist.Hop{},
		}
		hopOutput.Name = hop.Name


		hopOutput.Characteristics = strings.ReplaceAll(hop.Flavors, "\"", "'")
		hopOutput.Characteristics = strings.ReplaceAll(hopOutput.Characteristics, "\r", " ")
		hopOutput.Characteristics = strings.ReplaceAll(hopOutput.Characteristics, "\n", " ")
		hopOutput.AlphaAcidLow = hop.AlphaAcidsLow
		hopOutput.AlphaAcidHigh = hop.AlphaAcidsHigh
		hopOutput.BetaAcidLow = hop.BetaAcidsLow
		hopOutput.BetaAcidHigh = hop.BetaAcidsHigh
		hopOutput.TotalOilLow = hop.TotalOilLow
		hopOutput.TotalOilHigh = hop.TotalOilHigh
		hopOutput.MyrceneOilLow = hop.MyrceneLow
		hopOutput.MyrceneOilHigh = hop.MyrceneHigh
		hopOutput.LinaloolOilLow = hop.LinaloolLow
		hopOutput.LinaloolOilHigh = hop.LinaloolHigh
		hopOutput.PolyphenolsOilLow = hop.TotalPolyphenolsLow
		hopOutput.PolyphenolsOilHigh = hop.TotalPolyphenolsHigh
		hopOutput.Country = hop.Country
		hopOutput.Description = strings.ReplaceAll(hop.Description, "\"", "'")
		hopOutput.Description = strings.ReplaceAll(hopOutput.Description, "\r", " ")
		hopOutput.Description = strings.ReplaceAll(hopOutput.Description, "\n", " ")
		//hopOutput.Description = ""

		hopOutput.Purpose = use(hopOutput.Purpose)
		hopOutput.ID = uuid.New().String()
		hopsOuput = append(hopsOuput, hopOutput)
	}

	clientsFileOutput, err := os.OpenFile("../hops.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFileOutput.Close()

	err = gocsv.MarshalFile(&hopsOuput, clientsFileOutput)
	if err != nil {
		panic(err)
	}

}

type HopOuput struct {
	*hopslist.Hop
}


func loadFile(filePath string) (*os.File, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		return nil, fmt.Errorf("file not found %v: %w", filePath, err)
	}
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("file not found %v", filePath)
	}

	filePath, err = filepath.Abs(filePath)
	if err != nil {
		return nil, fmt.Errorf("file not found %v: %w", filePath, err)
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to open %v: %w", filePath, err)
	}
	return file, nil
}
