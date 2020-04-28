package router

import "github.com/gin-gonic/gin"

var GinApp *gin.Engine

func Init() {
	GinApp = gin.Default()

	GinApp.Use(gin.Logger())
	GinApp.Use(gin.Recovery())
}
