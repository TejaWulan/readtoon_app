package models

type PaymentMethod struct {
	Code        string `gorm:"primaryKey;size:20" json:"code"` // e.g. "GooglePlay", "GoPay"
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}
