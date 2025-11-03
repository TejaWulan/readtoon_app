package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Canvas struct {
	UID            string         `gorm:"column:uniqueid;primaryKey" json:"uid"`
	AccountName    string         `gorm:"column:account_name" json:"account_name"`
	AccountUID     string         `gorm:"column:account_uid" json:"account_uid"`
	CreatedDate    time.Time      `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate    time.Time      `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	Status         string         `gorm:"column:initial_status;default:ARCHIVE" json:"status"`
	CanvasPhoto    string         `gorm:"column:canvas_photo" json:"canvas_photo"`
	Episodes       StringArray    `gorm:"column:episodes"  json:"episodes"`
	MusicVideo     StringArray    `gorm:"column:music_video" json:"music_video"`
	AccountProfile AccountProfile `gorm:"foreignKey:AccountUID;references:UID"`
}

func (c *Canvas) BeforeCreate(tx *gorm.DB) (err error) {
	if c.UID == "" {
		c.UID = uuid.New().String()
	}
	return
}
