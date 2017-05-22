package template

import (
	"errors"
	"image/draw"
	"photo_builder/model/filter"
	"photo_builder/util"
)

type blockTemplate struct {
	ConfigPath string
	Base       base
	Blocks     []block
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

func newBlockTemplate(configPath string) *blockTemplate {
	return &blockTemplate{ConfigPath: configPath, Blocks: []block{}}
}

func (this *blockTemplate) ProcessPhoto(photos []draw.Image) (draw.Image, error) {
	if len(photos) < 1 {
		return nil, errors.New("at least 1 photos needed")
	}
	err := loadFromJSONFile(this, this.ConfigPath)
	if err != nil {
		return nil, errors.New("load config file err: " + err.Error())
	}
	base := util.WhiteBackground(this.Base.ResizeW, this.Base.ResizeH)
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
