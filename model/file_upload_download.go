package model

import (
	"gorm.io/gorm"
	"time"
)

type ExaFileUploadAndDownload struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Name      string         `json:"name" gorm:"comment:文件名"`
	Url       string         `json:"url" gorm:"comment:文件地址"`
	Tag       string         `json:"tag" gorm:"comment:文件标签"`
	Key       string         `json:"key" gorm:"comment:编号"`
}
