package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	log_traceable "github.com/udayangaac/turbit-nsi/internal/lib/log-traceable"
	"github.com/udayangaac/turbit-nsi/internal/service"
	"github.com/udayangaac/turbit-nsi/internal/transport/http/schema"
)

type WebService struct {
	Port     int
	Services service.Container
	server   *http.Server
}

func (w *WebService) Init() {
	rootRouter := mux.NewRouter()
	tnsiRouter := rootRouter.PathPrefix("/tnsi").Subrouter()

	tnsiRouter.HandleFunc("/notification",
		AddNotificationHandler(w.Services)).Methods(http.MethodPost)

	tnsiRouter.HandleFunc("/notification/{notificationId}",
		ModifyNotificationHandler(w.Services)).Methods(http.MethodPut)

	tnsiRouter.HandleFunc("/notifications",
		GetNotificationsHandler(w.Services)).Methods(http.MethodPost)

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

func AddNotificationHandler(services service.Container) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		ctx := getTraceableContext(request)
		req := schema.NotificationRequest{}
		decoder := json.NewDecoder(request.Body)
		err := decoder.Decode(&req)
		if err != nil {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusBadRequest)
			if err := json.NewEncoder(writer).Encode(schema.ErrorResp{Message: "invalid user id in path variable"}); err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				return
			} else {
				return
			}
		}

		locations := []service.Location{}

		for _, val := range req.Locations {
			l := service.Location{}
			l.Lon = val.Lon
			l.Lat = val.Lat
			locations = append(locations, l)
		}

		doc := service.Document{
			Id:               req.ID,
			CompanyName:      req.CompanyName,
			Content:          req.Content,
			NotificationType: req.NotificationType,
			StartTime:        req.StartTime,
			EndDate:          req.EndDate,
			LogoCompany:      req.LogoCompany,
			ImagePublisher:   req.ImagePublisher,
			Categories:       req.Categories,
			Locations:        locations,
		}
		docStr, _ := json.MarshalIndent(doc, "", "\t")
		fmt.Printf("Added document request details \n%s", docStr)

		err = services.GatewayService.Add(ctx, doc)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)

			msg := schema.SuccessMessage{
				Message: "Add notification successfully !",
			}
			if err = json.NewEncoder(writer).Encode(schema.SuccessResp{Data: msg}); err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				return
			} else {
				return
			}
		}
	}
}

func ModifyNotificationHandler(services service.Container) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		ctx := getTraceableContext(request)

		req := schema.NotificationRequest{}
		decoder := json.NewDecoder(request.Body)
		err := decoder.Decode(&req)
		if err != nil {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusBadRequest)
			if err := json.NewEncoder(writer).Encode(schema.ErrorResp{Message: "invalid request"}); err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				return
			} else {
				return
			}
		}

		vars := mux.Vars(request)
		if vars == nil {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusBadRequest)
			if err := json.NewEncoder(writer).Encode(schema.ErrorResp{Message: "invalid request"}); err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				return
			} else {
				return
			}
		}

		req.ID, err = strconv.ParseInt(vars["notificationId"], 10, 32)
		if err != nil {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusBadRequest)
			if err := json.NewEncoder(writer).Encode(schema.ErrorResp{Message: "invalid request"}); err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				return
			} else {
				return
			}
		}

		locations := []service.Location{}

		for _, val := range req.Locations {
			l := service.Location{}
			l.Lon = val.Lon
			l.Lat = val.Lat
			locations = append(locations, l)
		}

		doc := service.Document{
			CompanyName:      req.CompanyName,
			Content:          req.Content,
			NotificationType: req.NotificationType,
			StartTime:        req.StartTime,
			EndDate:          req.EndDate,
			LogoCompany:      req.LogoCompany,
			ImagePublisher:   req.ImagePublisher,
			Categories:       req.Categories,
			Locations:        locations,
		}
		err = services.GatewayService.Update(ctx, doc)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)

			msg := schema.SuccessMessage{
				Message: "Modified notification successfully !",
			}
			if err = json.NewEncoder(writer).Encode(schema.SuccessResp{Data: msg}); err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				return
			} else {
				return
			}
		}
	}
}

func GetNotificationsHandler(services service.Container) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		ctx := getTraceableContext(request)

		req := schema.GetNotificationRequest{}

		decoder := json.NewDecoder(request.Body)
		err := decoder.Decode(&req)
		if err != nil {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(writer).Encode(schema.ErrorResp{Message: "invalid request"}); err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				return
			} else {
				return
			}
		}

		param := service.Param{
			Lat:            req.Lat,
			Lon:            req.Lon,
			GeoRefId:       req.GeoRefId,
			UserId:         req.UserId,
			IsOffsetEnable: req.IsNewest,
			Categories:     req.Categories,
			SearchText:     req.SearchTerm,
		}

		//// latitude
		//param.Lat, err = strconv.ParseFloat(req.Lat, 64)
		//if err != nil {
		//	writer.Header().Set("Content-Type", "application/json")
		//	writer.WriteHeader(http.StatusBadRequest)
		//	if err := json.NewEncoder(writer).Encode(schema.ErrorResp{Message: "invalid latitude in path variable"}); err != nil {
		//		writer.WriteHeader(http.StatusInternalServerError)
		//		return
		//	} else {
		//		return
		//	}
		//}
		//
		//// latitude
		//param.Lon, err = strconv.ParseFloat(req.Lon, 64)
		//if err != nil {
		//	writer.Header().Set("Content-Type", "application/json")
		//	writer.WriteHeader(http.StatusBadRequest)
		//	if err := json.NewEncoder(writer).Encode(schema.ErrorResp{Message: "invalid longitude in path variable"}); err != nil {
		//		writer.WriteHeader(http.StatusInternalServerError)
		//		return
		//	} else {
		//		return
		//	}
		//}

		notifications := service.Notifications{}
		notifications, err = services.GatewayService.GetNotifications(ctx, param)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)

			data := schema.UserNotifications{
				GeoRefId:      notifications.RefId,
				Notifications: notifications.Documents,
			}

			if err = json.NewEncoder(writer).Encode(schema.SuccessResp{Data: data}); err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				return
			} else {
				return
			}
		}
	}
}

func getTraceableContext(req *http.Request) (ctx context.Context) {
	uuidStr := uuid.New().String()
	ctx = context.WithValue(req.Context(), "uuid_str", uuidStr)
	log.Trace(log_traceable.GetMessage(ctx, "Started to process request URL:", req.URL, "Method:", req.Method))
	return
}
