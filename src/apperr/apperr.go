package apperr

const BAD_REQUEST = "BAD_REQUEST"
const INTERNAL_ERROR = "INTERNAL_SERVER_ERROR"

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
