package template

import "image/draw"

type Template interface {
	ProcessPhoto(photos []draw.Image) (draw.Image, error)
}
