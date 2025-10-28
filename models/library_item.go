package models

import (
	"time"
)

type LibraryItem struct {
	ComicUID     string    `json:"comic_uid"`
	CreatedDate  time.Time `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate  time.Time `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	AccountName  string    `json:"account_name"` // nama_akun-owner
	ComicPhoto   string    `json:"comic_photo"`
	TotalEpisode int       `json:"total_episode"`
	Likes        int       `json:"likes"`
	Viewer       int       `json:"viewer"`
	Owner        Canvas    `gorm:"foreignKey:AccountName;references:AccountName"`
}
