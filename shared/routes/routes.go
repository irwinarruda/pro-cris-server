package routes

import "github.com/gin-gonic/gin"

func CreateRoutes(group *gin.RouterGroup) {
	CreateStatusRoutes(group)
	CreateStudentRoutes(group)
}
