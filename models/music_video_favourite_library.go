package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MusicVideoFavouriteLibrary struct {
	UID         string                    `gorm:"primaryKey;size:36" json:"uid"`
	Items       []MusicVideoFavouriteItem `gorm:"-" json:"items"`
	CreatedDate time.Time                 `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate time.Time                 `gorm:"column:updated_date;autoUpdateTime" json:"-"`
}

func (mvfl *MusicVideoFavouriteLibrary) BeforeCreate(tx *gorm.DB) (err error) {
	if mvfl.UID == "" {
		mvfl.UID = uuid.New().String()
	}
	return
}
