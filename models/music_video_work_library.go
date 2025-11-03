package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MusicVideoWorkLibrary struct {
	UID            string         `gorm:"column:uniqueid;primaryKey" json:"uid"`
	CreatedDate    time.Time      `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate    time.Time      `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	AccountName    string         `gorm:"column:account_name" json:"account_name"`
	AccountUID     string         `gorm:"column:account_uid" json:"account_uid"`
	Items          StringArray    `gorm:"column:items"  json:"items"`
	CreatedAt      time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"column:update_at" json:"updated_at"`
	AccountProfile AccountProfile `gorm:"foreignKey:AccountUID;references:UID"`
}

func (mvwl *MusicVideoWorkLibrary) BeforeCreate(tx *gorm.DB) (err error) {
	if mvwl.UID == "" {
		mvwl.UID = uuid.New().String()
	}
	return
}
