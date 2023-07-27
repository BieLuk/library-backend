package apperr

const BAD_REQUEST = "BAD_REQUEST"

type AppErr struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewAppErr(code, message string) *AppErr {
	return &AppErr{
		Code:    code,
		Message: message,
	}
}
