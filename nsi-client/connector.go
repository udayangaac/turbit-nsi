package nsi_client

import (
	"context"
	"time"
)

type ResponseBody struct {
	Data struct {
		GeoRefID      string         `json:"geo_ref_id"`
		Notifications []Notification `json:"notifications"`
	} `json:"data"`
}

type RequestBody struct {
	Lat        string   `json:"lat"`
	Lon        string   `json:"lon"`
	GeoRefID   string   `json:"geo_ref_id"`
	UserID     int      `json:"user_id"`
	Categories []string `json:"categories"`
	IsNewest   bool     `json:"is_newest"`
}

type Notification struct {
	ID               int       `json:"id"`
	CompanyName      string    `json:"company_name"`
	Content          string    `json:"content"`
	NotificationType int       `json:"notification_type"`
	StartTime        time.Time `json:"start_time"`
	EndDate          time.Time `json:"end_date"`
	LogoCompany      string    `json:"logo_company"`
	ImagePublisher   string    `json:"image_publisher"`
	Categories       []string  `json:"categories"`
	Locations        []struct {
		Lat string `json:"lat"`
		Lon string `json:"lon"`
	} `json:"locations"`
	GeoHexIds []string `json:"geo_hex_ids"`
}

type NSIConnector interface {
	GetNotifications(ctx context.Context, param RequestBody) (notifications []Notification, geoRefId string, err error)
}
