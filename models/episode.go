package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Episode struct {
	UID           string      `gorm:"column:uniqueid;primaryKey" json:"uid"`
	Title         string      `gorm:"column:title" json:"title"`
	Status        string      `gorm:"column:initial_status;default:ARCHIVE" json:"status"`
	PhotoEpisodes StringArray `gorm:"column:photo_episodes" json:"photo_episodes"` // list of photo URLs/paths for the episode
	ComicUID      string      `gorm:"column:comic_uid" json:"comic_uid"`
	CanvasUID     string      `gorm:"column:canvas_uid" json:"canvas_uid"`
	Music         string      `gorm:"column:music" json:"music"`
	Viewer        int         `gorm:"column:viewer" json:"viewer"`
	Likes         int         `gorm:"column:likes" json:"likes"`
	CreatedDate   time.Time   `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate   time.Time   `gorm:"column:updated_date;autoUpdateTime"  json:"-"`
	Canvas        Canvas      `gorm:"foreignKey:CanvasUID;references:UID"`
	Comic         Comic       `gorm:"foreignKey:ComicUID;references:UID"`
}

func (e *Episode) BeforeCreate(tx *gorm.DB) (err error) {
	if e.UID == "" {
		e.UID = uuid.New().String()
	}
	return
}
