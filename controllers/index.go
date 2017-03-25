package controllers

import (
	"photo_builder/service"

	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	cameras := service.PhotoStore.Cameras()
	this.Data["Cameras"] = cameras
	this.TplName = "index.tpl"
}

func (this *IndexController) Post() {
	cid := this.GetString("cameraId")
	photos := service.PhotoStore.GetPhotos(cid)
	this.Data["json"] = photos
	this.ServeJSON()
}
