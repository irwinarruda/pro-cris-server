package routes

import "github.com/gin-gonic/gin"

func CreateApiRoutes(group *gin.RouterGroup) {
	CreateStatusRoutes(group)
	CreateStudentRoutes(group)
	CreateAuthRoutes(group)
}
