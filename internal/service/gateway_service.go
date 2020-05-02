package service

import (
	"context"
)

type Notifications struct {
	//Offset    int                      `json:"offset"`
	RefId     string              `json:"ref_id"`
	Documents []FormattedDocument `json:"documents"`
}

type Param struct {
	Lat, Lon       string
	GeoRefId       string
	UserId         int
	IsOffsetEnable bool
	Categories     []string
}

type FormattedDocument struct {
	Id               int64    `json:"id"`
	CompanyName      string   `json:"company_name"`
	Content          string   `json:"content"`
	NotificationType int      `json:"notification_type"`
	StartTime        string   `json:"start_time"`
	EndDate          string   `json:"end_date"`
	LogoCompany      string   `json:"logo_company"`
	ImagePublisher   string   `json:"image_publisher"`
	Categories       []string `json:"categories"`
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
}

type Location struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

type GatewayService interface {
	Add(ctx context.Context, document Document) (err error)
	Update(ctx context.Context, document Document) (err error)
	GetNotifications(ctx context.Context, param Param) (notifications Notifications, err error)
}
