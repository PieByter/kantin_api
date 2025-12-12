package models

import (
	"database/sql"
)

type OrderRegulerItem struct {
	ID             int          `json:"id"`
	OrderRegulerID int          `json:"order_reguler_id"`
	MakananID      int          `json:"makanan_id"`
	Qty            int          `json:"qty"`
	HargaSatuan    int          `json:"harga_satuan"`
	Subtotal       int          `json:"subtotal"`
	CancelledAt    sql.NullTime `json:"cancelled_at"`
	CreatedAt      sql.NullTime `json:"created_at"`
	UpdatedAt      sql.NullTime `json:"updated_at"`
}
