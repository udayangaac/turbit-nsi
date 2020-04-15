package geo_classifier

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

// Integration test
// NOTE : Need to neglect from CI config
func TestNewGeoClassifier(t *testing.T) {
	t.Run("adding a record", TestGeoClassifier_AddRecord)
	t.Run("get location summery", TestGeoClassifier_GetLocationSummery)
}

func TestGeoClassifier_GetLocationSummery(t *testing.T) {
	gc := NewGeoClassifier("http://localhost:8001")
	detail := UserDetail{
		UserId:    2020,
		Latitude:  "6.714360",
		Longitude: "81.059219",
		Offset:    0,
	}
	data, err1 := gc.GetLocationSummery(context.Background(), detail)
	if err1 != nil {
		t.Log(err1.Error())
	}
	byteArr, err2 := json.MarshalIndent(data, "", "\t")
	if err2 != nil {
		t.Log(err2.Error())
	}
	t.Log(fmt.Sprintf("Success fully create the request %s", byteArr))
}

func TestGeoClassifier_AddRecord(t *testing.T) {
	gc := NewGeoClassifier("http://localhost:8001")
	detail := GeoRecordDetail{
		Latitude:  "6.714360",
		Longitude: "81.059219",
		Offset:    121,
	}
	data, err1 := gc.AddRecord(context.Background(), detail)
	if err1 != nil {
		t.Log(err1.Error())
	}
	byteArr, err2 := json.MarshalIndent(data, "", "\t")
	if err2 != nil {
		t.Log(err2.Error())
	}
	t.Log(fmt.Sprintf("Success fully create the request %s", byteArr))
}
