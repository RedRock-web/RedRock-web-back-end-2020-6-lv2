package database

type Class struct {
	gorm.Model
	Name      string
	StudentId int
	ClassId   string
	Location  string
	Day       string
	Lesson    string
	RawWeek   string
	Teacher   string
	Semester  string
}

type Electives struct {
	gorm.Model
	Name      string
	StudentId int
	ClassId   string
	Location  string
	Day       string
	Lesson    string
	RawWeek   string
	Teacher   string
	Semester  string
}


func Start()  {
	ConnetDb()
	CreateTable()
}

func ConnetDb() {
	db, err := gorm.Open("mysql", "root:mima@/students?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		errors.New("open database failed!")
	}
	G_db = db
}

func CreateTable() {
	if G_db.HasTable(&Class{}) {
		G_db.AutoMigrate()
	} else {
		G_db.CreateTable(&Class{})
	}
	if G_db.HasTable(&Electives{}) {
		G_db.AutoMigrate()
	} else {
		G_db.CreateTable(&Electives{})
	}
}
