package response

type ErrorResp struct {
	Code    int    `json:"code"`
	Err     error  `json:"error"`
	Message string `json:"message"`
}
