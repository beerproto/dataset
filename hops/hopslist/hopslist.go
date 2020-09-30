package hopslist


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
