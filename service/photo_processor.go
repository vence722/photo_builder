package service

import (
	"bytes"
	"encoding/base64"
	"image/draw"
	"image/jpeg"
	"photo_builder/model"
	"photo_builder/util"
)

type PhotoProcessor struct {
	rootPath string
}

func NewPhotoProcessor(rootPath string) *PhotoProcessor {
	return &PhotoProcessor{rootPath: rootPath}
}

func (this *PhotoProcessor) Process(photoBatch []*model.Photo, template Template) (*model.Photo, error) {
	var photos []draw.Image
	for _, photoRaw := range photoBatch {
		data, err := base64.RawStdEncoding.DecodeString(photoRaw.DataBase64)
		if err != nil {
			return nil, err
		}
		img, err := jpeg.Decode(bytes.NewReader(data))
		if err != nil {
			return nil, err
		}
		photo := util.ToDrawableImage(img)
		photos = append(photos, photo)
	}
	target, err := template.ProcessPhoto(photos)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer([]byte{})
	jpeg.Encode(buf, target, &jpeg.Options{Quality: 100})

	result := &model.Photo{}
	result.FileName = "target.jpg"
	result.DataBase64 = base64.RawStdEncoding.EncodeToString(buf.Bytes())

	return result, nil
}
