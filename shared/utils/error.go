package utils

type AppError struct {
	Message         string `json:"message"`
	IsPublicMessage bool   `json:"isPublicMessage"`
	RealError       error  `json:"-"`
}

func (e AppError) Error() string {
	return e.Message
}

func NewAppError(message string, isPublicMessage bool) AppError {
	return AppError{
		Message:         message,
		IsPublicMessage: isPublicMessage,
	}
}
