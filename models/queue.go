package models

import "gorm.io/gorm"

type Antri struct {
	gorm.Model
	Seq int `json:"ant"`
}