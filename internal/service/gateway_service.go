package service

import (
	"context"
	"github.com/udayangaac/turbit-nsi/internal/external-service/elasticsearch"
)

type Notifications struct {
	Offset    int64
	RefId     string
	Documents []elasticsearch.Document
}

type Param struct {
	Lat, Lon float64
	Offset   int64
	RefId    string
}

type GatewayService interface {
	Add(ctx context.Context, document elasticsearch.Document) (err error)
	Update(ctx context.Context, document elasticsearch.Document) (err error)
	GetNotifications(ctx context.Context, param Param) (notifications Notifications, err error)
}
