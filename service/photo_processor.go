package service

import (
	"bytes"
	"encoding/base64"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"os"
	"photo_builder/model"
	"photo_builder/model/template"
	"photo_builder/util"
	"time"

	"github.com/vence722/convert"
)

type photoProcessor struct {
	rootPath string
}

func newPhotoProcessor(rootPath string) *photoProcessor {
	return &photoProcessor{rootPath: rootPath}
}

func (this *photoProcessor) Process(photoBatch []*model.Photo, tmpl template.Template) (*model.Photo, error) {
	var photos []draw.Image
	for _, photoRaw := range photoBatch {
		data, err := base64.RawStdEncoding.DecodeString(photoRaw.DataBase64)
		if err != nil {
			return nil, err
		}
		img, err := util.DecodeAndHandleRotation(data)
		if err != nil {
			return nil, err
		}
		photo := util.ToDrawableImage(img)
		photos = append(photos, photo)
	}
	target, err := tmpl.ProcessPhoto(photos)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer([]byte{})
	jpeg.Encode(buf, target, &jpeg.Options{Quality: 100})

	targetPath := "./target/" + convert.Int2String(time.Now().Unix()) + ".jpg"
	ioutil.WriteFile(targetPath, buf.Bytes(), 0666)

	result := &model.Photo{}
	result.FileName = "target.jpg"
	result.DataBase64 = base64.RawStdEncoding.EncodeToString(buf.Bytes())

	// write to target folder
	var key = convert.Int2Str(time.Now().UnixNano())
	var targetFilePath = TargetPath + string(os.PathSeparator) + "target_" + key + ".jpg"
	_, err = os.Create(targetFilePath)
	if err == nil {
		ioutil.WriteFile(targetFilePath, buf.Bytes(), 0666)
	}

	return result, nil
}
