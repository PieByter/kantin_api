package models

import "database/sql"

type OrderReguler struct {
	ID            int          `json:"id"`
	UserID        int          `json:"user_id"`
	TanggalOrder  sql.NullTime `json:"tanggal_order"`
	TotalBayar    int          `json:"total_bayar"`
	Status        string       `json:"status"`
	PaymentMethod string       `json:"payment_method"`
	Keterangan    *string      `json:"keterangan"`
	CreatedAt     sql.NullTime `json:"created_at"`
	UpdatedAt     sql.NullTime `json:"updated_at"`
}
