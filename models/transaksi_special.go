package models

import "database/sql"

type TransaksiSpecial struct {
	ID               int          `json:"id"`
	OrderKuponItemID int          `json:"order_kupon_item_id"`
	TotalBayar       int          `json:"total_bayar"`
	MetodePembayaran string       `json:"metode_pembayaran"`
	TanggalBayar     sql.NullTime `json:"tanggal_bayar"`
	Status           string       `json:"status"`
	CreatedAt        sql.NullTime `json:"created_at"`
	UpdatedAt        sql.NullTime `json:"updated_at"`
}
