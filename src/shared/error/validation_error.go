package error

func NewValidationError(msg string) AppError {
	return AppError{
		ErrorCode:          ERROR_BAD_REQUEST,
		Message:            msg,
		SystemErrorMessage: msg,
	}
}
