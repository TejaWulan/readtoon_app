package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LibraryItem struct {
	UID          string    `gorm:"primaryKey;size:36" json:"uid"` // UUID string
	CreatedDate  time.Time `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	CreatedBy    string    `gorm:"column:created_by;default:null" json:"created_by"`
	UpdatedDate  time.Time `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	UpdatedBy    string    `gorm:"column:updated_by;default:null" json:"updated_by"`
	AccountUID   string    `json:"account_uid"`
	ComicUID     string    `json:"comic_uid"`
	ComicPhoto   string    `json:"comic_photo"`
	TotalEpisode int       `json:"total_episode"`
	Likes        int       `json:"likes"`
	Viewer       int       `json:"viewer"`
}

func (l *LibraryItem) BeforeCreate(tx *gorm.DB) (err error) {
	if l.UID == "" {
		l.UID = uuid.New().String()
	}
	return
}
