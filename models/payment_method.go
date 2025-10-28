package models

import "time"

type PaymentMethod struct {
	Code         string         `gorm:"primaryKey;size:20" json:"code"` // e.g. "GooglePlay", "GoPay"
	AccountUID   string         `json:"account_uid"`                    // who bought the coins
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	IsActive     bool           `json:"is_active"`
	CoinUID      string         `json:"coin_uid"` // related to Coin.UID (optional)
	CreatedDate  time.Time      `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	Account_Name AccountProfile `gorm:"foreignKey:Name;references:Name"`
	Coin         Coin           `gorm:"foreignKey:CoinUID;references:UID"`
}
