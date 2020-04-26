package app

import (
	"RedRock-web-back-end-2020-6-lv2/config"
	"RedRock-web-back-end-2020-6-lv2/database"
	"RedRock-web-back-end-2020-6-lv2/response"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"regexp"
)

const (
	CLASS = "'>(.*?)<\\/td>\\s*<td rowspan='1'>(.*?)-(.*?)<\\/td>\\s*<td rowspan='1'>(.*?)<\\/td>\\s*<td rowspan='1'>(.*?)<\\/td><td rowspan='1' align='center'>(.*?)<br> <\\/td><td>(.*?)<\\/td>\\s*<td>(星期\\d)(第.*?节) (.*?周)<\\/td><td>(.*?)<\\/td>\\s*<td rowspan='1' align='center'><a\\s*href='kb_stuList.php\\?jxb=(.*?)' target=_blank>名单<\\/a>"
)

func GetAllElectives() {
	var e database.Electives

	url := "http://jwc.cqupt.edu.cn/kebiao/kb_rw.php?NhGeWgGN=4_Qmrt.KklV.ZDgsbYCwypPJ0cVF3Cf0RgbAnLS19bY0reIJdbwB3xNqq3iE8oH0qGhTiy5Ilm1J45GIj_TV1DwJByQA8hk9YIsva6aGH2SbpnRPSWAE4bq4vPVQ6q00G8JRfJxXaNLJbLNW2L2GQyOez1Td7Z5nKioZAEtbr7PqvYX0HFCD04oUmEYRmq.oVsqRbo_B0YUkefnuu5TY3KdQ8l0qG64JwlxJZUVpArgQpdAKG_bTyyo6Ce6C4pxhmjYJVL048cnhW3PWYMUT3oNgaOgiosirXO9kLBFO.vZlKHbYdtKbU7vRha3bIK2BI7m.ZvkJxn6mBvSJrRhlZQ5Y4ODu2MYRFTdEWUYKGevsZcuo2ZT0yWuXbboRO_DuRTbWOVdfb3aW1MBXpKBwtIqt2nw4AU1Vmyt96OXtzMRyb1m4YpxWvM.FSddZvdrAXMqe"
	body := GetBody(url)
	re := regexp.MustCompile(CLASS)
	classs := re.FindAllStringSubmatch(body, -1)

	for _, v := range classs {
		e.ClassId = v[2]
		e.Name = v[3]
		e.Teacher = v[7]
		e.Day = v[8]
		e.RawWeek = v[10]
		e.Lesson = v[9]

		//Todo: fix bug
		database.G_db.Create(&e)
	}
}

func ChooseElective(g config.GetForm, c *gin.Context) (err error) {
	var class database.Class
	var e database.Electives

	tx := database.G_db.Begin()

	//Todo: fix bug
	if err = tx.Where("class_id = ? AND day = ? AND lesson = ? AND raw_week = ?", g.ClassId, g.Day, g.Lesson, g.Rawweek).Find(&e).Error; err != nil {
		response.Error(c, 10001, "该课不存在")
		fmt.Println(e)
		return err
	}

	tx.Where("class_id = ? AND day = ? AND lesson = ? AND raw_week = ?", g.ClassId, g.Day, g.Lesson, g.Rawweek).Find(&class)

	if IsConflict(class, e) {
		response.Error(c, 1002, "课程冲突！")
		return errors.New("configed!")
	}

	tx.Create(&database.Class{
		Name:      e.Name,
		StudentId: g.StudentId,
		ClassId:   e.ClassId,
		Day:       e.Day,
		Lesson:    e.Lesson,
		RawWeek:   e.RawWeek,
		Teacher:   e.Teacher,
	})

	response.OkWithData(c,"选课成功！")

	tx.Commit()

	return err
}

func IsConflict(c database.Class, e database.Electives) bool {
	return c.ClassId == e.ClassId && c.Day == e.Day && c.Lesson == e.Lesson && c.RawWeek == e.RawWeek
}
