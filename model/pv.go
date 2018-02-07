package model

import "time"

type Pv struct {
	Id          int       `json:"id"`
	Date        time.Time `json:"date"`
	Url         string    `json:"url"`
	ReferrerUrl string    `json:"referrerUrl"`
	BrowserUuId string    `json:"browserUuId"`
	UserAgent   string    `json:"userAgent"`
	UserId      string    `json:"userId"`
	UserIp      string    `json:"userIp"`
}
