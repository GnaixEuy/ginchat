package service

import "github.com/gin-gonic/gin"

func GetIndex(c *gin.Context) {
	c.JSONP(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{
			"message": "Welcome ginChat",
			"version": "1.0.0",
		},
	})
}
