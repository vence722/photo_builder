package filter

import "math"

type PizzaFilter struct {
	alpha   float64
	theta   float64
	offsetX int
	offsetY int
	r       int
}

func NewPizzaFilter(alpha float64, theta float64, offsetX int, offsetY int, r int) *PizzaFilter {
	return &PizzaFilter{
		alpha:   alpha,
		theta:   theta,
		offsetX: offsetX,
		offsetY: offsetY,
		r:       r,
	}
}

func (this *PizzaFilter) Filter(x int, y int) bool {
	xx := int(float64(x)*math.Cos(this.alpha) + float64(y)*math.Sin(this.alpha))
	yy := int(float64(y)*math.Cos(this.alpha) - float64(x)*math.Sin(this.alpha))
	offsetXX := int(float64(this.offsetX)*math.Cos(this.alpha) + float64(this.offsetY)*math.Sin(this.alpha))
	offsetYY := int(float64(this.offsetY)*math.Cos(this.alpha) - float64(this.offsetX)*math.Sin(this.alpha))
	if xx < offsetXX || yy < offsetYY || (xx-offsetXX)*(xx-offsetXX)+(yy-offsetYY)*(yy-offsetYY) >= this.r*this.r || float64(yy-offsetYY)/float64(xx-offsetXX) >= math.Tan(this.theta) {
		return false
	}
	return true
}
