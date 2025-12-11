package models

import "time"

type Kupon struct {
	ID               int       `json:"id"`
	PembelianKuponID *int      `json:"pembelian_kupon_id"`
	UserID           int       `json:"user_id"`
	Jenis            string    `json:"jenis"`
	Kode             string    `json:"kode"`
	TipeKupon        string    `json:"tipe_kupon"`
	NilaiDiskon      int       `json:"nilai_diskon"`
	HargaPembelian   int       `json:"harga_pembelian"`
	Status           string    `json:"status"`
	IsUsed           bool      `json:"is_used"`
	TanggalBerlaku   time.Time `json:"tanggal_berlaku"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
