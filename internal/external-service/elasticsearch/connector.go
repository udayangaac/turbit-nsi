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

type UserActionDocument struct {
	Id             string `json:"id"`
	UserId         int64  `json:"user_id"`
	NotificationId int64  `json:"notification_id"`
	UserReaction   int16  `json:"user_reaction"`
	IsViewed       bool   `json:"is_viewed"`
}

type Location struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

type Criteria struct {
	UserId           int
	GeoHexId         []string
	LastConsumedId   int64
	Categories       []string
	PageIndex        int
	PageSize         int
	TextSearchEnable bool
	SearchTerm       string
}

const (
	ActiveNotificationsIndex = "active_notifications_index"
	UserActionsIndex         = "user_actions_index"
)

type Connector interface {
	AddDocument(ctx context.Context, doc Document) (err error)
	GetDocuments(ctx context.Context, criteria Criteria) (docs []Document, err error)
	DeleteDocument(ctx context.Context, id int64) (err error)
	AddUserActionDocument(ctx context.Context, doc UserActionDocument) (err error)
	GetUserActionDocuments(ctx context.Context, criteria Criteria) (docs map[int64]UserActionDocument, err error)
}
