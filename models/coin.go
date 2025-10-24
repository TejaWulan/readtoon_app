package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Coin struct {
	UID             string    `gorm:"primaryKey;size:36" json:"uid"`
	AccountName     string    `json:"account_name"`
	AccountUID      string    `json:"account_uid"`
	Amount          int64     `json:"amount"`            // number of coins
	ConversionToIDR float64   `json:"conversion_to_idr"` // conversion rate or converted value
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (co *Coin) BeforeCreate(tx *gorm.DB) (err error) {
	if co.UID == "" {
		co.UID = uuid.New().String()
	}
	return
}
