package geo_classifier

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	log_traceable "github.com/udayangaac/turbit-nsi/internal/lib/log-traceable"
	"net/http"
	"time"
)

type geoClassifier struct {
	BaseUrl string
	Client  *http.Client
}

func NewGeoClassifier(baseUrl string) GeoClassifier {
	g := geoClassifier{
		BaseUrl: baseUrl,
	}
	g.init()
	return &g
}

func (g *geoClassifier) init() {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	g.Client = &http.Client{Transport: tr}
}

func (g *geoClassifier) GetLocationSummery(ctx context.Context, detail UserDetail) (summery LocationSummery, err error) {
	var (
		req *http.Request
		res *http.Response
	)
	url := fmt.Sprintf(
		"%s/tgc/location-summery/user/%v/loc/%v/%v?documentOffset=%v",
		g.BaseUrl,
		detail.UserId,
		detail.Latitude,
		detail.Longitude,
		detail.Offset,
	)
	if req, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		return
	}
	if g.Client == nil {
		g.init()
		log.Warn(log_traceable.GetMessage(ctx, "Re-initialized the get location summery"))
	}

	if res, err = g.Client.Do(req); err != nil {
		return
	}
	defer func() {
		err = res.Body.Close()
		if err != nil {
			log.Error(log_traceable.GetMessage(ctx, "Unable close the body of the request"))
		}
	}()
	err = json.NewDecoder(res.Body).Decode(&summery)
	return
}

func (g *geoClassifier) AddRecord(ctx context.Context, detail GeoRecordDetail) (record Record, err error) {
	var (
		req     *http.Request
		res     *http.Response
		payload []byte
	)

	url := fmt.Sprintf("%s/tgc/geo/record", g.BaseUrl)
	if payload, err = json.Marshal(detail); err != nil {
		return
	}
	if req, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload)); err != nil {
		return
	}
	req.Header.Add("content-type", "application/json")
	if res, err = g.Client.Do(req); err != nil {
		return
	}
	defer func() {
		err = res.Body.Close()
		if err != nil {
			log.Error(log_traceable.GetMessage(ctx, "Unable close the body of the request"))
		}
	}()
	err = json.NewDecoder(res.Body).Decode(&record)
	return
}
