package models

import (
	"time"
)

type LibraryItem struct {
	ComicUID     string    `gorm:"column:comic_uid" json:"comic_uid"`
	CreatedDate  time.Time `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate  time.Time `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	CanvasUID    string    `gorm:"canvas_uid" json:"episode_uid"`
	AccountName  string    `gorm:"column:account_name" json:"account_name"` // nama_akun-owner
	ComicPhoto   string    `gorm:"column:comic_photo" json:"comic_photo"`
	TotalEpisode int       `gorm:"column:total_episode" json:"total_episode"`
	Likes        int       `gorm:"column:likes" json:"likes"`
	Viewer       int       `gorm:"column:viewer" json:"viewer"`
	Owner        Canvas    `gorm:"foreignKey:CanvasUID;references:UID"`
}
