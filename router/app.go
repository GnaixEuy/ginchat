package router

import (
	"ginchat/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	engine := gin.Default()
	engine.GET("/index", service.GetIndex)
	return engine
}
