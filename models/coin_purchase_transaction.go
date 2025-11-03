package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CoinPurchaseTransaction struct {
	UID                 string         `gorm:"column:uniqueid;primaryKey" json:"uid"`
	AccountUID          string         `gorm:"column:account_uid" json:"account_uid"` // who bought the coins
	AccountName         string         `gorm:"column:account_name" json:"account_name"`
	CreatedDate         time.Time      `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate         time.Time      `gorm:"column:updated_date;autoUpdateTime" json:"-"`
	CoinUID             string         `gorm:"column:coin_uid" json:"coin_uid"`                                     // related to Coin.UID (optional)
	AmountCoin          int64          `gorm:"column:amount_uid" json:"amount_coin"`                                // how many coins purchased
	PriceIDR            float64        `gorm:"column:price_idr" json:"price_idr"`                                   // price paid in rupiah
	PaymentMethod       string         `gorm:"column:payment_method" json:"payment_method"`                         // e.g. "GooglePlay", "TransferBank", "CreditCard"
	PaymentStatus       string         `gorm:"column:payment_status" json:"payment_status"`                         // e.g. "pending", "completed", "failed", "refunded"
	GoogleOrderID       string         `gorm:"column:product_uid" json:"google_order_id,omitempty"`                 // from Google Play purchase
	GooglePurchaseToken string         `gorm:"column:google_purchase_token" json:"google_purchase_token,omitempty"` // for verifying via Play API
	Verified            bool           `gorm:"column:verified" json:"verified"`                                     // whether Google Play purchase verified
	CreatedAt           time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt           time.Time      `gorm:"column:update_at" json:"updated_at"`
	AccountProfile      AccountProfile `gorm:"foreignKey:AccountUID;references:UID"`
	Coin                Coin           `gorm:"foreignKey:CoinUID;references:UID"`
}

func (cpt *CoinPurchaseTransaction) BeforeCreate(tx *gorm.DB) (err error) {
	if cpt.UID == "" {
		cpt.UID = uuid.New().String()
	}
	return
}
