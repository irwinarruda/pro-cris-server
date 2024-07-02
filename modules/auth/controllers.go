package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type AuthCtrl struct {
	Validate configs.Validate `inject:"validate"`
}

func NewAuthCtrl() *AuthCtrl {
	return proinject.Resolve(&AuthCtrl{})
}

func (a *AuthCtrl) Login(c *gin.Context) {
	loginDTO := LoginDTO{}
	err := c.Bind(&loginDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewAppError("Invalid student data"+err.Error(), false, err))
		return
	}
	err = a.Validate.Struct(loginDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewAppError("Invalid student data"+err.Error(), false, nil))
		return
	}
	authService := NewAuthService()
	user, err := authService.Login(loginDTO)
	if err, ok := err.(utils.AppError); ok {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (a *AuthCtrl) EnsureAuthenticated(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, utils.NewAppError("No token provided.", false, nil))
		c.Abort()
		return
	}
	authService := NewAuthService()
	id, err := authService.EnsureAuthenticated(token, Google)
	if err, ok := err.(utils.AppError); ok {
		c.JSON(http.StatusUnauthorized, err)
		c.Abort()
		return
	}
	c.Set("id_user", id)
	c.Next()
}
