package models

import "database/sql"

type OrderKupon struct {
	ID                 int          `json:"id"`
	UserID             int          `json:"user_id"`
	TanggalOrder       sql.NullTime `json:"tanggal_order"`
	TotalHari          int          `json:"total_hari"`
	TotalKuponTerpakai int          `json:"total_kupon_terpakai"`
	TotalUangTambahan  int          `json:"total_uang_tambahan"`
	Status             string       `json:"status"`
	Keterangan         *string      `json:"keterangan"`
	CreatedAt          sql.NullTime `json:"created_at"`
	UpdatedAt          sql.NullTime `json:"updated_at"`
}
