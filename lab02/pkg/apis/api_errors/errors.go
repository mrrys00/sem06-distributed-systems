package api_errors

type ErrorPage struct {
	StatusCode int    `json:"status_code"`
	Err        string `json:"message"`
}

func NewErrorPage(status int, err string) *ErrorPage {
	return &ErrorPage{
		StatusCode: status,
		Err:        err,
	}
}
