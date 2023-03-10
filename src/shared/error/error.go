package error

import "fmt"

func NewAppError(errCode ErrorCode, msg, errMsg string) error {
	return AppError{
		ErrorCode:          errCode,
		Message:            msg,
		SystemErrorMessage: errMsg,
	}
}

type AppError struct {
	ErrorCode          ErrorCode `json:"error_code"`
	Message            string    `json:"message"`
	SystemErrorMessage string    `json:"system_error_message"`
}

func (er AppError) Error() string {
	return fmt.Sprintf("[%s] | %s | %s", er.ErrorCode, er.Message, er.SystemErrorMessage)
}

func (er AppError) GetErrorCode() ErrorCode {
	return er.ErrorCode
}

func (er AppError) GetMessage() string {
	return er.Message
}

func (er AppError) GetSystemErrorMessage() string {
	return er.SystemErrorMessage
}
