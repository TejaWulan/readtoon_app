package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Canvas struct {
	UID            string         `gorm:"column:uniqueid;primaryKey" json:"uid"`
	AccountName    string         `gorm:"column:account_name" json:"account_name"`
	CreatorUID     string         `gorm:"column:creator_uid" json:"creator_uid"`
	CreatedDate    time.Time      `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate    time.Time      `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	Status         string         `gorm:"column:initial_status;default:ARCHIVE" json:"status"`
	Title          string         `gorm:"column:title" json:"title"`
	Description    string         `gorm:"column:description" json:"description"`
	Thumbnail      string         `gorm:"column:thumbnail_photo" json:"thumbnail_photo"`
	CanvasPhoto    string         `gorm:"column:canvas_photo" json:"canvas_photo"`
	EpisodeUIDs    StringArray    `gorm:"column:episode_uids"  json:"episode_uids"` // slice of episode uid + photo
	MusicVideoUIDs StringArray    `gorm:"column:music_video_uids" json:"music_video_uids"`
	AccountProfile AccountProfile `gorm:"foreignKey:CreatorUID;references:UID"`
}

func (c *Canvas) BeforeCreate(tx *gorm.DB) (err error) {
	if c.UID == "" {
		c.UID = uuid.New().String()
	}
	return
}
