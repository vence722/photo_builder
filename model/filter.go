package model

import (
	"math"
)

type PhotoFilter interface {
	Filter(int, int) bool
}

type NonFilter struct{}

func NewNonFilter() *NonFilter {
	return &NonFilter{}
}

func (this *NonFilter) Filter(x int, y int) bool {
	return true
}

type PizzaFilter struct {
	alpha    float64
	theta    float64
	offset_x int
	offset_y int
	r        int
}

func NewPizzaFilter(alpha float64, theta float64, offset_x int, offset_y int, r int) *PizzaFilter {
	return &PizzaFilter{
		alpha:    alpha,
		theta:    theta,
		offset_x: offset_x,
		offset_y: offset_y,
		r:        r,
	}
}

func (this *PizzaFilter) Filter(x int, y int) bool {
	xx := int(float64(x)*math.Cos(this.alpha) + float64(y)*math.Sin(this.alpha))
	yy := int(float64(y)*math.Cos(this.alpha) - float64(x)*math.Sin(this.alpha))
	offset_xx := int(float64(this.offset_x)*math.Cos(this.alpha) + float64(this.offset_y)*math.Sin(this.alpha))
	offset_yy := int(float64(this.offset_y)*math.Cos(this.alpha) - float64(this.offset_x)*math.Sin(this.alpha))
	if xx < offset_xx || yy < offset_yy || (xx-offset_xx)*(xx-offset_xx)+(yy-offset_yy)*(yy-offset_yy) >= this.r*this.r || float64(yy-offset_yy)/float64(xx-offset_xx) >= math.Tan(this.theta) {
		return false
	}
	return true
}
