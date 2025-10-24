package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MusicVideoFavouriteLibrary struct {
	UID       string                    `gorm:"primaryKey;size:36" json:"uid"`
	Items     []MusicVideoFavouriteItem `gorm:"-" json:"items"`
	CreatedAt time.Time                 `json:"created_at"`
	UpdatedAt time.Time                 `json:"updated_at"`
}

func (mvfl *MusicVideoFavouriteLibrary) BeforeCreate(tx *gorm.DB) (err error) {
	if mvfl.UID == "" {
		mvfl.UID = uuid.New().String()
	}
	return
}
