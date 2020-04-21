package service

import (
	"context"
	external_service "github.com/udayangaac/turbit-nsi/internal/external-service"
	"github.com/udayangaac/turbit-nsi/internal/external-service/elasticsearch"
	geo_classifier "github.com/udayangaac/turbit-nsi/internal/external-service/geo-classifier"
)

type gatewayService struct {
	ExtServiceContainer external_service.Container
}

func NewGatewayService(extServiceContainer external_service.Container) GatewayService {
	return &gatewayService{
		ExtServiceContainer: extServiceContainer,
	}
}

func (g *gatewayService) Add(ctx context.Context, document Document) (err error) {

	// Get hex ids from hex ids
	hexIds := make([]string, 0)
	for _, val := range document.Locations {
		geoRecord := geo_classifier.GeoRecordDetail{
			Latitude:  val.Lat,
			Longitude: val.Lon,
			Offset:    document.Id,
		}
		if record, errExt := g.ExtServiceContainer.GeoClassifier.AddRecord(ctx, geoRecord); errExt != nil {
		} else {
			hexIds = append(hexIds, record.Data.GeoHexID)
		}
	}

	// Get locations
	locations := []elasticsearch.Location{}
	for _, val := range document.Locations {
		l := elasticsearch.Location{}
		l.Lon = val.Lon
		l.Lat = val.Lat
		locations = append(locations, l)
	}

	// Add document to elastic search
	doc := elasticsearch.Document{
		Id:               document.Id,
		CompanyName:      document.CompanyName,
		Content:          document.Content,
		NotificationType: document.NotificationType,
		StartTime:        document.StartTime,
		EndDate:          document.EndDate,
		LogoCompany:      document.LogoCompany,
		ImagePublisher:   document.ImagePublisher,
		Category:         document.Category,
		GeoHexIds:        hexIds,
		Locations:        locations,
	}
	return g.ExtServiceContainer.ESConnector.AddDocument(ctx, "active_notifications_index", doc)
}

func (g *gatewayService) Update(ctx context.Context, document Document) (err error) {
	// Get hex ids from hex ids
	hexIds := make([]string, 0)
	for _, val := range document.Locations {
		geoRecord := geo_classifier.GeoRecordDetail{
			Latitude:  val.Lat,
			Longitude: val.Lon,
			Offset:    document.Id,
		}
		if record, errExt := g.ExtServiceContainer.GeoClassifier.AddRecord(ctx, geoRecord); errExt != nil {
		} else {
			hexIds = append(hexIds, record.Data.GeoHexID)
		}
	}

	// Get locations
	locations := []elasticsearch.Location{}
	for _, val := range document.Locations {
		l := elasticsearch.Location{}
		l.Lon = val.Lon
		l.Lat = val.Lat
		locations = append(locations, l)
	}

	// Add document to elastic search
	doc := elasticsearch.Document{
		Id:               document.Id,
		CompanyName:      document.CompanyName,
		Content:          document.Content,
		NotificationType: document.NotificationType,
		StartTime:        document.StartTime,
		EndDate:          document.EndDate,
		LogoCompany:      document.LogoCompany,
		ImagePublisher:   document.ImagePublisher,
		Category:         document.Category,
		GeoHexIds:        hexIds,
		Locations:        locations,
	}
	return g.ExtServiceContainer.ESConnector.AddDocument(ctx, "active_notifications_index", doc)
}

func (g *gatewayService) GetNotifications(ctx context.Context, param Param) (notifications Notifications, err error) {

	notifications = Notifications{}
	summery := geo_classifier.LocationSummery{}
	userDetails := geo_classifier.UserDetail{
		Latitude:  param.Lon,
		Longitude: param.Lon,
		Offset:    0,
		UserId:    param.UserId,
	}
	if summery, err = g.ExtServiceContainer.GeoClassifier.GetLocationSummery(ctx, userDetails); err != nil {
		return
	}

	// geo reference generator
	criteria := elasticsearch.Criteria{
		Index:          elasticsearch.ActiveNotificationsIndex,
		GeoHexId:       []string{summery.Data.GeoRef},
		LastConsumedId: int64(summery.Data.CurrentOffset),
		Category:       "",
		PageIndex:      0,
		PageSize:       10,
	}
	notifications.Documents, err = g.ExtServiceContainer.ESConnector.GetDocuments(ctx, criteria)
	return
}
