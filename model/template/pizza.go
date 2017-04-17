package template

import (
	"errors"
	"image/draw"
	"photo_builder/model/filter"
	"photo_builder/util"
)

type pizzaTemplate struct {
	Base        plate
	PizzaPieces []pizzaPiece
}

type plate struct {
	ResizeW int
	ResizeH int
}

type pizzaPiece struct {
	ResizeW int
	ResizeH int
	PutX    int
	PutY    int
	Alpha   float64
	Theta   float64
	OffsetX int
	OffsetY int
	R       int
}

func newPizzaTemplate() *pizzaTemplate {
	return &pizzaTemplate{PizzaPieces: []pizzaPiece{}}
}

func (this *pizzaTemplate) ProcessPhoto(photos []draw.Image) (draw.Image, error) {
	if len(photos) < 1 {
		return nil, errors.New("at least 1 photos needed")
	}

	base := util.Resize(photos[0], this.Base.ResizeW, this.Base.ResizeH)
	photos = photos[1:]
	for i, photo := range photos {
		if i < len(this.PizzaPieces) {
			photos[i] = util.Resize(photo, this.PizzaPieces[i].ResizeW, this.PizzaPieces[i].ResizeH)
		}
	}
	for i, photo := range photos {
		if i < len(this.PizzaPieces) {
			pizzaFilter := filter.NewPizzaFilter(this.PizzaPieces[i].Alpha, this.PizzaPieces[i].Theta, this.PizzaPieces[i].OffsetX, this.PizzaPieces[i].OffsetY, this.PizzaPieces[i].R)
			base = putPhoto(base, photo, this.PizzaPieces[i].PutX, this.PizzaPieces[i].PutY, pizzaFilter)
		}
	}
	return base, nil
}
