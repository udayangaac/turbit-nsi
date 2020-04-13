package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	log_traceable "github.com/udayangaac/turbit-nsi/internal/lib/log-traceable"
	"github.com/udayangaac/turbit-nsi/internal/service"
	"github.com/udayangaac/turbit-nsi/internal/transport/http/schema"
	"net/http"
	"strconv"
	"time"
)

type WebService struct {
	Port     int
	Services service.Container
	server   *http.Server
}

func (w *WebService) Init() {
	rootRouter := mux.NewRouter()
	tnsiRouter := rootRouter.PathPrefix("/tnsi").Subrouter()

	log.Info(log_traceable.GetMessage(context.Background(), "Server is starting, Port:"+fmt.Sprintf("%v", w.Port)))
	w.server = &http.Server{
		Addr:         fmt.Sprintf(":%v", w.Port),
		WriteTimeout: time.Second * 1,
		ReadTimeout:  time.Second * 1,
		IdleTimeout:  time.Second * 1,
		Handler:      rootRouter,
	}
	go func() {
		if err := w.server.ListenAndServe(); err != nil {
			log.Error(log_traceable.GetMessage(context.Background(), "Unable to start the server error.", err.Error()))
		}
	}()

}

func (w *WebService) Stop() {
	err := w.server.Shutdown(context.Background())
	if err != nil {
		log.Error(log_traceable.GetMessage(context.Background(), "Error shutting down application error.", err.Error()))
	}
}

func getTraceableContext(req *http.Request) (ctx context.Context) {
	uuidStr := uuid.New().String()
	ctx = context.WithValue(req.Context(), "uuid_str", uuidStr)
	log.Trace(log_traceable.GetMessage(ctx, "Started to process request URL:", req.URL, "Method:", req.Method))
	return
}
