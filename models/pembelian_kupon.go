package models

import "database/sql"

type PembelianKupon struct {
	ID               int          `json:"id"`
	UserID           int          `json:"user_id"`
	TanggalPembelian sql.NullTime `json:"tanggal_pembelian"`
	Jumlah           int          `json:"jumlah"`
	TotalBayar       int          `json:"total_bayar"`
	Status           string       `json:"status"`
	CreatedAt        sql.NullTime `json:"created_at"`
	UpdatedAt        sql.NullTime `json:"updated_at"`
}
