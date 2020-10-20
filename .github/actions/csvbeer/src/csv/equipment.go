package csv

import (
	beerproto "github.com/beerproto/beerproto_go"
)

type Equipment struct {
	ID    string            `csv:"ID"`
	Name  string            `csv:"Name"`
	Items []*EquipmentItems `csv:"-"`
}

type EquipmentItems struct {
	ID                  string   `csv:"ID"`
	EquipmentID         string   `csv:"Equipment ID"`
	Name                string   `csv:"Name"`
	MaximumVolume       *float64 `csv:"Maximum Volume (L)"`
	BoilRatePerHour     *float64 `csv:"Boil Rate Per Hour (L)"`
	DrainRatePerMinute  *float64 `csv:"Drain Rate Per Minute (L)"`
	SpecificHeat        *float64 `csv:"Specific Heat (Cal/(g*C))"`
	GrainAbsorptionRate *float64 `csv:"Grain Absorption Rate (L/kg)"`
	Loss                *float64 `csv:"Loss (L)"`
	Weight              *float64 `csv:"Weight (kg)"`
	Form                int      `csv:"Form (EquipmentItemType.EquipmentBaseForm)"`
}

func (s *Equipment) ToEquipmentType(items []*beerproto.EquipmentItemType) *beerproto.EquipmentType {

	return &beerproto.EquipmentType{
		Id:             s.ID,
		Name:           s.Name,
		EquipmentItems: items,
	}
}

func (s *EquipmentItems) ToEquipmentItems() *beerproto.EquipmentItemType {
	return &beerproto.EquipmentItemType{
		Id:                  s.ID,
		Name:                s.Name,
		BoilRatePerHour:     toVolumeType(s.BoilRatePerHour, beerproto.VolumeType_L),
		MaximumVolume:       toVolumeType(s.MaximumVolume, beerproto.VolumeType_L),
		DrainRatePerMinute:  toVolumeType(s.DrainRatePerMinute, beerproto.VolumeType_L),
		SpecificHeat:        toSpecificHeatType(s.SpecificHeat),
		GrainAbsorptionRate: toSpecificVolumeType(s.GrainAbsorptionRate),
		Loss:                toVolumeType(s.Loss, beerproto.VolumeType_L),
		Weight:              toMassType(s.Weight),
		Form:                beerproto.EquipmentItemType_EquipmentBaseForm(s.Form),
	}
}
