package models

import "database/sql"

type Kupon struct {
	ID               int          `json:"id"`
	PembelianKuponID *int         `json:"pembelian_kupon_id"`
	UserID           int          `json:"user_id"`
	Jenis            string       `json:"jenis"`
	Kode             string       `json:"kode"`
	TipeKupon        string       `json:"tipe_kupon"`
	NilaiDiskon      int          `json:"nilai_diskon"`
	HargaPembelian   int          `json:"harga_pembelian"`
	Status           string       `json:"status"`
	IsUsed           bool         `json:"is_used"`
	TanggalBerlaku   sql.NullTime `json:"tanggal_berlaku"`
	CreatedAt        sql.NullTime `json:"created_at"`
	UpdatedAt        sql.NullTime `json:"updated_at"`
}
