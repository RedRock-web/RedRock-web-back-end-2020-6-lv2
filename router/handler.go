package router

import (
	"RedRock-web-back-end-2020-6-lv2/app"
	"RedRock-web-back-end-2020-6-lv2/config"
	"RedRock-web-back-end-2020-6-lv2/response"
	"errors"
	"github.com/gin-gonic/gin"
)

func Handle(c *gin.Context) {
	g := config.GetForm{}
	if err := c.ShouldBindJSON(&g); err != nil {
		errors.New("bind json error!")
		response.FormError(c)
		return
	}
	app.GetAllClassInfo(g.StudentId)
	app.ChooseElective(g, c)
}
