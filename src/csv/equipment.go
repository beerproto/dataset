package csv

import (
	beerproto "github.com/beerproto/beerproto_go"
)

type Equipment struct {
	ID    string               `csv:"ID"`
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
		BoilRatePerHour:     s.toVolumeType(s.BoilRatePerHour),
		MaximumVolume:       s.toVolumeType(s.MaximumVolume),
		DrainRatePerMinute:  s.toVolumeType(s.DrainRatePerMinute),
		SpecificHeat:        s.toSpecificHeatType(s.SpecificHeat),
		GrainAbsorptionRate: s.toSpecificVolumeType(s.GrainAbsorptionRate),
		Loss:                s.toVolumeType(s.Loss),
		Weight:              s.toMassType(s.Weight),
		Form:                beerproto.EquipmentItemType_EquipmentBaseForm(s.Form),
	}
}

func (s *EquipmentItems) toMassType(value *float64) *beerproto.MassType {
	if value == nil {
		return nil
	}

	return &beerproto.MassType{
		Unit:  beerproto.MassUnitType_KG,
		Value: *value,
	}
}

func (s *EquipmentItems) toSpecificVolumeType(value *float64) *beerproto.SpecificVolumeType {
	if value == nil {
		return nil
	}

	return &beerproto.SpecificVolumeType{
		Unit:  beerproto.SpecificVolumeType_LKG,
		Value: *value,
	}
}

func (s *EquipmentItems) toSpecificHeatType(value *float64) *beerproto.SpecificHeatType {
	if value == nil {
		return nil
	}

	return &beerproto.SpecificHeatType{
		Unit:  beerproto.SpecificHeatUnitType_CALGC,
		Value: *value,
	}
}

func (s *EquipmentItems) toVolumeType(value *float64) *beerproto.VolumeType {
	if value == nil {
		return nil
	}

	return &beerproto.VolumeType{
		Unit:  beerproto.VolumeType_L,
		Value: *value,
	}
}
