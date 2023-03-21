package core

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func ResponseOK() Message {
	return Message{
		Status:  "OK",
		Message: "",
	}
}

func ResponseError(err error) Message {
	return Message{
		Status:  "error",
		Message: err.Error(),
	}
}

func ResponseData(data []byte) Message {
	return Message{
		Status:  "OK",
		Message: string(data),
	}
}
