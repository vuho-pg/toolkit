package api

import (
	"fmt"
	"github.com/vuho-pg/toolkit/pagination"
	"net/http"
)

type Metadata struct {
	*pagination.Pagination
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response interface {
	GetResponse() interface{}
	StatusCode() int
}

type DefaultResponse struct {
	Meta Metadata    `json:"meta"`
	Data interface{} `json:"data"`
}

func (r DefaultResponse) GetResponse() interface{} {
	return r
}

func (r DefaultResponse) StatusCode() int {
	return r.Meta.Code
}

func transformMessage(status int, msg string) string {
	if len(msg) != 0 {
		return msg
	}
	switch status {
	case http.StatusOK:
		return "success"
	case http.StatusBadRequest:
		return "bad request"
	case http.StatusInternalServerError:
		return "error"
	default:
		return fmt.Sprintf("status code %v", status)
	}
}

func CodeAndMessage(code int, message string) Response {
	return DefaultResponse{
		Meta: Metadata{
			Code:    code,
			Message: message,
		},
		Data: nil,
	}
}

func Success(data interface{}, message string) Response {
	return DefaultResponse{
		Meta: Metadata{Code: http.StatusOK, Message: transformMessage(http.StatusOK, message)},
		Data: data,
	}
}

func SuccessPagination(data interface{}, pagin pagination.Pagination, message string) Response {
	return DefaultResponse{
		Meta: Metadata{
			Pagination: &pagin,
			Code:       http.StatusOK,
			Message:    transformMessage(http.StatusOK, message),
		},
		Data: data,
	}
}

func BadRequest(message string) Response {
	return DefaultResponse{
		Meta: Metadata{
			Code:    http.StatusBadRequest,
			Message: transformMessage(http.StatusOK, message),
		},
	}
}

func InternalError(message string) Response {
	return DefaultResponse{
		Meta: Metadata{
			Code:    http.StatusInternalServerError,
			Message: message,
		},
	}
}
