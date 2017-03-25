package template

var BlockTemplate Template
var PizzaTemplate Template

func init() {
	// init block template
	BlockTemplate = newBlockTemplate()
	BlockTemplate.LoadFromJSONFile("./conf/template/block.json")

	// init pizza template
	PizzaTemplate = newPizzaTemplate()
	PizzaTemplate.LoadFromJSONFile("./conf/template/pizza.json")
}
