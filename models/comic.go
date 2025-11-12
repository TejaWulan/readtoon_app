package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comic struct {
	UID            string         `gorm:"column:uniqueid;primaryKey" json:"uid"`
	CreatedDate    time.Time      `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	CreatedBy      string         `gorm:"column:created_by;default:null" json:"created_by"`
	UpdatedDate    time.Time      `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	Title          string         `gorm:"column:title" json:"title"`
	Description    string         `gorm:"column:description" json:"description"`
	Thumbnail      string         `gorm:"column:thumbnail_photo" json:"thumbnail_photo"`
	AuthorUID      string         `gorm:"column:author_uid" json:"author_uid"`
	CanvasWorkUID  string         `gorm:"column:canvas_work_uid" json:"canvas_work_uid"` // related canvas work uid (if any)
	EpisodeUIDs    StringArray    `gorm:"column:episode_uids"  json:"episode_uids"`      // slice of episode uid + photo
	MusicVideoUIDs StringArray    `gorm:"column:music_video_uids" json:"music_video_uids"`
	Canvas         Canvas         `gorm:"foreignKey:CanvasWorkUID;references:UID"`
	Author         AccountProfile `gorm:"foreignKey:AuthorUID;references:UID"`
}

func (c *Comic) BeforeCreate(tx *gorm.DB) (err error) {
	if c.UID == "" {
		c.UID = uuid.New().String()
	}
	return
}
