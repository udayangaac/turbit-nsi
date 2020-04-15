package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/udayangaac/turbit-nsi/internal/config"
	external_service "github.com/udayangaac/turbit-nsi/internal/external-service"
	"github.com/udayangaac/turbit-nsi/internal/external-service/elasticsearch"
	geo_classifier "github.com/udayangaac/turbit-nsi/internal/external-service/geo-classifier"
	log_traceable "github.com/udayangaac/turbit-nsi/internal/lib/log-traceable"
	"github.com/udayangaac/turbit-nsi/internal/lib/orm"
	"github.com/udayangaac/turbit-nsi/internal/service"
	"github.com/udayangaac/turbit-nsi/internal/transport/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ctx := context.Background()
	osSigChan := make(chan os.Signal, 0)
	signal.Notify(osSigChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGINT)

	//opts := badger.DefaultOptions(config.EnvConf.Badger.DbPath)
	//opts.Dir = config.EnvConf.Badger.DbPath
	//opts.ValueDir = config.EnvConf.Badger.DbPath
	//db, err := badger.Open(opts)
	//
	//if err != nil {
	//	log.Fatal(log_traceable.GetMessage(ctx, "Unable to open database"))
	//}

	config.InitConfigurations()
	if err := orm.InitDatabase(config.DatabaseConf); err != nil {
		log.Fatal(log_traceable.GetMessage(ctx, "Unable to open the database error :"+err.Error()))
	}

	//repos := repo.Container{
	//
	//}

	esOptns := elasticsearch.ClientOptions{
		Addresses: config.ElasticsearchConf.Addresses,
		Username:  config.ElasticsearchConf.Username,
		Password:  config.ElasticsearchConf.Password,
	}

	extServices := external_service.Container{
		ESConnector:   elasticsearch.NewConnectorImpl(esOptns),
		GeoClassifier: geo_classifier.NewGeoClassifier(config.GeoClassifierConf.BaseUrl),
	}

	// Inject external services
	services := service.Container{
		GatewayService: service.NewGatewayService(extServices),
	}

	// Inject services
	webService := http.WebService{
		Port:     config.EnvConf.Port,
		Services: services,
	}

	webService.Init()

	select {
	case <-osSigChan:
		webService.Stop()
		err := orm.CloseDatabase()
		if err != nil {
			log.Fatal(log_traceable.GetMessage(ctx, "Unable to open database ="))
		}
	}
}
