package response

import (
	"errors"
	"net/http"
)

var (
	ErrNoUserArg = errors.New("username argument missing")
	ErrNoUser    = errors.New("username does not exist")
)

type Response struct {
	StatusCode int                    `json:"statusCode,omitempty"`
	Headers    map[string]string      `json:"headers,omitempty"`
	Body       map[string]interface{} `json:"body,omitempty"`
}

func NewResponse(statusCode int, body map[string]interface{}) *Response {
	return &Response{
		StatusCode: statusCode,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       body,
	}
}

func OkResponse() *Response {
	return NewResponse(http.StatusOK, map[string]interface{}{"message": "OK", "success": true})
}

func OkJsonReponse(body map[string]interface{}) *Response {
	return NewResponse(http.StatusOK, body)
}

// only first provided arg will be used. (OPTIONAL)
func OkMsgResponse(msg_optional ...string) *Response {
	jsonRes := make(map[string]interface{})
	if len(msg_optional) > 0 {
		jsonRes["message"] = msg_optional[0]
	} else {
		jsonRes["message"] = "OK"
	}
	jsonRes["success"] = true
	return NewResponse(http.StatusOK, jsonRes)
}

// only first provided arg will be used. (OPTIONAL)
func ClientErrResponse(msg_optional ...string) *Response {
	jsonRes := make(map[string]interface{})
	if len(msg_optional) > 0 {
		jsonRes["message"] = msg_optional[0]
	} else {
		jsonRes["message"] = "FAILURE. Bad API request."
	}
	jsonRes["success"] = false
	return NewResponse(http.StatusBadRequest, jsonRes)
}

// only first provided arg will be used. (OPTIONAL)
func ServerErrResponse(msg_optional ...string) *Response {
	jsonRes := make(map[string]interface{})
	if len(msg_optional) > 0 {
		jsonRes["message"] = msg_optional[0]
	} else {
		jsonRes["message"] = "FAILURE. Inernal server error."
	}

	jsonRes["success"] = false
	return NewResponse(http.StatusInternalServerError, jsonRes)
}
