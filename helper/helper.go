package helper

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Code    int    `json:"code"`
}

func APIResponse(message string, status string, code int, data interface{}) Response {
	meta := Meta{
		Message: message,
		Status:  status,
		Code:    code,
	}
	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}
	return jsonResponse
}
