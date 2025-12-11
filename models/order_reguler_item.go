package models

import "time"

type OrderRegulerItem struct {
	ID             int        `json:"id"`
	OrderRegulerID int        `json:"order_reguler_id"`
	MakananID      int        `json:"makanan_id"`
	Qty            int        `json:"qty"`
	HargaSatuan    int        `json:"harga_satuan"`
	Subtotal       int        `json:"subtotal"`
	CancelledAt    *time.Time `json:"cancelled_at"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}
