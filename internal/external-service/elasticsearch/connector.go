// Copyright 2020. All rights reserved.
// Author : Chamith Udayanga.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package elasticsearch

import "context"

type ClientOptions struct {
	Addresses []string
	Username  string
	Password  string
}

type Document struct {
	Id               int64      `json:"id"`
	CompanyName      string     `json:"company_name"`
	Content          string     `json:"content"`
	NotificationType int        `json:"notification_type"`
	StartTime        string     `json:"start_time"`
	EndDate          string     `json:"end_date"`
	LogoCompany      string     `json:"logo_company"`
	ImagePublisher   string     `json:"image_publisher"`
	Categories       []string   `json:"categories"`
	Locations        []Location `json:"locations"`
	GeoHexIds        []string   `json:"geo_hex_ids"`
}

type Location struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

type Criteria struct {
	Index            string
	GeoHexId         []string
	LastConsumedId   int64
	Categories       []string
	PageIndex        int
	PageSize         int
	TextSearchEnable bool
	SearchTerm       string
}

const ActiveNotificationsIndex = "active_notifications_index"

type Connector interface {
	AddDocument(ctx context.Context, index string, doc Document) (err error)
	GetDocuments(ctx context.Context, criteria Criteria) (docs []Document, err error)
	DeleteDocument(ctx context.Context, id int64) (err error)
}
