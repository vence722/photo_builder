package template

var BlockTemplate1 Template
var BlockTemplate2 Template
var PizzaTemplate Template
var GalleryTemplate Template

func init() {
	// init block templates
	BlockTemplate1 = newBlockTemplate("./conf/template/block1.json")
	BlockTemplate2 = newBlockTemplate("./conf/template/block2.json")

	// init pizza template
	PizzaTemplate = newPizzaTemplate("./conf/template/pizza.json")

	// init gallery template
	GalleryTemplate = newGalleryTemplate("./conf/template/gallery.json")
}
