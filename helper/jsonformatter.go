package helper

type ResponseWithData struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseErrorWithData struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

type ResponseWithOutData struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type StructKosong struct {
}

func ApiResponse(status bool, message string, data interface{}) interface{} {
	if data == nil {
		var structkosong StructKosong
		jsonResponse := ResponseErrorWithData{
			Success: status,
			Message: message,
			Error:   structkosong,
		}
		return jsonResponse
	} else {
		jsonResponse := ResponseWithData{
			Success: status,
			Message: message,
			Data:    data,
		}
		return jsonResponse
	}

}

func ApiWithOutData(status bool, message string) interface{} {
	jsonResponse := ResponseWithOutData{
		Success: status,
		Message: message,
	}
	return jsonResponse
}
