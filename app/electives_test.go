package app

import (
	"RedRock-web-back-end-2020-6-lv2/config"
	"RedRock-web-back-end-2020-6-lv2/database"
	"RedRock-web-back-end-2020-6-lv2/response"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestGetAllElectives(t *testing.T) {
	database.Start()
	//app.GetAllElectives()
	//router.Start()
	r := gin.Default()
	r.POST("/class",H)
	r.Run()
}

func H(c *gin.Context)  {
	g := config.GetForm{}

	if err := c.ShouldBindJSON(&g); err != nil {
		errors.New("bind json error!")
		response.FormError(c)
		return
	}

	var e database.Electives
	database.G_db.Where("class_id = ? AND day = ? AND lesson = ? AND raw_week = ?", g.ClassId, g.Day, g.Lesson, g.Rawweek).Find(&e)
	fmt.Println(e)

}
