package external_service

import (
	"github.com/udayangaac/turbit-nsi/internal/external-service/elasticsearch"
	geo_classifier "github.com/udayangaac/turbit-nsi/internal/external-service/geo-classifier"
)

type Container struct {
	ESConnector   elasticsearch.Connector
	GeoClassifier geo_classifier.GeoClassifier
}
