package models

import "database/sql"

type TransaksiReguler struct {
	ID             int          `json:"id"`
	OrderRegulerID int          `json:"order_reguler_id"`
	TotalHarga     int          `json:"total_harga"`
	TotalDiskon    int          `json:"total_diskon"`
	TotalBayar     int          `json:"total_bayar"`
	TanggalBayar   sql.NullTime `json:"tanggal_bayar"`
	Status         string       `json:"status"`
	Keterangan     *string      `json:"keterangan"`
	CreatedAt      sql.NullTime `json:"created_at"`
	UpdatedAt      sql.NullTime `json:"updated_at"`
}
