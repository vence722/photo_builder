package main

import (
	"photo_builder/controllers"

	"github.com/astaxie/beego"
)

func configRouter() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/proess", &controllers.ProcessController{})
}

func main() {
	configRouter()

	beego.Run()
}
