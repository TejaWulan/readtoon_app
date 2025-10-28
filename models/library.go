package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Library struct {
	UID            string         `gorm:"primaryKey;size:36" json:"uid"`
	CreatedDate    time.Time      `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate    time.Time      `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	Items          []LibraryItem  `gorm:"-" json:"items"` // gorm "-" means stored as JSON or separate table in actual DB design
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	AccountProfile AccountProfile `gorm:"foreignKey:AcccountUID;references:UID"`
	Account_Name   AccountProfile `gorm:"foreignKey:AcccountName;references:Name"`
}

func (l *Library) BeforeCreate(tx *gorm.DB) (err error) {
	if l.UID == "" {
		l.UID = uuid.New().String()
	}
	return
}
