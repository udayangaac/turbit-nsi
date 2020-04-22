package nsi_client

import (
	"context"
	"testing"
)

func TestNewNSIConnector(t *testing.T) {
	connector := NewNSIConnector("http://134.122.123.88:3001")
	//{
	//	"lat": "6.814360",
	//	"lon": "81.059219",
	//	"geo_ref_id":"0x886102cde9fffff",
	//	"user_id":3,
	//	"categories":["food"],
	//	"is_newest": false
	//}
	req := RequestBody{
		Lat:        "6.814360",
		Lon:        "81.059219",
		GeoRefID:   "0x886102cde9fffff",
		UserID:     3,
		Categories: []string{"food"},
		IsNewest:   false,
	}
	notifications, gid, err := connector.GetNotifications(context.Background(), req)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(notifications, gid)
}
