package csv

import (
	beerproto "github.com/beerproto/beerproto_go"
	"strings"
)

type Fermentable struct {
	ID               string   `csv:"ID"`
	Name             string   `csv:"Name"`
	Producer         string   `csv:"Producer"`
	Group            string   `csv:"Group"`
	Type             string   `csv:"Type"`
	Standard         string   `csv:"Standard (ASBC/EBC/ION)"`
	Moisture         *float64 `csv:"Moisture (%)"`
	ExtractYield     *float64 `csv:"Extract, F.G. dry (%)"`
	Extract          *float64 `csv:"Extract (L/KG)"`
	FineCoarse       *float64 `csv:"Fine-coarse difference (EBC)"`
	Potential        *float64 `csv:"Potential (SG)"`
	Color            *float64 `csv:"Color (EBC)"`
	DiastaticPower   *float64 `csv:"Diastatic power (ºWK)"`
	ProteinTotal     *float64 `csv:"Protein Total (%)"`
	Solubleprotein   *float64 `csv:"Soluble protein (%)"`
	TotalNitrogen    *float64 `csv:"Total Nitrogen (%)"`
	SolubleNitrogen  *float64 `csv:"Soluble Nitrogen (mg/100g)"`
	MaxInBatch       *float64 `csv:"Max in Batch (%)"`
	FAN              *float64 `csv:"FAN (mg/g)"`
	AlphaAmylase     *float64 `csv:"Alpha Amylase (Dry min)"`
	BetaGlucans      *float64 `csv:"Beta Glucans (mg/l)"`
	Saccharification *int     `csv:"Saccharification time (min)"`
	Viscosity        *float64 `csv:"Viscosity (mPa.s)"`
	DMSP             *float64 `csv:"DMS-P"`
	Kolbach          *float64 `csv:"Kolbach Index"`
	Friability       *float64 `csv:"Friability (%)"`
	WortPH           *float64 `csv:"Wort pH"`
	Country          string   `csv:"Country"`
	Notes            string   `csv:"Notes"`
}

func (s Fermentable) ToFermentableType() *beerproto.FermentableType {
	return &beerproto.FermentableType{
		Id:             s.ID,
		MaxInBatch:     toPercent(s.MaxInBatch),
		Protein:        toPercent(s.ProteinTotal),
		GrainGroup:     s.ToGrainGroup(),
		Yield:          s.ToYield(),
		Type:           s.ToType(),
		Producer:       s.Producer,
		AlphaAmylase:   *s.AlphaAmylase,
		Color:          toColor(s.Color),
		Name:           s.Name,
		DiastaticPower: toDiastaticPowerType(s.DiastaticPower),
		Moisture:       toPercent(s.Moisture),
		Origin:         s.Country,
		KolbachIndex:   *s.Kolbach,
		Friability:     toPercent(s.Friability),
		DiPh:           toAcidityType(s.WortPH),
		Viscosity:      toViscosityType(s.Viscosity),
		DmsP:           toConcentrationType(s.DMSP),
		Fan:            toConcentrationType(s.FAN),
		BetaGlucan:     toConcentrationType(s.BetaGlucans),
		Notes:          s.Notes,
		Inventory: &beerproto.FermentableInventoryType{
			Amount: &beerproto.FermentableInventoryType_Mass{
				Mass: &beerproto.MassType{
					Value: 0,
					Unit: beerproto.MassUnitType_KG,
				},
			},
		},
	}
}

func (s *Fermentable) ToYield() *beerproto.YieldType {
	return &beerproto.YieldType{
		Potential:            toGravity(s.Potential),
		FineCoarseDifference: toPercent(s.FineCoarse),
		FineGrind:            toPercent(s.ExtractYield),
	}
}

func (s *Fermentable) ToType() beerproto.FermentableBaseType {
	if t, ok := beerproto.FermentableBaseType_value[strings.ToUpper(s.Type)]; ok {
		return beerproto.FermentableBaseType(t)
	}

	return beerproto.FermentableBaseType_NULL_FERMENTABLEBASETYPE
}

func (s *Fermentable) ToGrainGroup() beerproto.GrainGroup {
	if t, ok := beerproto.GrainGroup_value[strings.ToUpper(s.Type)]; ok {
		return beerproto.GrainGroup(t)
	}

	return beerproto.GrainGroup_NULL_GRAINGROUP
}
