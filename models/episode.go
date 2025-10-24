package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Episode struct {
	UID       string    `gorm:"primaryKey;size:36" json:"uid"`
	Photos    []string  `gorm:"-" json:"photos"` // list of photo URLs/paths for the episode
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (e *Episode) BeforeCreate(tx *gorm.DB) (err error) {
	if e.UID == "" {
		e.UID = uuid.New().String()
	}
	return
}
