package web

type WebResponse struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
