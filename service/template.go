package service

import (
	"errors"
	"image/draw"
	"math"

	"photo_builder/model"
	"photo_builder/util"
)

type Template interface {
	ProcessPhoto(photos []draw.Image) (draw.Image, error)
}

type BlockTemplate struct{}

func (this *BlockTemplate) ProcessPhoto(photos []draw.Image) (draw.Image, error) {
	if len(photos) <= 0 {
		return nil, errors.New("at least 3 photos needed")
	}
	base := util.Resize(photos[0], 1000, 0)
	srcLeft := util.Resize(photos[1], 500, 0)
	srcRight := util.Resize(photos[2], 500, 0)

	base = util.PutPhoto(base, srcLeft, 150, 50, model.NewNonFilter())
	base = util.PutPhoto(base, srcRight, 450, 100, model.NewNonFilter())

	return base, nil
}

type PizzaTemplate struct{}

func (this *PizzaTemplate) ProcessPhoto(photos []draw.Image) (draw.Image, error) {
	if len(photos) <= 0 {
		return nil, errors.New("at least 3 photos needed")
	}
	base := util.Resize(photos[0], 1000, 0)
	srcLeft := util.Resize(photos[1], 500, 0)
	srcRight := util.Resize(photos[2], 500, 0)

	filterLeft := model.NewPizzaFilter(math.Pi*2/3, math.Pi/6, 350, 140, 350)
	filterRight := model.NewPizzaFilter(math.Pi/7, math.Pi/6, 50, 50, 350)
	base = util.PutPhoto(base, srcLeft, 150, 20, filterLeft)
	base = util.PutPhoto(base, srcRight, 450, 100, filterRight)

	return base, nil
}
