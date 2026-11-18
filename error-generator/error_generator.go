package error_generator

import (
	"fmt"
	"time"
)

type ErrorMessageRequest struct {
	AppName    string
	HTTPStatus int
}

type ErrorMessageResult struct {
	Error string `json:"error"`
	Code  string `json:"code"`
}

// NewClient membuat instance HTTP Client dengan default timeout
func NewErrorGenerator(req ErrorMessageRequest) ErrorMessageResult {

	return ErrorMessageResult{
		Error: GetHTTPErrorMessageID(req.HTTPStatus),
		Code:  fmt.Sprintf("[%s] %s", req.AppName, time.Now().Format("20060102150405")),
	}
}
