package csv

import (
	"strings"

	beerproto "github.com/beerproto/beerproto_go"
)

type Culture struct {
	ID             string   `csv:"ID"`
	Name           string   `csv:"Name"`
	Producer       string   `csv:"Producer"`
	ProductID      string   `csv:"Product ID"`
	Type           string   `csv:"Type"`
	Form           string   `csv:"Form"`
	Species        string   `csv:"Species"`
	PhMin          *float64 `csv:"End pH Min"`
	PhMax          *float64 `csv:"End pH Max"`
	FlavorAroma    string   `csv:"Flavor/Aroma"`
	Pitch          string   `csv:"Pitch"`
	PitchTemp      *float64 `csv:"Pitch Temp ©"`
	Tolerance      *float64 `csv:"Tolerance (%)"`
	AttenuationMin *float64 `csv:"Attenuation Min (%)"`
	AttenuationMax *float64 `csv:"Attenuation Max (%)"`
	Flocculation   string   `csv:"Flocculation (%)"`
	TemperatureMin *float64 `csv:"Temperature Min ©"`
	TemperatureMax *float64 `csv:"Temperature Max ©"`
	BestFor        string   `csv:"Best For"`
	Description    string   `csv:"Description"`
}

func (s Culture) ToCultureInformation() *beerproto.CultureInformation {
	return &beerproto.CultureInformation{
		Id:               s.ID,
		Form:             s.ToForm(),
		Producer:         s.Producer,
		TemperatureRange: toTemperatureRangeType(s.TemperatureMin, s.TemperatureMax),
		Notes:            s.Description,
		BestFor:          s.BestFor,
		Inventory: &beerproto.CultureInventoryType{
			Liquid: &beerproto.VolumeType{
				Unit: beerproto.VolumeType_L,
			},
			Dry: &beerproto.MassType{
				Unit: beerproto.MassUnitType_KG,
			},
			Slant: &beerproto.VolumeType{
				Unit: beerproto.VolumeType_L,
			},
			Culture: &beerproto.VolumeType{
				Unit: beerproto.VolumeType_L,
			},
		},
		ProductId:        s.ProductID,
		Name:             s.Name,
		AlcoholTolerance: toPercent(s.Tolerance),
		Type:             s.ToType(),
		Flocculation:     s.ToFlocculation(),
		AttenuationRange: toPercentRangeType(s.AttenuationMin, s.AttenuationMax),
	}
}

func (s *Culture) ToType() beerproto.CultureBaseType {
	if t, ok := beerproto.CultureBaseType_value[strings.ToUpper(s.Type)]; ok {
		return beerproto.CultureBaseType(t)
	}

	return beerproto.CultureBaseType_NULL_CULTUREBASETYPE
}

func (s *Culture) ToForm() beerproto.CultureBaseForm {
	if t, ok := beerproto.CultureBaseForm_value[strings.ToUpper(s.Form)]; ok {
		return beerproto.CultureBaseForm(t)
	}

	return beerproto.CultureBaseForm_NULL_CULTUREBASEFORM
}

func (s *Culture) ToFlocculation() beerproto.QualitativeRangeType {
	key := strings.ReplaceAll(strings.TrimSpace(strings.ToUpper(s.Flocculation)), " ", "_")
	if t, ok := beerproto.QualitativeRangeType_value[key]; ok {
		return beerproto.QualitativeRangeType(t)
	}

	return beerproto.QualitativeRangeType_NULL_QUALITATIVERANGETYPE
}
