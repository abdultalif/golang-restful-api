package web

type WebResponseFailed struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Error   interface{} `json:"error"`
}
