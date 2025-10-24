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
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (mv *MusicVideo) BeforeCreate(tx *gorm.DB) (err error) {
	if mv.UID == "" {
		mv.UID = uuid.New().String()
	}
	return
}
