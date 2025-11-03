package models

import "time"

type PaymentMethod struct {
	Code        string    `gorm:"primaryKey;size:20" json:"code"`                // e.g. "GooglePlay", "GoPay"
	AccountUID  string    `gorm:"column:account_uid;size:36" json:"account_uid"` // who bought the coins
	Name        string    `gorm:"column:name" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	IsActive    bool      `gorm:"column:is_active" json:"is_active"`
	CoinUID     string    `gorm:"column:coin_uid;size:36" json:"coin_uid"` // related to Coin.UID (optional)
	CreatedDate time.Time `gorm:"column:created_date;autoCreateTime" json:"created_date"`

	AccountProfile AccountProfile `gorm:"foreignKey:AccountUID;references:UID" json:"account_profile"`
	Coin           Coin           `gorm:"foreignKey:CoinUID;references:UID" json:"coin"`
}
