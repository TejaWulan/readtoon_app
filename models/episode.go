package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Episode struct {
	UID         string      `gorm:"column:uniqueid;primaryKey" json:"uid"`
	Status      string      `gorm:"column:initial_status;default:ARCHIVE" json:"status"`
	Photos      StringArray `gorm:"column:photos" json:"photos"` // list of photo URLs/paths for the episode
	CreatedDate time.Time   `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate time.Time   `gorm:"column:updated_date;autoUpdateTime"  json:"-"`
}

func (e *Episode) BeforeCreate(tx *gorm.DB) (err error) {
	if e.UID == "" {
		e.UID = uuid.New().String()
	}
	return
}
