package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type HttpAppError struct {
	Message         *string  `json:"message,omitempty"`
	Errors          []string `json:"messages,omitempty"`
	StatusCode      *int     `json:"statusCode,omitempty"`
	IsPublicMessage bool     `json:"isPublicMessage"`
}

type AppError struct {
	message         *string
	messages        []string
	statusCode      *int
	isPublicMessage bool
	baseMessage     string
}

func (e AppError) Error() string {
	return e.baseMessage
}

func NewAppError(message string, isPublicMessage bool, statusCode int) AppError {
	return AppError{
		isPublicMessage: isPublicMessage,
		message:         &message,
		statusCode:      &statusCode,
		baseMessage:     message,
	}
}

func HandleHttpError(c *gin.Context, err error) {
	if err, ok := err.(*AppError); ok {
		c.JSON(*err.statusCode, HttpAppError{
			Message:         err.message,
			Errors:          err.messages,
			StatusCode:      err.statusCode,
			IsPublicMessage: err.isPublicMessage,
		})
		return
	}

	if err, ok := err.(validator.ValidationErrors); ok {
		c.JSON(http.StatusBadRequest, HttpAppError{
			Errors: Map(err, func(item validator.FieldError, _ int) string {
				return item.Error()
			}),
			Message:         StringP("One or more fields are invalid."),
			StatusCode:      IntP(http.StatusBadRequest),
			IsPublicMessage: true,
		})
		return
	}

	c.JSON(http.StatusBadRequest, HttpAppError{
		Message:         StringP(err.Error()),
		Errors:          nil,
		StatusCode:      IntP(http.StatusInternalServerError),
		IsPublicMessage: false,
	})
}
