package models

import "time"

type MusicVideoFavouriteItem struct {
	MusicVideoUID string     `gorm:"column:music_video_uid" json:"music_video_uid"`
	CreatedDate   time.Time  `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate   time.Time  `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	Photo         string     `gorm:"column:photo" json:"photo"` // thumbnail
	Viewer        int        `gorm:"column:viewer" json:"viewer"`
	Likes         int        `gorm:"column:likes" json:"likes"`
	AccountName   string     `gorm:"column:account_name" json:"account_name"` // nama_akun-owner
	MusicVideo    MusicVideo `gorm:"foreignKey:MusicVideoUID;references:UID"`
}
