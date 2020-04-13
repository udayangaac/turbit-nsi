package schema

type SuccessResp struct {
	Data interface{} `json:"data"`
}

type ErrorResp struct {
	Message string `json:"message"`
}
