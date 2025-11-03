package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MusicVideo struct {
	UID            string         `gorm:"column:uniqueid;primaryKey" json:"uid"`
	AccountName    string         `gorm:"column:account_name" json:"account_name"`
	AccountUID     string         `gorm:"column:account_uid" json:"account_uid"`
	Photo          string         `gorm:"column:photo" json:"photo"`             // thumbnail
	VideoComic     string         `gorm:"column:video_comic" json:"video_comic"` // associated comic video file/path
	Music          string         `gorm:"column:music" json:"music"`             // audio file/path
	Viewer         int            `gorm:"column:viewer" json:"viewer"`
	Likes          int            `gorm:"column:likes" json:"likes"`
	CreatedDate    time.Time      `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate    time.Time      `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	AccountProfile AccountProfile `gorm:"foreignKey:AccountUID;references:UID"`
}

func (mv *MusicVideo) BeforeCreate(tx *gorm.DB) (err error) {
	if mv.UID == "" {
		mv.UID = uuid.New().String()
	}
	return
}
