package main

import (
	"github.com/gocarina/gocsv"
	"github.com/google/uuid"
	"os"
	"strings"
	)

func main() {
	clientsFile, err := os.OpenFile("hops.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()

	hops := []*Hop{}

	if err := gocsv.UnmarshalFile(clientsFile, &hops); err != nil { // Load clients from file
		panic(err)
	}

	clientsFileMore, err := os.OpenFile("hopsMore.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFileMore.Close()

	hopsMore := []*HopMore{}

	if err := gocsv.UnmarshalFile(clientsFileMore, &hopsMore); err != nil { // Load clients from file
		panic(err)
	}

	hopsMoreMap := map[string]*HopMore{}
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

	clientsFileOutput, err := os.OpenFile("output.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFileOutput.Close()

	err = gocsv.MarshalFile(&hopsOuput, clientsFileOutput)
	if err != nil {
		panic(err)
	}

}

type HopMore struct {
	Name                   string `csv:"Name"`
	Flavors              string `csv:"Flavors"`
	AlphaAcidsLow        string `csv:"Alpha-Acids Low (%)"`
	AlphaAcidsHigh       string `csv:"Alpha-Acids High (%)"`
	BetaAcidsLow         string `csv:"Beta-Acids Low (%)"`
	BetaAcidsHigh        string `csv:"Beta-Acids High (%)"`
	TotalOilLow          string `csv:"Total Oil Low (ML/100G)"`
	TotalOilHigh         string `csv:"Total Oil High (ML/100G)"`
	MyrceneLow           string `csv:"Myrcene Low (%)"`
	MyrceneHigh          string `csv:"Myrcene High (%)"`
	LinaloolLow          string `csv:"Linalool Low (%)"`
	LinaloolHigh         string `csv:"Linalool High (%)"`
	TotalPolyphenolsLow  string `csv:"Total Polyphenols Low (%)"`
	TotalPolyphenolsHigh string `csv:"Total Polyphenols High (%)"`
	SumOfTerpeneAlcohols string `csv:"Sum Of Terpene Alcohols"`
	TotalOilMinusMyrcene string `csv:"Total Oil Minus Myrcene"`
}

type Hop struct {
	ID                   string `csv:"ID"`
	Name                 string `csv:"Name"`
	Purpose              string `csv:"Purpose"`
	AlphaAcidLow         string `csv:"Alpha Acid Low (%)"`
	AlphaAcidHigh        string `csv:"Alpha Acid High (%)"`
	BetaAcidLow          string `csv:"Beta Acid Low (%)"`
	BetaAcidHigh         string `csv:"Beta Acid High (%)"`
	CoHumuloneLow        string `csv:"Co-Humulone Low (%)"`
	CoHumuloneHigh       string `csv:"Co-Humulone High (%)"`
	Country              string `csv:"Country"`
	Storability          string `csv:"Storability"`
	TotalOilLow          string `csv:"Total Oil Composition Low (mL/100g)"`
	TotalOilHigh         string `csv:"Total Oil Composition High (mL/100g)"`
	MyrceneOilLow        string `csv:"Myrcene Oil Low (%)"`
	MyrceneOilHigh       string `csv:"Myrcene Oil High (%)"`
	HumuleneOilLow       string `csv:"Humulene Oil Low (%)"`
	HumuleneOilHigh      string `csv:"Humulene Oil High (%)"`
	CaryophylleneOilLow  string `csv:"Caryophyllene Oil Low (%)"`
	CaryophylleneOilHigh string `csv:"Caryophyllene Oil High (%)"`
	FarneseneOilLow      string `csv:"Farnesene Oil Low (%)"`
	FarneseneOilHigh     string `csv:"Farnesene Oil High (%)"`
	LinaloolOilLow       string `csv:"Linalool Oil Low (%)"`
	LinaloolOilHigh      string `csv:"Linalool Oil High (%)"`
	PolyphenolsOilLow    string `csv:"Polyphenols Oil Low (%)"`
	PolyphenolsOilHigh   string `csv:"Polyphenols Oil High (%)"`
	Substitutes          string `csv:"Substitutes"`
	StyleGuide           string `csv:"Style Guide"`
	AlsoKnownAs          string `csv:"Also Known As"`
	Characteristics      string `csv:"Characteristics"`
	Description          string `csv:"Description"`
}


type HopOuput struct {
	*Hop
}
