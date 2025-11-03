package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CanvasMusicVideoWork struct {
	UID            string         `gorm:"column:uniqueid;primaryKey" json:"uid"`
	CreatedDate    time.Time      `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate    time.Time      `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	AccountName    string         `gorm:"column:account_name" json:"account_name"`
	AccountUID     string         `gorm:"column:account_uid" json:"account_uid"`
	VideoComic     string         `gorm:"column:video_comic" json:"video_comic"`
	Music          string         `gorm:"column:music" json:"music"`
	Viewer         int            `gorm:"column:viewer" json:"viewer"`
	Likes          int            `gorm:"column:likes" json:"likes"`
	AccountProfile AccountProfile `gorm:"foreignKey:AccountUID;references:UID"`
}

func (cmvw *CanvasMusicVideoWork) BeforeCreate(tx *gorm.DB) (err error) {
	if cmvw.UID == "" {
		cmvw.UID = uuid.New().String()
	}
	return
}
