package models

import "time"

type MusicVideoSummary struct {
	MusicVideoUID string    `json:"music_video_uid"`
	CreatedDate   time.Time `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate   time.Time `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	Photo         string    `json:"photo"`
	Viewer        int       `json:"viewer"`
	Likes         int       `json:"likes"`
}
