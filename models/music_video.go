package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MusicVideo struct {
	UID         string    `gorm:"primaryKey;size:36" json:"uid"`
	AccountName string    `json:"account_name"`
	AccountUID  string    `json:"account_uid"`
	Photo       string    `json:"photo"`       // thumbnail
	VideoComic  string    `json:"video_comic"` // associated comic video file/path
	Music       string    `json:"music"`       // audio file/path
	Viewer      int       `json:"viewer"`
	Likes       int       `json:"likes"`
	CreatedDate time.Time `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate time.Time `gorm:"column:updated_date;autoUpdateTime" json:"-"`
}

func (mv *MusicVideo) BeforeCreate(tx *gorm.DB) (err error) {
	if mv.UID == "" {
		mv.UID = uuid.New().String()
	}
	return
}
