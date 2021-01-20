package respond

type Response struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func Builder(status bool, msg string, data interface{}) Response {
	return Response{
		Status: status,
		Message: msg,
		Data: data,
	}
}
