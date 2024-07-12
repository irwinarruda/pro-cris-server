package authdrivers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/auth"
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
	loginDTO := auth.LoginDTO{}
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
	authService := auth.NewAuthService()
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
	authService := auth.NewAuthService()
	id, err := authService.EnsureAuthenticated(token, auth.Google)
	if err, ok := err.(utils.AppError); ok {
		c.JSON(http.StatusUnauthorized, err)
		c.Abort()
		return
	}
	c.Set("id_user", id)
	c.Next()
}
