package models

import "time"

type TransaksiRegulerDetail struct {
	ID          int       `json:"id"`
	TransaksiID int       `json:"transaksi_id"`
	MakananID   int       `json:"makanan_id"`
	Jumlah      int       `json:"jumlah"`
	HargaSatuan int       `json:"harga_satuan"`
	DiskonKupon int       `json:"diskon_kupon"`
	Subtotal    int       `json:"subtotal"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
