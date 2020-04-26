package main

import (
	"RedRock-web-back-end-2020-6-lv2/app"
	"RedRock-web-back-end-2020-6-lv2/database"
)

func main() {
	database.Start()
	app.GetAllElectives()
	app.GetAllClassInfo(2019211548)
}

