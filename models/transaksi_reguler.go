package models

import "time"

type TransaksiReguler struct {
	ID             int       `json:"id"`
	OrderRegulerID int       `json:"order_reguler_id"`
	TotalHarga     int       `json:"total_harga"`
	TotalDiskon    int       `json:"total_diskon"`
	TotalBayar     int       `json:"total_bayar"`
	TanggalBayar   time.Time `json:"tanggal_bayar"`
	Status         string    `json:"status"`
	Keterangan     *string   `json:"keterangan"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
