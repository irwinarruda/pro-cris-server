package utils

type AppError struct {
	Message         string `json:"message"`
	IsPublicMessage bool   `json:"isPublicMessage"`
	realError       error
}

func (e AppError) Error() string {
	return e.Message
}

func NewAppError(message string, isPublicMessage bool, err error) AppError {
	return AppError{
		Message:         message,
		IsPublicMessage: isPublicMessage,
		realError:       err,
	}
}
