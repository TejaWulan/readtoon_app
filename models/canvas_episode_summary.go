package models

import "time"

type CanvasEpisodeSummary struct {
	EpisodeUID  string      `gorm:"episode_uid" json:"episode_uid"`
	CreatedDate time.Time   `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate time.Time   `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	Photos      StringArray `gorm:"column:photos" json:"photos"`
	Episode     Episode     `gorm:"foreignKey:EpisodeUID;references:UID"`
}
