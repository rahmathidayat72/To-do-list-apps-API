package helper

type ResponseWeb struct {
	Code    int         `json:"code"`
	Massage string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func WebResponse(code int, message string, data interface{}) ResponseWeb {
	return ResponseWeb{
		Code:    code,
		Massage: message,
		Data:    data,
	}
}
