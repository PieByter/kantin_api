package models

import "time"

type OrderReguler struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	TanggalOrder  time.Time `json:"tanggal_order"`
	TotalBayar    int       `json:"total_bayar"`
	Status        string    `json:"status"`
	PaymentMethod string    `json:"payment_method"`
	Keterangan    *string   `json:"keterangan"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
