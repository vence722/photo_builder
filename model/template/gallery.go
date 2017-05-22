package template

import (
	"errors"
	"image"
	"image/draw"
	"photo_builder/model/filter"
	"photo_builder/util"
)

type galleryTmplate struct {
	ConfigPath string
	NumX       int
	NumY       int
	ResizeX    int
	ResizeY    int
}

func newGalleryTemplate(configPath string) *galleryTmplate {
	return &galleryTmplate{ConfigPath: configPath}
}

func (this *galleryTmplate) ProcessPhoto(photos []draw.Image) (draw.Image, error) {
	if len(photos) < 1 {
		return nil, errors.New("at least 1 photos needed")
	}
	err := loadFromJSONFile(this, this.ConfigPath)
	if err != nil {
		return nil, errors.New("load config file err: " + err.Error())
	}
	rect := image.Rect(0, 0, this.NumX*this.ResizeX, this.NumY*this.ResizeY)
	base := image.NewRGBA(rect)
	for j := 0; j < this.NumY; j++ {
		for i := 0; i < this.NumX; i++ {
			photo := util.Resize(photos[j*this.NumY+i], this.ResizeX, this.ResizeY)
			fromX := i * this.ResizeX
			fromY := j * this.ResizeY
			putPhoto(base, photo, fromX, fromY, filter.NewNonFilter())
		}
	}
	return base, nil
}
