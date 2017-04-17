package service

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"photo_builder/model"
)

type photoStore struct {
	rootPath string
}

func newPhotoStore(rootPath string) *photoStore {
	return &photoStore{rootPath: rootPath}
}

func (this *photoStore) Cameras() []string {
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

func (this *photoStore) GetPhotos(cameraID string) []*model.Photo {
	return this.GetPhotosWithNum(cameraID, -1)
}

func (this *photoStore) GetPhotosWithNum(cameraID string, numOfPhotos int) []*model.Photo {
	var photos []*model.Photo
	fileInfos, err := ioutil.ReadDir(this.rootPath + string(os.PathSeparator) + cameraID)
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
			path, err := filepath.Abs(this.rootPath + string(os.PathSeparator) + cameraID + string(os.PathSeparator) + fileInfo.Name())
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

func (this *photoStore) GetNextBatch() []*model.Photo {
	var photos []*model.Photo
	cameras := this.Cameras()
	for _, camera := range cameras {
		photoSigle := this.GetPhotosWithNum(camera, 1)
		photos = append(photos, photoSigle...)
	}
	return photos
}

func (this *photoStore) GetSelectedBatch(photoSelects []*model.PhotoSelect) []*model.Photo {
	var photos []*model.Photo
	cameras := this.Cameras()
	for _, camera := range cameras {
		cameraPhotos := this.GetPhotos(camera)
		var photoSelect *model.PhotoSelect
		for _, sel := range photoSelects {
			if sel.Cid == camera {
				photoSelect = sel
				break
			}
		}
		if photoSelect == nil {
			fmt.Println("error: should select 1 photo of each camera")
			return nil
		}
		var resultPhoto *model.Photo
		for _, photo := range cameraPhotos {
			if photo.FileName == photoSelect.Filename {
				resultPhoto = photo
				break
			}
		}
		if resultPhoto == nil {
			fmt.Println("error: should select 1 photo of each camera")
			return nil
		}
		photos = append(photos, resultPhoto)
	}
	return photos
}
