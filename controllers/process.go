package controllers

import (
	"fmt"
	"photo_builder/service"

	"github.com/astaxie/beego"
)

type ProcessController struct {
	beego.Controller
}

func (this *ProcessController) Post() {
	tempId := this.GetString("tempId")
	var template service.Template
	if tempId == "1" {
		template = service.TEMPLATE_BLOCK
	} else {
		template = service.TEMPLATE_PIZZA
	}
	batch := service.PHOTO_STORE.GetNextBatch()
	photo, err := service.PHOTO_PROCESSOR.Process(batch, template)
	if err != nil {
		fmt.Println(err)
		return
	}
	this.Data["json"] = photo
	this.ServeJSON()
}
