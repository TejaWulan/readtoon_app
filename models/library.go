package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Library struct {
	UID         string      `gorm:"column:uniqueid;primaryKey" json:"uid"`
	CreatedDate time.Time   `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate time.Time   `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	Items       StringArray `gorm:"column:items" json:"items"`
	CreatedAt   time.Time   `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time   `gorm:"column:updated_at" json:"updated_at"`

	AccountUID     string         `gorm:"column:account_uid;size:36" json:"account_uid"` // foreign key field
	AccountProfile AccountProfile `gorm:"foreignKey:AccountUID;references:UID"`          // relation
}

func (l *Library) BeforeCreate(tx *gorm.DB) (err error) {
	if l.UID == "" {
		l.UID = uuid.New().String()
	}
	return
}
