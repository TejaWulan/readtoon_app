package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MusicVideoFavouriteLibrary struct {
	UID            string         `gorm:"column:uniqueid;primaryKey" json:"uid"`
	Items          StringArray    `gorm:"column:items"  json:"items"`
	CreatedDate    time.Time      `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate    time.Time      `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	AccountName    string         `gorm:"column:account_name" json:"account_name"`
	AccountUID     string         `gorm:"column:account_uid" json:"account_uid"`
	AccountProfile AccountProfile `gorm:"foreignKey:AccountUID;references:UID"`
}

func (mvfl *MusicVideoFavouriteLibrary) BeforeCreate(tx *gorm.DB) (err error) {
	if mvfl.UID == "" {
		mvfl.UID = uuid.New().String()
	}
	return
}
