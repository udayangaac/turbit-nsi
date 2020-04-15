package external_service

import "github.com/udayangaac/turbit-nsi/internal/external-service/elasticsearch"

type Container struct {
	ESConnector elasticsearch.Connector
}
