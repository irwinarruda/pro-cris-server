package authdrivers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/auth"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type AuthCtrl struct{}

func NewAuthCtrl() *AuthCtrl {
	return proinject.Resolve(&AuthCtrl{})
}

func (a *AuthCtrl) Login(c *gin.Context) {
	loginDTO := auth.LoginDTO{}
	err := c.Bind(&loginDTO)
	if err != nil {
		utils.HandleHttpError(c, utils.NewAppError("Invalid login."+err.Error(), false, http.StatusBadRequest))
		return
	}
	authService := auth.NewAuthService()
	account, err := authService.Login(loginDTO)
	if utils.HandleHttpError(c, err) {
		return
	}
	c.JSON(http.StatusOK, account)
}

func (a *AuthCtrl) EnsureAuthenticated(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		utils.HandleHttpError(c, utils.NewAppError("No token provided.", false, http.StatusBadRequest))
		c.Abort()
		return
	}
	authService := auth.NewAuthService()
	id, err := authService.EnsureAuthenticated(token, auth.LoginProviderGoogle)
	if utils.HandleHttpError(c, err) {
		c.Abort()
		return
	}
	c.Set("id_account", id)
	c.Next()
}
