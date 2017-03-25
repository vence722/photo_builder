package controllers

import (
	"fmt"
	"photo_builder/model/template"
	"photo_builder/service"

	"github.com/astaxie/beego"
)

type ProcessController struct {
	beego.Controller
}

func (this *ProcessController) Post() {
	tmplID := this.GetString("tmplID")
	var tmpl template.Template
	if tmplID == "1" {
		tmpl = template.BlockTemplate
	} else {
		tmpl = template.PizzaTemplate
	}
	batch := service.PhotoStore.GetNextBatch()
	photo, err := service.PhotoProcessor.Process(batch, tmpl)
	if err != nil {
		fmt.Println(err)
		return
	}
	this.Data["json"] = photo
	this.ServeJSON()
}
