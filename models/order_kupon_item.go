package models

import "time"

type OrderKuponItem struct {
	ID              int        `json:"id"`
	OrderKuponID    int        `json:"order_kupon_id"`
	MakananID       int        `json:"makanan_id"`
	KuponID         int        `json:"kupon_id"`
	Qty             int        `json:"qty"`
	KuponTerpakai   int        `json:"kupon_terpakai"`
	TambahanBayar   int        `json:"tambahan_bayar"`
	Status          string     `json:"status"`
	TanggalKonsumsi time.Time  `json:"tanggal_konsumsi"`
	CancelledAt     *time.Time `json:"cancelled_at"`
	UsedAt          *time.Time `json:"used_at"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}
