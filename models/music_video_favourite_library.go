package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MusicVideoFavouriteLibrary struct {
	UID            string                    `gorm:"primaryKey;size:36" json:"uid"`
	Items          []MusicVideoFavouriteItem `gorm:"-" json:"items"`
	CreatedDate    time.Time                 `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate    time.Time                 `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	AccountName    string                    `json:"account_name"`
	AccountUID     string                    `json:"account_uid"`
	AccountProfile AccountProfile            `gorm:"foreignKey:AcccountUID;references:UID"`
	Account_Name   AccountProfile            `gorm:"foreignKey:AcccountName;references:Name"`
}

func (mvfl *MusicVideoFavouriteLibrary) BeforeCreate(tx *gorm.DB) (err error) {
	if mvfl.UID == "" {
		mvfl.UID = uuid.New().String()
	}
	return
}
