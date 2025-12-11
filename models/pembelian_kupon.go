package models

import "time"

type PembelianKupon struct {
	ID               int       `json:"id"`
	UserID           int       `json:"user_id"`
	TanggalPembelian time.Time `json:"tanggal_pembelian"`
	Jumlah           int       `json:"jumlah"`
	TotalBayar       int       `json:"total_bayar"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
