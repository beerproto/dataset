package csv

import (
	beerproto "github.com/beerproto/beerproto_go"
)

func toTemperature(value *float64) *beerproto.TemperatureType {

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
