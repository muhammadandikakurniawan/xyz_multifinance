package model

import (
	"net/http"

	sharedErr "github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/error"
)

type BaseResponseModel[T_DATA any] struct {
	Data           T_DATA `json:"data"`
	Message        string `json:"message"`
	ErrorMessage   string `json:"error_message"`
	Success        bool   `json:"success"`
	StatusCode     string `json:"status_code"`
	HttpStatusCode int    `json:"http_status_code"`
}

func (m *BaseResponseModel[T_DATA]) SetData(data T_DATA) *BaseResponseModel[T_DATA] {
	m.Data = data
	return m
}

func (m *BaseResponseModel[T_DATA]) SetMessage(data string) *BaseResponseModel[T_DATA] {
	m.Message = data
	return m
}

func (m *BaseResponseModel[T_DATA]) SetErrorMessage(data string) *BaseResponseModel[T_DATA] {
	m.ErrorMessage = data
	return m
}

func (m *BaseResponseModel[T_DATA]) SetSuccess(data bool) *BaseResponseModel[T_DATA] {
	m.Success = data
	return m
}

func (m *BaseResponseModel[T_DATA]) SetStatusCode(data string) *BaseResponseModel[T_DATA] {
	m.StatusCode = data
	return m
}

func (m *BaseResponseModel[T_DATA]) SetHttpStatusCode(data int) *BaseResponseModel[T_DATA] {
	m.HttpStatusCode = data
	return m
}

func (m *BaseResponseModel[T_DATA]) SetError(err error) *BaseResponseModel[T_DATA] {
	m.Success = false
	if appErr, ok := err.(sharedErr.AppError); ok {
		m.Message = appErr.Message
		m.ErrorMessage = appErr.SystemErrorMessage
		m.StatusCode = string(appErr.ErrorCode)
		m.HttpStatusCode = appErr.ErrorCode.ToHttpStatus()
		return m
	}
	m.Message = "internal server errorr"
	m.ErrorMessage = err.Error()
	m.HttpStatusCode = http.StatusInternalServerError
	m.StatusCode = string(sharedErr.INTERNAL_SERVER_ERROR)
	return m
}
