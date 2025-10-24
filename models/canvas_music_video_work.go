package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CanvasMusicVideoWork struct {
	UID         string    `gorm:"primaryKey;size:36" json:"uid"`
	AccountName string    `json:"account_name"`
	AccountUID  string    `json:"account_uid"`
	VideoComic  string    `json:"video_comic"`
	Music       string    `json:"music"`
	Viewer      int       `json:"viewer"`
	Likes       int       `json:"likes"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (cmvw *CanvasMusicVideoWork) BeforeCreate(tx *gorm.DB) (err error) {
	if cmvw.UID == "" {
		cmvw.UID = uuid.New().String()
	}
	return
}
