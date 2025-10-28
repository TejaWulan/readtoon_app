package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CoinPurchaseTransaction struct {
	UID                 string         `gorm:"primaryKey;size:36" json:"uid"`
	AccountUID          string         `json:"account_uid"` // who bought the coins
	AccountName         string         `json:"account_name"`
	CreatedDate         time.Time      `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate         time.Time      `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	CoinUID             string         `json:"coin_uid"`                        // related to Coin.UID (optional)
	AmountCoin          int64          `json:"amount_coin"`                     // how many coins purchased
	PriceIDR            float64        `json:"price_idr"`                       // price paid in rupiah
	PaymentMethod       string         `json:"payment_method"`                  // e.g. "GooglePlay", "TransferBank", "CreditCard"
	PaymentStatus       string         `json:"payment_status"`                  // e.g. "pending", "completed", "failed", "refunded"
	GoogleOrderID       string         `json:"google_order_id,omitempty"`       // from Google Play purchase
	GooglePurchaseToken string         `json:"google_purchase_token,omitempty"` // for verifying via Play API
	Verified            bool           `json:"verified"`                        // whether Google Play purchase verified
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	AccountProfile      AccountProfile `gorm:"foreignKey:AcccountUID;references:UID"`
	Account_Name        AccountProfile `gorm:"foreignKey:AcccountName;references:Name"`
	Coin                Coin           `gorm:"foreignKey:CoinUID;references:UID"`
}

func (cpt *CoinPurchaseTransaction) BeforeCreate(tx *gorm.DB) (err error) {
	if cpt.UID == "" {
		cpt.UID = uuid.New().String()
	}
	return
}
