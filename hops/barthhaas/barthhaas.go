package barthhaas


type Hop struct {
	Name                 string `csv:"Name"`
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
	Country              string `csv:"Country"`
	Description          string `csv:"Description"`
}

