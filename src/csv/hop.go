package csv

type Hop struct {
	Name                 string `csv:"Name"`
	AlsoKnownAs          string `csv:"Also Known As"`
	Characteristics      string `csv:"Characteristics"`
	Purpose              string `csv:"Purpose"`
	AlphaAcidLow         string `csv:"Alpha Acid Low"`
	AlphaAcidHigh        string `csv:"Alpha Acid High"`
	BetaAcidLow          string `csv:"Beta Acid Low"`
	BetaAcidHigh         string `csv:"Beta Acid High"`
	CoHumuloneLow        string `csv:"Co-Humulone Low"`
	CoHumuloneHigh       string `csv:"Co-Humulone High"`
	Country              string `csv:"Country"`
	ConeSize             string `csv:"Cone Size"`
	ConeDensity          string `csv:"Cone Density"`
	SeasonalMaturity     string `csv:"Seasonal Maturity"`
	YieldAmount          string `csv:"Yield Amount"`
	GrowthRate           string `csv:"Growth Rate"`
	Resistantto          string `csv:"Resistant to"`
	Susceptibleto        string `csv:"Susceptible to"`
	Storability          string `csv:"Storability"`
	EaseofHarvest        string `csv:"Easeof Harvest"`
	TotalOilComposition  string `csv:"Total Oil Composition"`
	MyrceneOilLow        string `csv:"Myrcene Oil Low"`
	MyrceneOilHigh       string `csv:"Myrcene Oil High"`
	HumuleneOilLow       string `csv:"Humulene Oil Low"`
	HumuleneOilHigh      string `csv:"Humulene Oil High"`
	CaryophylleneOilLow  string `csv:"Caryophyllene Oil Low"`
	CaryophylleneOilHigh string `csv:"Caryophyllene Oil High"`
	FarneseneOilLow      string `csv:"Farnesene Oil Low"`
	FarneseneOilHigh     string `csv:"Farnesene Oil High"`
	Substitutes          string `csv:"Substitutes"`
	StyleGuide           string `csv:"Style Guide"`
	Description          string `csv:"Description"`
}
