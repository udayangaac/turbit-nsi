package service

import (
	"context"
	"errors"
	external_service "github.com/udayangaac/turbit-nsi/internal/external-service"
	"github.com/udayangaac/turbit-nsi/internal/external-service/elasticsearch"
	geo_classifier "github.com/udayangaac/turbit-nsi/internal/external-service/geo-classifier"
	"strings"
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
		Categories:       document.Categories,
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
		Categories:       document.Categories,
		GeoHexIds:        hexIds,
		Locations:        locations,
	}
	return g.ExtServiceContainer.ESConnector.AddDocument(ctx, "active_notifications_index", doc)
}

func (g *gatewayService) GetNotifications(ctx context.Context, param Param) (notifications Notifications, err error) {

	geoHexId := ""
	summery := geo_classifier.LocationSummery{}

	notifications = Notifications{
		// Offset:    -1,
		RefId:     "",
		Documents: make([]FormattedDocument, 0),
	}

	userDetails := geo_classifier.UserDetail{
		Latitude:  param.Lat,
		Longitude: param.Lon,
		Offset:    0,
		UserId:    param.UserId,
	}

	if summery, err = g.ExtServiceContainer.GeoClassifier.GetLocationSummery(ctx, userDetails); err != nil {
		return // err
	}

	if summery.Data.CurrentOffset == -1 && param.IsOffsetEnable {
		notifications.RefId = summery.Data.GeoRef
		return
	}

	if geoHexId, err = getGeoHexId(summery.Data.GeoRef); err != nil {
		return // err
	}

	// geo reference generator
	criteria := elasticsearch.Criteria{
		Index:          elasticsearch.ActiveNotificationsIndex,
		GeoHexId:       []string{geoHexId},
		LastConsumedId: int64(summery.Data.CurrentOffset),
		Categories:     param.Categories,
		PageIndex:      0,
		PageSize:       20,
	}

	formattedDocuments := make([]FormattedDocument, 0)
	documents := make([]elasticsearch.Document, 0)
	documents, err = g.ExtServiceContainer.ESConnector.GetDocuments(ctx, criteria)
	if err != nil {
		return
	}

	for _, v := range documents {
		formattedDocuments = append(formattedDocuments, FormattedDocument{
			Id:               v.Id,
			CompanyName:      v.CompanyName,
			Content:          v.Content,
			NotificationType: v.NotificationType,
			StartTime:        v.StartTime,
			EndDate:          v.EndDate,
			LogoCompany:      v.LogoCompany,
			ImagePublisher:   v.ImagePublisher,
			Categories:       v.Categories,
		})
	}

	// notifications.Offset = summery.Data.CurrentOffset
	// Geo reference id
	notifications.Documents = formattedDocuments
	notifications.RefId = summery.Data.GeoRef
	return
}

func getGeoHexId(userGeoRef string) (geoRef string, err error) {
	arr := strings.Split(userGeoRef, "_")
	if arr == nil {
		err = errors.New("invalid user geo reference id")
		return
	}
	if len(arr) != 2 {
		err = errors.New("invalid user geo reference id")
		return
	}
	return arr[0], nil
}
