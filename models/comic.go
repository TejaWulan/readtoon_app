package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comic struct {
	UID           string      `gorm:"column:uniqueid;primaryKey" json:"uid"`
	CreatedDate   time.Time   `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	CreatedBy     string      `gorm:"column:created_by;default:null" json:"created_by"`
	UpdatedDate   time.Time   `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	AccountName   string      `gorm:"column:account_name" json:"account_name"`       // nama_akun
	ComicPhoto    string      `gorm:"column:comic_photo" json:"comic_photo"`         // cover / thumbnail
	CanvasWorkUID string      `gorm:"column:canvas_work_uid" json:"canvas_work_uid"` // related canvas work uid (if any)
	Episodes      StringArray `gorm:"column:episodes"  json:"episodes"`              // slice of episode uid + photo
	MusicVideo    StringArray `gorm:"column:music_video" json:"music_video"`         // list of related music video UIDs
	Canvas        Canvas      `gorm:"foreignKey:CanvasWorkUID;references:UID"`
}

func (c *Comic) BeforeCreate(tx *gorm.DB) (err error) {
	if c.UID == "" {
		c.UID = uuid.New().String()
	}
	return
}
