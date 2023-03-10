package moduleregistry

func NewInvalidParameter(msg string) error {
	return InvalidParameter{
		Message: msg,
	}
}

type InvalidParameter struct {
	Message string `json:"message"`
}

func (er InvalidParameter) Error() string {
	return er.Message
}
