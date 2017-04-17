package model

type Photo struct {
	FileName   string
	Path       string
	DataBase64 string
}

type PhotoSelect struct {
	Cid      string `json:"cid"`
	Filename string `json:"filename"`
}
