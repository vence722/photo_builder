package service

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"photo_builder/model"
)

type PhotoStore struct {
	rootPath string
}

func NewPhotoStore(rootPath string) *PhotoStore {
	return &PhotoStore{rootPath: rootPath}
}

func (this *PhotoStore) Cameras() []string {
	var cameras []string
	fileInfos, err := ioutil.ReadDir(this.rootPath)
	if err != nil {
		fmt.Println(err)
		return cameras
	}
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			cameras = append(cameras, fileInfo.Name())
		}
	}
	return cameras
}

func (this *PhotoStore) GetPhotos(cameraId string) []*model.Photo {
	return this.GetPhotosWithNum(cameraId, -1)
}

func (this *PhotoStore) GetPhotosWithNum(cameraId string, numOfPhotos int) []*model.Photo {
	var photos []*model.Photo
	fileInfos, err := ioutil.ReadDir(this.rootPath + string(os.PathSeparator) + cameraId)
	if err != nil {
		fmt.Println(err)
		return photos
	}
	count := 0
	for _, fileInfo := range fileInfos {
		if count == numOfPhotos {
			break
		}
		if !fileInfo.IsDir() && ".DS_Store" != fileInfo.Name() {
			photo := &model.Photo{
				FileName: fileInfo.Name(),
			}
			path, err := filepath.Abs(this.rootPath + string(os.PathSeparator) + cameraId + string(os.PathSeparator) + fileInfo.Name())
			if err != nil {
				fmt.Println(err)
				continue
			}
			photo.Path = path
			data, err := ioutil.ReadFile(photo.Path)
			if err != nil {
				fmt.Println(err)
				continue
			}
			photo.DataBase64 = base64.RawStdEncoding.EncodeToString(data)
			photos = append(photos, photo)
			count++
		}
	}
	return photos
}

func (this *PhotoStore) GetNextBatch() []*model.Photo {
	var photos []*model.Photo
	cameras := this.Cameras()
	for _, camera := range cameras {
		photoSigle := this.GetPhotosWithNum(camera, 1)
		photos = append(photos, photoSigle...)
	}
	return photos
}
