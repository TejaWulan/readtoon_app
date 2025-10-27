package models

import "time"

type PaymentMethod struct {
	Code        string    `gorm:"primaryKey;size:20" json:"code"` // e.g. "GooglePlay", "GoPay"
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsActive    bool      `json:"is_active"`
	CreatedDate time.Time `gorm:"column:created_date;autoCreateTime" json:"created_date"`
}
