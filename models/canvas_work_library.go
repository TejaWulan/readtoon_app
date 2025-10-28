package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CanvasWorkLibrary struct {
	UID            string         `gorm:"primaryKey;size:36" json:"uid"`
	AccountUID     string         `json:"account_uid"` // account this library belongs to
	AccountName    string         `json:"account_name"`
	CreatedDate    time.Time      `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate    time.Time      `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	Items          []LibraryItem  `gorm:"-" json:"items"` // items for the same account
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	AccountProfile AccountProfile `gorm:"foreignKey:AcccountUID;references:UID"`
	Account_Name   AccountProfile `gorm:"foreignKey:AcccountName;references:Name"`
}

func (cwl *CanvasWorkLibrary) BeforeCreate(tx *gorm.DB) (err error) {
	if cwl.UID == "" {
		cwl.UID = uuid.New().String()
	}
	return
}
