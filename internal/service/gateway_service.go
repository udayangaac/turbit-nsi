package service

import (
	"context"
	"github.com/udayangaac/turbit-nsi/internal/external-service/elasticsearch"
)

type Notifications struct {
	Offset    int
	RefId     string
	Documents []elasticsearch.Document
}

type Param struct {
	Lat, Lon string
	GeoRefId string
	UserId   int
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
	Category         string     `json:"category"`
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
