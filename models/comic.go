package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comic struct {
	UID           string                `gorm:"primaryKey;size:36" json:"uid"`
	CreatedDate   time.Time             `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	CreatedBy     string                `gorm:"column:created_by;default:null" json:"created_by"`
	UpdatedDate   time.Time             `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	UpdatedBy     string                `gorm:"column:updated_by;default:null" json:"updated_by"`
	AccountName   string                `json:"account_name"`      // nama_akun
	AccountUID    string                `json:"account_uid"`       // nama_akun_uid
	ComicPhoto    string                `json:"comic_photo"`       // cover / thumbnail
	CanvasWorkUID string                `json:"canvas_work_uid"`   // related canvas work uid (if any)
	Episodes      []ComicEpisodeSummary `gorm:"-" json:"episodes"` // slice of episode uid + photo
	MusicVideo    []string              `json:"music_video"`       // list of related music video UIDs
	CreatedAt     time.Time             `json:"created_at"`
	UpdatedAt     time.Time             `json:"updated_at"`
}

func (c *Comic) BeforeCreate(tx *gorm.DB) (err error) {
	if c.UID == "" {
		c.UID = uuid.New().String()
	}
	return
}
