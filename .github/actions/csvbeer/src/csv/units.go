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

func toTemperatureRangeType(low, high *float64) *beerproto.TemperatureRangeType {

	return &beerproto.TemperatureRangeType{
		Minimum: toTemperature(low),
		Maximum: toTemperature(high),
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


func toVolumeType(value *float64, t beerproto.VolumeType_VolumeUnitType) *beerproto.VolumeType {
	if value == nil {
		return nil
	}

	return &beerproto.VolumeType{
		Unit: t,
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

func toPercent(value *float64) *beerproto.PercentType {
	if value == nil {
		return nil
	}
	return &beerproto.PercentType{
		Value: *value,
		Unit: beerproto.PercentType_PERCENT_SIGN,
	}
}


func toPercentRangeType(low, high *float64) *beerproto.PercentRangeType {
	return &beerproto.PercentRangeType{
		Minimum: toPercent(low),
		Maximum: toPercent(high),
	}
}


func toGravity(value *float64) *beerproto.GravityType {
	if value == nil {
		return nil
	}
	return &beerproto.GravityType{
		Value: *value,
		Unit: beerproto.GravityUnitType_SG,
	}
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

func toColor(value *float64) *beerproto.ColorType {
	if value == nil {
		return nil
	}
	return &beerproto.ColorType{
		Unit:  beerproto.ColorUnitType_EBC,
		Value: *value,
	}
}

func toDiastaticPowerType(value *float64) *beerproto.DiastaticPowerType {
	if value == nil {
		return nil
	}

	return &beerproto.DiastaticPowerType{
		Unit:  beerproto.DiastaticPowerUnitType_WK,
		Value: *value,
	}
}

func toAcidityType(value *float64) *beerproto.AcidityType {
	if value == nil {
		return nil
	}

	return &beerproto.AcidityType{
		Unit:  beerproto.AcidityUnitType_PH,
		Value: *value,
	}
}

func toViscosityType(value *float64) *beerproto.ViscosityType {
	if value == nil {
		return nil
	}

	return &beerproto.ViscosityType{
		Unit:  beerproto.ViscosityUnitType_MPAS,
		Value: *value,
	}
}