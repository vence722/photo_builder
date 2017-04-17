package template

var BlockTemplate1 Template
var BlockTemplate2 Template
var PizzaTemplate Template
var GalleryTemplate Template

func init() {
	// init block templates
	BlockTemplate1 = newBlockTemplate()
	loadFromJSONFile(BlockTemplate1, "./conf/template/block1.json")

	BlockTemplate2 = newBlockTemplate()
	loadFromJSONFile(BlockTemplate2, "./conf/template/block2.json")

	// init pizza template
	PizzaTemplate = newPizzaTemplate()
	loadFromJSONFile(PizzaTemplate, "./conf/template/pizza.json")

	// init gallery template
	GalleryTemplate = newGalleryTemplate()
	loadFromJSONFile(GalleryTemplate, "./conf/template/gallery.json")
}
