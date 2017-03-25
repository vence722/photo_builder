package service

const ROOT_PATH = "./photos"

var TEMPLATE_BLOCK Template = &BlockTemplate{}
var TEMPLATE_PIZZA Template = &PizzaTemplate{}

var PHOTO_STORE = NewPhotoStore(ROOT_PATH)
var PHOTO_PROCESSOR = NewPhotoProcessor(ROOT_PATH)
