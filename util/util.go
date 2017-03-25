package util

import (
	"image"
	"image/draw"
	"image/jpeg"
	"os"
	"photo_builder/model"

	"github.com/nfnt/resize"
)

func ReadImage(path string) (draw.Image, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	img, err := jpeg.Decode(f)
	if err != nil {
		return nil, err
	}
	return ToDrawableImage(img), nil
}

func WriteImage(img draw.Image, path string) error {
	tf, err := os.Create(path)
	defer tf.Close()
	if err != nil {
		return err
	}
	return jpeg.Encode(tf, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
}

func PutPhoto(base draw.Image, src draw.Image, offsetX int, offsetY int, photoFilter model.PhotoFilter) draw.Image {
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

func Resize(img draw.Image, width uint, height uint) draw.Image {
	resized := resize.Resize(width, height, img, resize.Lanczos3)
	return ToDrawableImage(resized)
}

func ToDrawableImage(img image.Image) draw.Image {
	target := image.NewRGBA(img.Bounds())
	for i := 0; i < img.Bounds().Dx(); i++ {
		for j := 0; j < img.Bounds().Dy(); j++ {
			target.Set(i, j, img.At(i, j))
		}
	}
	return target
}
