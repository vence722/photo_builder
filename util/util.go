package util

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"math"
	"os"

	"code.google.com/p/graphics-go/graphics"
	"github.com/nfnt/resize"
	"github.com/rwcarlsen/goexif/exif"
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

func DecodeAndHandleRotation(data []byte) (draw.Image, error) {
	img, err0 := jpeg.Decode(bytes.NewBuffer(data))
	if err0 != nil {
		return nil, err0
	}
	ex1, err1 := exif.Decode(bytes.NewBuffer(data))
	if err1 != nil {
		return nil, err1
	}
	orientation, err2 := ex1.Get(exif.Orientation)
	if err2 != nil {
		if err2.Error() == "exif: tag \"Orientation\" is not present" {
			// tag "Orientation" is not present
			return ToDrawableImage(img), nil
		}
		return nil, err2
	}
	// Should rotate 90 degrees
	if 6 == orientation.Val[0] {
		target := image.NewRGBA(image.Rect(0, 0, img.Bounds().Dy(), img.Bounds().Dx()))
		err3 := graphics.Rotate(target, img, &graphics.RotateOptions{Angle: math.Pi / 2})
		if err3 != nil {
			return nil, err3
		}
		return target, nil
	} else if 8 == orientation.Val[0] {
		target := image.NewRGBA(image.Rect(0, 0, img.Bounds().Dy(), img.Bounds().Dx()))
		err3 := graphics.Rotate(target, img, &graphics.RotateOptions{Angle: -math.Pi / 2})
		if err3 != nil {
			return nil, err3
		}
		return target, nil
	}
	return ToDrawableImage(img), nil
}

func Encode(img image.Image) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	if err := jpeg.Encode(buf, img, &jpeg.Options{Quality: 100}); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
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

func WhiteBackground(width, height int) draw.Image {
	background := image.NewRGBA(image.Rect(0, 0, width, height))
	for i := 0; i < background.Bounds().Dx(); i++ {
		for j := 0; j < background.Bounds().Dy(); j++ {
			background.Set(i, j, color.White)
		}
	}
	return ToDrawableImage(background)
}
