package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Canvas struct {
	UID         string                 `gorm:"primaryKey;size:36" json:"uid"`
	AccountName string                 `json:"account_name"`
	AccountUID  string                 `json:"account_uid"`
	CreatedDate time.Time              `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate time.Time              `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	CanvasPhoto string                 `json:"canvas_photo"`
	Episodes    []CanvasEpisodeSummary `gorm:"-" json:"episodes"`
	MusicVideo  []string               `json:"music_video"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
}

func (c *Canvas) BeforeCreate(tx *gorm.DB) (err error) {
	if c.UID == "" {
		c.UID = uuid.New().String()
	}
	return
}
