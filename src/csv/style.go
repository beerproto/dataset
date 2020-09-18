package csv

import (
	"fmt"
	"strconv"

	beerproto "github.com/beerproto/beerproto_go"
)

type Index struct {
	ID    int    `csv:"ID"`
	Style string `csv:"Style"`
}

type Style struct {
	ID                          int      `csv:"ID"`
	StyleID                     int      `csv:"Style ID"`
	Name                        string   `csv:"Name"`
	OGLow                       *float64 `csv:"OG Low,omitempty"`
	OGHigh                      *float64 `csv:"OG High,omitempty"`
	OGPlatoLow                  *float64 `csv:"OG Plato Low,omitempty"`
	OGPlatoHigh                 *float64 `csv:"OG Plato High,omitempty"`
	FGLow                       *float64 `csv:"FG Low,omitempty"`
	FGHigh                      *float64 `csv:"FG High,omitempty"`
	FGPlatoLow                  *float64 `csv:"FG Plato Low,omitempty"`
	FGPlatoHigh                 *float64 `csv:"FG Plato High,omitempty"`
	AlcoholbyWeightLow          *float64 `csv:"Alcohol by Weight Low (%),omitempty"`
	AlcoholbyWeightHigh         *float64 `csv:"Alcohol by Weight High (%),omitempty"`
	VolumeLow                   *float64 `csv:"Volume Low (%),omitempty"`
	VolumeHigh                  *float64 `csv:"Volume High (%),omitempty"`
	BitternessLow               *int     `csv:"Bitterness Low (IBU),omitempty"`
	BitternessHigh              *int     `csv:"Bitterness High (IBU),omitempty"`
	ColorLowSRM                 *int     `csv:"Color Low SRM,omitempty"`
	ColorHighSRM                *int     `csv:"Color High SRM,omitempty"`
	ColorLowEBC                 *int     `csv:"Color Low EBC,omitempty"`
	ColorHighEBC                *int     `csv:"Color High EBC,omitempty"`
	Color                       string   `csv:"Color"`
	Clarity                     string   `csv:"Clarity"`
	PerceivedMaltAromaFlavor    string   `csv:"Perceived Malt Aroma & Flavor"`
	PerceivedHopAromaFlavor     string   `csv:"Perceived Hop Aroma & Flavor"`
	PerceivedBitterness         string   `csv:"Perceived Bitterness"`
	FermentationCharacteristics string   `csv:"Fermentation Characteristics"`
	Body                        string   `csv:"Body"`
	AdditionalNotes             string   `csv:"Additional notes"`
}

func (s *Style) ToStyleType(category string) *beerproto.StyleType {
	return &beerproto.StyleType{
		Id:             strconv.Itoa(s.ID),
		Aroma:          fmt.Sprintf("%s \n%s", s.PerceivedHopAromaFlavor, s.PerceivedMaltAromaFlavor),
		Flavor:         fmt.Sprintf("%s \n%s", s.PerceivedHopAromaFlavor, s.PerceivedMaltAromaFlavor),
		CategoryNumber: int32(s.StyleID),
		Notes:          s.FermentationCharacteristics + s.AdditionalNotes,
		Mouthfeel:      s.Body,
		FinalGravity: s.toFG(),
		Color: s.toColor(),
		OriginalGravity: s.toOG(),
		Name: s.Name,
		AlcoholByVolume: s.toAlcoholByVolume(),
		InternationalBitternessUnits: s.toInternationalBitternessUnits(),
		Appearance: s.Clarity,
		Category:   category,
		Type:       toStyleCategories(s.StyleID),
	}
}

func (s *Style) toInternationalBitternessUnits() *beerproto.BitternessRangeType {
	toInternationalBitternessUnits := func(value *int) *beerproto.BitternessType {
		if value == nil {
			return nil
		}
		return &beerproto.BitternessType{
			Unit:  beerproto.BitternessType_IBUs,
			Value: float64(*value),
		}
	}

	return  &beerproto.BitternessRangeType{
		Minimum: toInternationalBitternessUnits(s.BitternessLow),
		Maximum: toInternationalBitternessUnits(s.BitternessHigh),
	}
}


func (s *Style) toAlcoholByVolume() *beerproto.PercentRangeType {
	toAlcoholByVolume := func(value *float64) *beerproto.PercentType {
		if value == nil {
			return nil
		}
		return &beerproto.PercentType{
			Unit:  beerproto.PercentType_PERCENT_SIGN,
			Value: *value,
		}
	}

	return  &beerproto.PercentRangeType{
		Minimum: toAlcoholByVolume(s.AlcoholbyWeightLow),
		Maximum: toAlcoholByVolume(s.AlcoholbyWeightHigh),
	}
}


func (s *Style) toOG() *beerproto.GravityRangeType {
	toOG := func(value *float64) *beerproto.GravityType {
		if value == nil {
			return nil
		}
		return &beerproto.GravityType{
			Unit:  beerproto.GravityUnitType_SG,
			Value: *value,
		}
	}

	return  &beerproto.GravityRangeType{
		Minimum: toOG(s.OGLow),
		Maximum: toOG(s.OGHigh),
	}
}

func (s *Style) toFG() *beerproto.GravityRangeType {
	toFG := func(value *float64) *beerproto.GravityType {
		if value == nil {
			return nil
		}
		return &beerproto.GravityType{
			Unit:  beerproto.GravityUnitType_SG,
			Value: *value,
		}
	}

	return  &beerproto.GravityRangeType{
		Minimum: toFG(s.FGLow),
		Maximum: toFG(s.FGHigh),
	}
}

func (s *Style) toColor() *beerproto.ColorRangeType {

	toColor := func(value *int) *beerproto.ColorType {
		if value == nil {
			return nil
		}
		return &beerproto.ColorType{
			Unit:  beerproto.ColorUnitType_SRM,
			Value: float64(*value),
		}
	}

	return &beerproto.ColorRangeType{
		Minimum: toColor(s.ColorLowSRM),
		Maximum: toColor(s.ColorHighSRM),
	}
}

func toStyleCategories(index int) beerproto.StyleType_StyleCategories {
	switch index {
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10:
		return beerproto.StyleType_BEER
	}

	return beerproto.StyleType_NULL_STYLECATEGORIES
}
