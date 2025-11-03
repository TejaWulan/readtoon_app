package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Coin struct {
	UID             string         `gorm:"primaryKey;size:36" json:"uid"`
	AccountName     string         `gorm:"column:account_name" json:"account_name"`
	AccountUID      string         `gorm:"column:account_uid" json:"account_uid"`
	CreatedDate     time.Time      `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate     time.Time      `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	Amount          int64          `gorm:"column:amount" json:"amount"`                       // number of coins
	ConversionToIDR float64        `gorm:"column:conversion_to_idr" json:"conversion_to_idr"` // conversion rate or converted value
	CreatedAt       time.Time      `gorm:"column:product_uid" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"column:updated_at" json:"updated_at"`
	AccountProfile  AccountProfile `gorm:"foreignKey:AccountUID;references:UID"`
}

func (co *Coin) BeforeCreate(tx *gorm.DB) (err error) {
	if co.UID == "" {
		co.UID = uuid.New().String()
	}
	return
}
