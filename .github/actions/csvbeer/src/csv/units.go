package csv

import (
	beerproto "github.com/beerproto/beerproto_go"
)

func toTemperature(value *float64) *beerproto.TemperatureType {

	if value == nil {
		return &beerproto.TemperatureType{
			Unit: beerproto.TemperatureUnitType_C,
		}
	}

	return &beerproto.TemperatureType{
		Unit: beerproto.TemperatureUnitType_C,
		Value: *value,
	}
}


func toTimeType(value *int64) *beerproto.TimeType {
	return &beerproto.TimeType{
		Unit: beerproto.TimeType_MIN,
		Value: *value,
	}
}

func toTimeTypeDays(value *int64) *beerproto.TimeType {
	return &beerproto.TimeType{
		Unit: beerproto.TimeType_DAY,
		Value: *value,
	}
}

func toConcentrationType(value *float64) *beerproto.ConcentrationType {
	if value == nil {
		return nil
	}

	return &beerproto.ConcentrationType{
		Unit: beerproto.ConcentrationUnitType_MGL,
		Value: *value,
	}
}


func toVolumeType(value *float64) *beerproto.VolumeType {
	if value == nil {
		return nil
	}

	return &beerproto.VolumeType{
		Unit:  beerproto.VolumeType_L,
		Value: *value,
	}
}


func toMassType(value *float64) *beerproto.MassType {
	if value == nil {
		return nil
	}

	return &beerproto.MassType{
		Unit:  beerproto.MassUnitType_KG,
		Value: *value,
	}
}

func toSpecificVolumeType(value *float64) *beerproto.SpecificVolumeType {
	if value == nil {
		return nil
	}

	return &beerproto.SpecificVolumeType{
		Unit:  beerproto.SpecificVolumeType_LKG,
		Value: *value,
	}
}

func toSpecificHeatType(value *float64) *beerproto.SpecificHeatType {
	if value == nil {
		return nil
	}

	return &beerproto.SpecificHeatType{
		Unit:  beerproto.SpecificHeatUnitType_CALGC,
		Value: *value,
	}
}

