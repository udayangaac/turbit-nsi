// Copyright 2020. All rights reserved.
// Author : Chamith Udayanga.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package elasticsearch

import (
	"context"
	"testing"
)

func TestConnector_AddDocument(t *testing.T) {
	otns := ClientOptions{
		Addresses: []string{"http://localhost:9200"},
		Username:  "username",
		Password:  "password",
	}
	conn := NewConnectorImpl(otns)

	l1 := Location{
		Lat: "1",
		Lon: "2",
	}

	geoHexIds := []string{"a"}

	doc1 := Document{
		Id:               1,
		CompanyName:      "Turbit",
		Content:          "This is a test document",
		NotificationType: 1,
		StartTime:        "",
		EndDate:          "",
		LogoCompany:      "",
		ImagePublisher:   "",
		Categories:       []string{"a"},
		GeoHexIds:        geoHexIds,
	}

	doc2 := Document{
		Id:               2,
		CompanyName:      "Turbit ",
		Content:          "This is a test document",
		NotificationType: 1,
		StartTime:        "",
		EndDate:          "",
		LogoCompany:      "",
		ImagePublisher:   "",
		Categories:       []string{"a"},
		GeoHexIds:        geoHexIds,
	}

	doc1.Locations = append(doc1.Locations, l1)
	doc2.Locations = append(doc2.Locations, l1)

	err := conn.AddDocument(context.Background(), "active_notifications_index", doc1)
	if err != nil {
		t.Error(err.Error())
	}

	err = conn.AddDocument(context.Background(), "active_notifications_index", doc2)
	if err != nil {
		t.Error(err.Error())
	}

	ds := []Document{}
	ds, err = conn.GetDocuments(context.Background(), Criteria{
		Index:          "active_notifications_index",
		GeoHexId:       []string{"a"},
		LastConsumedId: 1,
		Categories:     []string{"a"},
		PageIndex:      0,
		PageSize:       10,
	})

	if err != nil {
		t.Error(err.Error())
	}
	t.Log(ds)
}
