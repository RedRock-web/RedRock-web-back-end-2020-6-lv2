package router

import "github.com/gin-gonic/gin"

func Start() {
	r := gin.Default()
	r.POST("/info", Handle)
	r.Run()
}
