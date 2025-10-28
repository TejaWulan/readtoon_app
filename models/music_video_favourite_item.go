package models

import "time"

type MusicVideoFavouriteItem struct {
	MusicVideoUID string     `json:"music_video_uid"`
	CreatedDate   time.Time  `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate   time.Time  `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	Photo         string     `json:"photo"` // thumbnail
	Viewer        int        `json:"viewer"`
	Likes         int        `json:"likes"`
	AccountName   string     `json:"account_name"` // nama_akun-owner
	MusicVideo    MusicVideo `gorm:"foreignKey:MusicVideoUID;references:UID"`
	Owner         MusicVideo `gorm:"foreignKey:AccountName;references:AccountName"`
}
