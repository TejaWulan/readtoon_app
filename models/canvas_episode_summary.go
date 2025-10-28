package models

import "time"

type CanvasEpisodeSummary struct {
	EpisodeUID  string    `json:"episode_uid"`
	CreatedDate time.Time `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate time.Time `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	Photos      []string  `json:"photos"`
	Episode     Episode   `gorm:"foreignKey:EpisodeUID;references:UID"`
}
