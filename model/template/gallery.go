package template

import (
	"image"
	"image/draw"
	"photo_builder/model/filter"
	"photo_builder/util"
)

type galleryTmplate struct {
	NumX    int
	NumY    int
	ResizeX int
	ResizeY int
}

func newGalleryTemplate() *galleryTmplate {
	return &galleryTmplate{}
}

func (this *galleryTmplate) ProcessPhoto(photos []draw.Image) (draw.Image, error) {
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
