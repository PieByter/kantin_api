package models

import "time"

type OrderKupon struct {
	ID                 int       `json:"id"`
	UserID             int       `json:"user_id"`
	TanggalOrder       time.Time `json:"tanggal_order"`
	TotalHari          int       `json:"total_hari"`
	TotalKuponTerpakai int       `json:"total_kupon_terpakai"`
	TotalUangTambahan  int       `json:"total_uang_tambahan"`
	Status             string    `json:"status"`
	Keterangan         *string   `json:"keterangan"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
