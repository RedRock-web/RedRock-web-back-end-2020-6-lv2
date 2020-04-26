package app

import (
	"RedRock-web-back-end-2020-6-lv2/database"
	"regexp"
	"strconv"
)

const (
	// 课表头信息
	HEADEINFO = "<li>〉〉(.*?学年)(\\d学期) 学生课表>>(\\d{10})(.*?)  <\\/li>"
	// 将课表筛选出来
	CLASSTABLE = "<div id=\"kbStuTabs-list\">(.*?|\\s*)*<\\/div>"
	// 将每门课筛选出来
	CLASSS = "<tr ?>(.|\\n|\\r)*?<\\/tr>"
	// 获取班级名和 id
	CLASSNAMEANDID = "<td rowspan='\\d'>(.*?)-(.*?)<\\/td>"
	// 获取其他班级信息
	OTHERCLASSINFO = "<td ?>(.*?)(<\\/td>|>)\\s*<td>(星期\\d) ?(第.*?节) ?(.*?周) ?<\\/td><td>(.*?)(<\\/td>|>)"
)

// 获取所有班级信息
func GetAllClassInfo(id int) {
	var c database.Class

	c.StudentId = id

	url := "http://jwc.cqupt.edu.cn/kebiao/kb_stu.php?xh=" + strconv.Itoa(id)
	body := GetBody(url)

	// 获取学期
	headInfoReg := regexp.MustCompile(HEADEINFO)
	headInfo := headInfoReg.FindAllStringSubmatch(body, -1)
	c.Semester = headInfo[0][2]

	// 筛选班级表
	tableReg := regexp.MustCompile(CLASSTABLE)
	table := tableReg.FindAllStringSubmatch(body, -1)

	// 筛选班级
	classsReg := regexp.MustCompile(CLASSS)
	classs := classsReg.FindAllStringSubmatch(table[0][0], -1)

	//获取班级信息
	for k, v := range classs {
		if k != 0 && k != len(classs)-1 {
			GetClassInfo(v[0], c)
		}
	}
}

// 获取详细的班级信息
func GetClassInfo(class string, c database.Class) {
	classsInfo1Reg := regexp.MustCompile(CLASSNAMEANDID)
	classInfo1 := classsInfo1Reg.FindAllStringSubmatch(class, -1)

	if len(classInfo1) == 0 {
		return
	}

	c.ClassId = classInfo1[0][1]
	c.Name = classInfo1[0][2]

	classsInfo2Reg := regexp.MustCompile(OTHERCLASSINFO)
	classInfo2 := classsInfo2Reg.FindAllStringSubmatch(class, -1)

	c.Teacher = classInfo2[0][1]
	c.Day = classInfo2[0][3]
	c.Lesson = classInfo2[0][4]
	c.RawWeek = classInfo2[0][5]
	c.Location = classInfo2[0][6]

	database.G_db.Create(&c)

}
