package models

import "time"

type MusicVideoSummary struct {
	MusicVideoUID string     `gorm:"column:music_video_uid" json:"music_video_uid"`
	CreatedDate   time.Time  `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate   time.Time  `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	Photo         string     `gorm:"column:photo" json:"photo"`
	Viewer        int        `gorm:"column:viewer" json:"viewer"`
	Likes         int        `gorm:"column:likes" json:"likes"`
	MusicVideo    MusicVideo `gorm:"foreignKey:MusicVideoUID;references:UID"`
}
