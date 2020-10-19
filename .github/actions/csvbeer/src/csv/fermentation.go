package csv

import (
	beerproto "github.com/beerproto/beerproto_go"
)

type Fermentation struct {
	ID    string              `csv:"ID"`
	Name  string              `csv:"Name"`
	Steps []*FermentationStep `csv:"-"`
}

func (s *Fermentation) ToFermentationProcedureType(steps []*beerproto.FermentationStepType) *beerproto.FermentationProcedureType {

	return &beerproto.FermentationProcedureType{
		Id:                s.ID,
		Name:              s.Name,
		FermentationSteps: steps,
	}
}

type FermentationStep struct {
	ID             string   `csv:"ID"`
	FermentationID string   `csv:"Fermentation ID"`
	Name           string   `csv:"Name"`
	Temp           *float64 `csv:"Temp (c)"`
	Days           *int64   `csv:"Days"`
}

func (s *FermentationStep) ToFermentationStepType() *beerproto.FermentationStepType {
	return &beerproto.FermentationStepType{
		Id:               s.ID,
		Name:             s.Name,
		StartTemperature: toTemperature(s.Temp),
		StepTime:         toTimeTypeDays(s.Days),
	}
}
