package csv

import (
	beerproto "github.com/beerproto/beerproto_go"
	"strings"
)

type Hop struct {
	ID                   string   `csv:"ID"`
	Name                 string   `csv:"Name"`
	Purpose              string   `csv:"Purpose"`
	AlphaAcidLow         *float64 `csv:"Alpha Acid Low (%)"`
	AlphaAcidHigh        *float64 `csv:"Alpha Acid High (%)"`
	BetaAcidLow          *float64 `csv:"Beta Acid Low (%)"`
	BetaAcidHigh         *float64 `csv:"Beta Acid High (%)"`
	CoHumuloneLow        *float64 `csv:"Co-Humulone Low (%)"`
	CoHumuloneHigh       *float64 `csv:"Co-Humulone High (%)"`
	Country              string   `csv:"Country"`
	Storability          string   `csv:"Storability"`
	TotalOilLow          *float64 `csv:"Total Oil Composition Low (mL/100g)"`
	TotalOilHigh         *float64 `csv:"Total Oil Composition High (mL/100g)"`
	MyrceneOilLow        *float64 `csv:"Myrcene Oil Low (%)"`
	MyrceneOilHigh       *float64 `csv:"Myrcene Oil High (%)"`
	HumuleneOilLow       *float64 `csv:"Humulene Oil Low (%)"`
	HumuleneOilHigh      *float64 `csv:"Humulene Oil High (%)"`
	CaryophylleneOilLow  *float64 `csv:"Caryophyllene Oil Low (%)"`
	CaryophylleneOilHigh *float64 `csv:"Caryophyllene Oil High (%)"`
	FarneseneOilLow      *float64 `csv:"Farnesene Oil Low (%)"`
	FarneseneOilHigh     *float64 `csv:"Farnesene Oil High (%)"`
	LinaloolOilLow       *float64 `csv:"Linalool Oil Low (%)"`
	LinaloolOilHigh      *float64 `csv:"Linalool Oil High (%)"`
	PolyphenolsOilLow    *float64 `csv:"Polyphenols Oil Low (%)"`
	PolyphenolsOilHigh   *float64 `csv:"Polyphenols Oil High (%)"`
	Substitutes          string   `csv:"Substitutes"`
	StyleGuide           string   `csv:"Style Guide"`
	AlsoKnownAs          string   `csv:"Also Known As"`
	Characteristics      string   `csv:"Characteristics"`
	Description          string   `csv:"Description"`
}

func (s *Hop) ToVarietyInformation(partition string) *beerproto.VarietyInformation {
	return &beerproto.VarietyInformation{
		Id:        partition + s.ID,
		Inventory: &beerproto.HopInventoryType{},
		Type:      s.ToType(),
		OilContent: s.ToOilContentType(),
		//PercentLost:
		AlphaAcid:   averagePercent(s.AlphaAcidLow, s.AlphaAcidHigh),
		BetaAcid:    averagePercent(s.BetaAcidLow, s.BetaAcidHigh),
		Name:        s.Name,
		Origin:      s.Country,
		Substitutes: s.Substitutes,
		Notes:       s.Characteristics,
	}
}

func (s *Hop) ToOilContentType() *beerproto.OilContentType {
	return &beerproto.OilContentType{
		Polyphenols:        averagePercent(s.PolyphenolsOilLow, s.PolyphenolsOilHigh),
		TotalOilMlPer_100G: total(s.TotalOilLow, s.TotalOilHigh),
		Farnesene:          averagePercent(s.FarneseneOilLow, s.FarneseneOilHigh),
		//Limonene: averagePercent(s.l, s.PolyphenolsOilHigh),
		//Nerol: averagePercent(s., s.PolyphenolsOilHigh),
		//Geraniol: averagePercent(s.g, s.PolyphenolsOilHigh),
		//BPinene: averagePercent(s.b, s.PolyphenolsOilHigh),
		Linalool:      averagePercent(s.LinaloolOilLow, s.LinaloolOilHigh),
		Caryophyllene: averagePercent(s.CaryophylleneOilLow, s.CaryophylleneOilHigh),
		Cohumulone:    averagePercent(s.CoHumuloneLow, s.CoHumuloneHigh),
		//Xanthohumol: averagePercent(s.x, s.PolyphenolsOilHigh),
		Humulene: averagePercent(s.HumuleneOilLow, s.HumuleneOilHigh),
		Myrcene:  averagePercent(s.MyrceneOilLow, s.MyrceneOilHigh),
		//Pinene: averagePercent(s.p, s.PolyphenolsOilHigh),
	}
}

func total(low, high *float64) float64 {
	if low == nil {
		return 0
	}
	if high == nil {
		return *low
	}

	sum := *low + *high
	avg := (float64(sum)) / (float64(2))
	return avg
}


func averagePercent(low, high *float64) *beerproto.PercentType {
	if low == nil {
		return nil
	}
	if high == nil || *high == 0{
		return &beerproto.PercentType{
			Value: *low,
			Unit: beerproto.PercentType_PERCENT_SIGN,
		}
	}

	sum := *low + *high
	avg := (float64(sum)) / (float64(2))
	return &beerproto.PercentType{
		Value: avg,
		Unit: beerproto.PercentType_PERCENT_SIGN,
	}
}

func (s *Hop) ToType() beerproto.VarietyInformation_VarietyInformationType {

	hasBittering := strings.Contains(s.Purpose, "bittering")
	hasAroma := strings.Contains(s.Purpose, "aroma")

	if hasAroma && hasBittering {
		return beerproto.VarietyInformation_AROMA_BITTERING
	}
	if hasAroma {
		return beerproto.VarietyInformation_AROMA
	}
	if hasBittering {
		return beerproto.VarietyInformation_BITTERING
	}

	return beerproto.VarietyInformation_NULL_VARIETYINFORMATIONTYPE
}
