package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
)

type AuthCtrl struct {
	Validate configs.Validate `inject:"validate"`
}

func (a *AuthCtrl) Login(c *gin.Context) {
}
