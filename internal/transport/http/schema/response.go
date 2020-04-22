package schema

type SuccessResp struct {
	Data interface{} `json:"data"`
}

type ErrorResp struct {
	Message string `json:"message"`
}

type SuccessMessage struct {
	Message string `json:"message"`
}

type UserNotifications struct {
	GeoRefId      string      `json:"geo_ref_id"`
	Notifications interface{} `json:"notification"`
}
