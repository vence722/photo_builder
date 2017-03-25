package template

import (
	"image/draw"
	"photo_builder/model/filter"
)

func putPhoto(base draw.Image, src draw.Image, offsetX int, offsetY int, photoFilter filter.PhotoFilter) draw.Image {
	for i := 0; i < src.Bounds().Dx(); i++ {
		for j := 0; j < src.Bounds().Dy(); j++ {
			if photoFilter == nil || photoFilter.Filter(i, j) {
				x := offsetX + i
				y := offsetY + j

				base.Set(x, y, src.At(i, j))
			}
		}
	}
	return base
}
