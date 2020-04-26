package config

type GetForm struct {
	StudentId int    `json:"student_id"`
	ClassId   string `json:"class_id"`
	Day       string `json:"day"`
	Lesson    string `json:"lesson"`
	Rawweek   string `json:"raw_week"`
}
