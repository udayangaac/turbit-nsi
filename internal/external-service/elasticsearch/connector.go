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
	Logger    Logger
}

type Document struct {
	Id               int64  `json:"id"`
	CompanyName      string `json:"company_name"`
	Content          string `json:"content"`
	NotificationType int    `json:"notification_type"`
	StartTime        string `json:"start_time"`
	EndDate          string `json:"end_date"`
	LogoCompany      string `json:"logo_company"`
	ImagePublisher   string `json:"image_publisher"`
	Locations        []struct {
		Lat      string `json:"lat"`
		Lon      string `json:"lon"`
		GeoHexId string `json:"geo_hex_id"`
	} `json:"locations"`
}

type Connector interface {
	AddDocument(ctx context.Context, index string, doc Document) (err error)
}
