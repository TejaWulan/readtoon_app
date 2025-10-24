package models

import (
	"time"
)

type ComicEpisodeSummary struct {
	EpisodeUID   string    `json:"episode_uid"`
	CreatedDate  time.Time `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	CreatedBy    string    `gorm:"column:created_by;default:null" json:"created_by"`
	UpdatedDate  time.Time `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	UpdatedBy    string    `gorm:"column:updated_by;default:null" json:"updated_by"`
	EpisodePhoto []string  `json:"episode_photo"` // multiple photos for episode summary
}
