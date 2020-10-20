package csv

import (
	beerproto "github.com/beerproto/beerproto_go"
	"strings"
)

type Packaging struct {
	ID          string `csv:"ID"`
	Name        string `csv:"Name"`
	Type        string `csv:"Type"`
	Volume      *float64 `csv:"Volume (ml)"`
	Description string `csv:"Description"`
}

func (s Packaging) ToPackagingProcedureType() *beerproto.PackagingProcedureType {
	return &beerproto.PackagingProcedureType{
		Id:s.ID,
		Name: s.Name,
		PackagingVessels: []*beerproto.PackagingVesselType{
			&beerproto.PackagingVesselType{
				Id: s.ID,
				Name: s.Name,
				Description: s.Description,
				VesselQuantity: 1,
				Type: s.ToPackagingVesselType(),
				VesselVolume: toVolumeType(s.Volume, beerproto.VolumeType_ML),
			},
		},
	}
}


func (s *Packaging) ToPackagingVesselType() beerproto.PackagingVesselType_PackagingVesselTypeType {
	if v, ok := beerproto.PackagingVesselType_PackagingVesselTypeType_value[strings.ToUpper(s.Type)]; ok {
		return beerproto.PackagingVesselType_PackagingVesselTypeType(v)
	}
	return beerproto.PackagingVesselType_NULL_PACKAGINGVESSELTYPETYPE

}
