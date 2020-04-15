package service

import (
	"context"
	external_service "github.com/udayangaac/turbit-nsi/internal/external-service"
	"github.com/udayangaac/turbit-nsi/internal/external-service/elasticsearch"
)

type gatewayService struct {
	ExtServiceContainer external_service.Container
}

func NewGatewayService(extServiceContainer external_service.Container) GatewayService {
	return &gatewayService{
		ExtServiceContainer: extServiceContainer,
	}
}

func (g *gatewayService) Add(ctx context.Context, document elasticsearch.Document) (err error) {

}

func (g *gatewayService) Update(ctx context.Context, document elasticsearch.Document) (err error) {

}

func (g *gatewayService) GetNotifications(ctx context.Context, param Param) (notifications Notifications, err error) {

}
