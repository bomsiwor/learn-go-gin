package response

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Meta Meta `json:"meta"`
	Data any  `json:"data"`
}

func GenerateResponse(code int, message string, data any) Response {
	response := Response{
		Meta: Meta{
			Code:    code,
			Message: message,
		},
		Data: data,
	}

	return response
}
