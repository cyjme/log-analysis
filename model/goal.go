package model

type Goal struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	DataUrl string `json:"dataUrl"`
	TargetData int `json:"targetData"`
}
