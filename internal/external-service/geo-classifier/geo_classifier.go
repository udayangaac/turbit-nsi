package geo_classifier

import "context"

type UserDetail struct {
	UserId    int
	Latitude  string
	Longitude string
	Offset    int
}

type LocationSummery struct {
	Data struct {
		OldOffset     int    `json:"old_offset"`
		CurrentOffset int    `json:"current_offset"`
		GeoRef        string `json:"geo_ref"`
	} `json:"data"`
}

type GeoRecordDetail struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Offset    int64  `json:"offset"`
}

type Record struct {
	Data struct {
		GeoHexID string `json:"geo_hex_id"`
		Offset   int    `json:"offset"`
	} `json:"data"`
}

type GeoClassifier interface {
	GetLocationSummery(ctx context.Context, detail UserDetail) (summery LocationSummery, err error)
	AddRecord(ctx context.Context, detail GeoRecordDetail) (record Record, err error)
}
