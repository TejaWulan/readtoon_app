package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MusicVideoWorkLibrary struct {
	UID         string              `gorm:"primaryKey;size:36" json:"uid"`
	AccountName string              `json:"account_name"`
	AccountUID  string              `json:"account_uid"`
	Items       []MusicVideoSummary `gorm:"-" json:"items"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
}

func (mvwl *MusicVideoWorkLibrary) BeforeCreate(tx *gorm.DB) (err error) {
	if mvwl.UID == "" {
		mvwl.UID = uuid.New().String()
	}
	return
}
