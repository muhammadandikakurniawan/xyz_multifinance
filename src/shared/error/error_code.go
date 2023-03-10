package error

import "net/http"

type ErrorCode string

func (e ErrorCode) ToHttpStatus() int {
	switch e {
	case ERROR_UNAUTHORIZED:
		return http.StatusUnauthorized
	case ERROR_SURROUNDING:
		return http.StatusInternalServerError
	case ERROR_BAD_REQUEST:
		return http.StatusBadRequest
	case SUCCESS:
		return http.StatusOK
	}

	return http.StatusInternalServerError
}

const (
	ERROR_UNAUTHORIZED    ErrorCode = ErrorCode("CONS401")
	ERROR_SURROUNDING     ErrorCode = ErrorCode("CONS501")
	ERROR_BAD_REQUEST     ErrorCode = ErrorCode("CONS400")
	INTERNAL_SERVER_ERROR ErrorCode = ErrorCode("CONS500")
	SUCCESS               ErrorCode = ErrorCode("CONS200")
)
