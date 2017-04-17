package template

import (
	"errors"
	"image/draw"
	"photo_builder/model/filter"
	"photo_builder/util"
)

type blockTemplate struct {
	Base   base
	Blocks []block
}

type base struct {
	ResizeW int
	ResizeH int
}

type block struct {
	ResizeW int
	ResizeH int
	PutX    int
	PutY    int
}

func newBlockTemplate() *blockTemplate {
	return &blockTemplate{Blocks: []block{}}
}

func (this *blockTemplate) ProcessPhoto(photos []draw.Image) (draw.Image, error) {
	if len(photos) < 1 {
		return nil, errors.New("at least 1 photos needed")
	}
	base := util.Resize(photos[0], this.Base.ResizeW, this.Base.ResizeH)
	photos = photos[1:]
	for i, photo := range photos {
		if i < len(this.Blocks) {
			photos[i] = util.Resize(photo, this.Blocks[i].ResizeW, this.Blocks[i].ResizeH)
		}
	}
	for i, photo := range photos {
		if i < len(this.Blocks) {
			base = putPhoto(base, photo, this.Blocks[i].PutX, this.Blocks[i].PutY, filter.NewNonFilter())
		}
	}
	return base, nil
}
