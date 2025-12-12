package models

import (
	"database/sql"
)

type OrderKuponItem struct {
	ID              int          `json:"id"`
	OrderKuponID    int          `json:"order_kupon_id"`
	MakananID       int          `json:"makanan_id"`
	KuponID         int          `json:"kupon_id"`
	Qty             int          `json:"qty"`
	KuponTerpakai   int          `json:"kupon_terpakai"`
	TambahanBayar   int          `json:"tambahan_bayar"`
	Status          string       `json:"status"`
	TanggalKonsumsi sql.NullTime `json:"tanggal_konsumsi"`
	CancelledAt     sql.NullTime `json:"cancelled_at"`
	UsedAt          sql.NullTime `json:"used_at"`
	CreatedAt       sql.NullTime `json:"created_at"`
	UpdatedAt       sql.NullTime `json:"updated_at"`
}
