package models

import (
	"time"
)

type ComicEpisodeSummary struct {
	EpisodeUID   string    `json:"episode_uid"`
	CreatedDate  time.Time `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate  time.Time `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	EpisodePhoto []string  `json:"episode_photo"` // multiple photos for episode summary
	Episode      Episode   `gorm:"foreignKey:EpisodeUID;references:UID"`
}
