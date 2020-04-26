package main

import (
	"RedRock-web-back-end-2020-6-lv2/app"
	"RedRock-web-back-end-2020-6-lv2/database"
	"RedRock-web-back-end-2020-6-lv2/router"
)

func main() {
	database.Start()
	app.GetAllElectives()
	router.Start()
}

