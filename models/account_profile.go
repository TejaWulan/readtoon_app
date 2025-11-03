package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountProfile struct {
	UID         string    `gorm:"column:uniqueid;primaryKey;size:36" json:"uid"`
	CreatedDate time.Time `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	CreatedBy   string    `gorm:"column:created_by;default:null" json:"created_by"`
	UpdatedDate time.Time `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	UpdatedBy   string    `gorm:"column:updated_by;default:null" json:"updated_by"`
	Name        string    `gorm:"column:name" json:"name"`
	Email       string    `gorm:"column:email" json:"email"`
	Password    string    `gorm:"column:password" json:"password"`
	AvatarPhoto string    `gorm:"column:avatar_photo;default:null" json:"avatar_photo"`
}

func (a *AccountProfile) BeforeCreate(tx *gorm.DB) (err error) {
	if a.UID == "" {
		a.UID = uuid.New().String()
	}
	return
}
