package csv

import (
	beerproto "github.com/beerproto/beerproto_go"
)

type Water struct {
	ID          string `csv:"ID"`
	Name        string `csv:"Name"`
	Calcium     *float64 `csv:"Calcium ion as Ca+2 (mg/L)"`
	Magnesium   *float64 `csv:"Magnesium ion as Mg+2  (mg/L)"`
	Sulfate     *float64 `csv:"Sulfate ion as SO42-  (mg/L)"`
	Sodium      *float64 `csv:"Sodium ions as Na+  (mg/L)"`
	Chloride    *float64 `csv:"Chloride as Cl- (mg/L)"`
	Bicarbonate *float64 `csv:"Bicarbonate as HCO3- (mg/L)"`
	Nitrite     *float64 `csv:"Nitrite as NO3"`
	Potassium   *float64 `csv:"Potassium as K (mg/L)"`
	Iron        *float64 `csv:"Iron (mg/L)"`
	Flouride    *float64 `csv:"Flouride as F"`
	Nitrate     *float64 `csv:"Nitrate as NO2"`
	Description string `csv:"Description"`
}

func (s Water) ToWaterBase() *beerproto.WaterBase {
	return &beerproto.WaterBase{
		Id: s.ID,
		Name: s.Name,
		Calcium: toConcentrationType(s.Calcium),
		Nitrite: toConcentrationType(s.Nitrite),
		Chloride: toConcentrationType(s.Chloride),
		Potassium: toConcentrationType(s.Potassium),
		Iron: toConcentrationType(s.Iron),
		Flouride: toConcentrationType(s.Flouride),
		Sulfate: toConcentrationType(s.Sulfate),
		Magnesium: toConcentrationType(s.Magnesium),
		Bicarbonate: toConcentrationType(s.Bicarbonate),
		Nitrate: toConcentrationType(s.Nitrate),
		Sodium: toConcentrationType(s.Sodium),
	}
}
