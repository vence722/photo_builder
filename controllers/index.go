package controllers

import (
	"photo_builder/service"

	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	cameras := service.PHOTO_STORE.Cameras()
	this.Data["Cameras"] = cameras
	this.TplName = "index.tpl"
}

func (this *IndexController) Post() {
	cid := this.GetString("cameraId")
	photos := service.PHOTO_STORE.GetPhotos(cid)
	this.Data["json"] = photos
	this.ServeJSON()
}
