package models

import "database/sql"

type TransaksiRegulerDetail struct {
	ID          int          `json:"id"`
	TransaksiID int          `json:"transaksi_reguler_id"`
	MakananID   int          `json:"makanan_id"`
	Jumlah      int          `json:"jumlah"`
	HargaSatuan int          `json:"harga_satuan"`
	DiskonKupon int          `json:"diskon_kupon"`
	Subtotal    int          `json:"subtotal"`
	CreatedAt   sql.NullTime `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
}
