package schema

//{
//	"id":1,
//	"company_name":"pizza hut",
//	"content":"one to one large pizza offer",
//	"notification_type":1,
//	"start_time":"2020-04-14 10:00:00",
//	"end_date":"2020-04-14 20:00:00",
//	"logo_company":"",
//	"image_publisher":"",
//	"category" : "lat"
//	"location":[
//		{
//		"lat":"6.948676",
//		"lon":"79.859658"
//		}]
//}

type NotificationRequest struct {
	ID               int64  `json:"id"`
	CompanyName      string `json:"company_name"`
	Content          string `json:"content"`
	NotificationType int    `json:"notification_type"`
	StartTime        string `json:"start_time"`
	EndDate          string `json:"end_date"`
	LogoCompany      string `json:"logo_company"`
	ImagePublisher   string `json:"image_publisher"`
	Category         string `json:"category"`
	Locations        []struct {
		Lat string `json:"lat"`
		Lon string `json:"lon"`
	} `json:"locations"`
}

type GetNotificationRequest struct {
	Lat      string `json:"lat"`
	Lon      string `json:"lon"`
	GeoRefId string `json:"geo_ref_id"`
	UserId   int    `json:"user_id"`
}
