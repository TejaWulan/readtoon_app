package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountProfile struct {
	UID                      string    `gorm:"primaryKey;size:36" json:"uid"` // UUID string
	CreatedDate              time.Time `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	CreatedBy                string    `gorm:"column:created_by;default:null" json:"created_by"`
	UpdatedDate              time.Time `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	UpdatedBy                string    `gorm:"column:updated_by;default:null" json:"updated_by"`
	Name                     string    `json:"name"`
	Email                    string    `json:"email"`
	Password                 string    `json:"password"`                     // hashed password
	AvatarPhoto              string    `json:"avatar_photo"`                 // URL or path
	CoinUID                  string    `json:"coin_uid"`                     // FK to Coin.UID
	CanvasWorkLibraryUID     string    `json:"canvas_work_library_uid"`      // FK to CanvasWorkLibrary.UID
	LibraryUID               string    `json:"library_uid"`                  // FK to Library.UID
	MusicVideoFavLibraryUID  string    `json:"music_video_fav_library_uid"`  // FK to MusicVideoFavouriteLibrary.UID
	MusicVideoWorkLibraryUID string    `json:"music_video_work_library_uid"` // FK to MusicVideoWorkLibrary.UID

}

func (a *AccountProfile) BeforeCreate(tx *gorm.DB) (err error) {
	if a.UID == "" {
		a.UID = uuid.New().String()
	}
	return
}
