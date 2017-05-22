package controllers

import (
	"encoding/json"
	"fmt"
	"photo_builder/model"
	"photo_builder/model/template"
	"photo_builder/service"

	"github.com/astaxie/beego"
)

type ProcessController struct {
	beego.Controller
}

func (this *ProcessController) Post() {
	tmplID := this.GetString("tmplID")
	selected := this.GetString("selected")
	var tmpl template.Template
	if tmplID == "1" {
		tmpl = template.BlockTemplate1
	} else {
		tmpl = template.BlockTemplate2
	}
	photoSels := parsePhotoSelects(selected)
	batch := service.PhotoStore.GetSelectedBatch(photoSels)
	photo1, err1 := service.PhotoProcessor.Process(batch, tmpl)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	photo2, err2 := service.PhotoProcessor.Process(batch, template.GalleryTemplate)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	this.Data["json"] = []*model.Photo{photo1, photo2}
	// move files to archive
	service.PhotoStore.MoveAllToArchive()
	this.ServeJSON()
}

func parsePhotoSelects(selected string) []*model.PhotoSelect {
	var photoSelects []*model.PhotoSelect
	err := json.Unmarshal([]byte(selected), &photoSelects)
	if err != nil {
		fmt.Println("failed to decode json at parsePhotoSelects(), input:", selected)
	}
	return photoSelects
}
