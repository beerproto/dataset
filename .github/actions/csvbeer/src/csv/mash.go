package csv

import (
	beerproto "github.com/beerproto/beerproto_go"
	"strings"
)

type Mash struct {
	ID    string      `csv:"ID"`
	Name  string      `csv:"Name"`
	Steps []*MashStep `csv:"-"`
}

func (m Mash) ToMashProcedureType(items []*beerproto.MashStepType) *beerproto.MashProcedureType {
	return &beerproto.MashProcedureType{
		Id:        m.ID,
		Name:      m.Name,
		MashSteps: items,
	}
}

type MashStep struct {
	ID          string   `csv:"ID"`
	EquipmentID string   `csv:"Mash ID"`
	Name        string   `csv:"Name"`
	Type        string   `csv:"Type"`
	Temp        *float64 `csv:"Temp ©"`
	Time        *int64   `csv:"Time(min)"`
	RampTime    *int64   `csv:"Ramp Time(min)"`
}

func (s MashStep) ToMashStepType() *beerproto.MashStepType {
	return &beerproto.MashStepType{
		Id:              s.ID,
		Name:            s.Name,
		Type:            s.ToType(),
		StepTemperature: toTemperature(s.Temp),
		StepTime:        toTimeType(s.Time),
		RampTime:        toTimeType(s.RampTime),
	}
}

func (s *MashStep) ToType() beerproto.MashStepType_MashStepTypeType {
	if t, ok := beerproto.MashStepType_MashStepTypeType_value[strings.ToUpper(s.Type)]; ok {
		return beerproto.MashStepType_MashStepTypeType(t)
	}

	return beerproto.MashStepType_NULL
}
