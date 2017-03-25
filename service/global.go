package service

const RootPath = "./photos"

var PhotoStore = newPhotoStore(RootPath)
var PhotoProcessor = newPhotoProcessor(RootPath)
