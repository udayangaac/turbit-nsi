package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/udayangaac/turbit-nsi/internal/config"
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
	// Inject repositories
	services := service.Container{}

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
