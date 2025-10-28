package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CanvasMusicVideoWork struct {
	UID            string         `gorm:"primaryKey;size:36" json:"uid"`
	CreatedDate    time.Time      `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate    time.Time      `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	AccountName    string         `json:"account_name"`
	AccountUID     string         `json:"account_uid"`
	VideoComic     string         `json:"video_comic"`
	Music          string         `json:"music"`
	Viewer         int            `json:"viewer"`
	Likes          int            `json:"likes"`
	AccountProfile AccountProfile `gorm:"foreignKey:AcccountUID;references:UID"`
	Account_Name   AccountProfile `gorm:"foreignKey:AcccountName;references:Name"`
}

func (cmvw *CanvasMusicVideoWork) BeforeCreate(tx *gorm.DB) (err error) {
	if cmvw.UID == "" {
		cmvw.UID = uuid.New().String()
	}
	return
}
