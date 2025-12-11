package models

import "time"

type TransaksiSpecial struct {
	ID               int       `json:"id"`
	OrderKuponItemID int       `json:"order_kupon_item_id"`
	TotalBayar       int       `json:"total_bayar"`
	MetodePembayaran string    `json:"metode_pembayaran"`
	TanggalBayar     time.Time `json:"tanggal_bayar"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
