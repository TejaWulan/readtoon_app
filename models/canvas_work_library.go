package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CanvasWorkLibrary struct {
	UID            string         `gorm:"column:uniqueid;primaryKey" json:"uid"`
	AccountUID     string         `gorm:"column:account_uid" json:"account_uid"` // account this library belongs to
	AccountName    string         `gorm:"column:account_name" json:"account_name"`
	CreatedDate    time.Time      `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate    time.Time      `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	Items          StringArray    `gorm:"column:items"  json:"items"` // items for the same account
	CreatedAt      time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"column:update_at" json:"updated_at"`
	AccountProfile AccountProfile `gorm:"foreignKey:AccountUID;references:UID"`
}

func (cwl *CanvasWorkLibrary) BeforeCreate(tx *gorm.DB) (err error) {
	if cwl.UID == "" {
		cwl.UID = uuid.New().String()
	}
	return
}
