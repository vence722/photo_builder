package util

import (
	"image"
	"image/draw"
	"image/jpeg"
	"os"

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

func Resize(img draw.Image, width int, height int) draw.Image {
	resized := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
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
