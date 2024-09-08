package model

type Response struct {
	Success bool   `json:"success"`
	Data    any    `json:"data"`
	Msg     string `json:"msg"`
}

func Success(data any) Response {
	return Response{
		Success: true,
		Data:    data,
		Msg:     "success",
	}
}

func Failed(msg string) Response {
	return Response{
		Success: false,
		Data:    "",
		Msg:     msg,
	}
}
