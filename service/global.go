package service

import "github.com/astaxie/beego"

var RootPath = beego.AppConfig.String("path::inputpath")
var TargetPath = beego.AppConfig.String("path::outputpath")

var PhotoStore = newPhotoStore(RootPath)
var PhotoProcessor = newPhotoProcessor(RootPath)
