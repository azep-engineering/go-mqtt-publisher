package response

type Response struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func getResponse(errorBool bool, messageKey string, data interface{}) (res Response) {
	response := Response{
		Error:   errorBool,
		Message: messageKey,
		Data:    data,
	}
	return response
}
