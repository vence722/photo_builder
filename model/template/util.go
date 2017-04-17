package template

import (
	"encoding/json"
	"image/draw"
	"io/ioutil"
	"os"
	"photo_builder/model/filter"
)

func saveToJSONFile(tmpl Template, filename string) error {
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		return err
	}
	enc := json.NewEncoder(f)
	enc.SetIndent("", "	")
	err = enc.Encode(tmpl)
	if err != nil {
		return err
	}
	return nil
}

func loadFromJSONFile(tmpl Template, filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, tmpl)
	if err != nil {
		return err
	}
	return nil
}

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
