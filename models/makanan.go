package models

import "time"

type Makanan struct {
	ID            int       `json:"id"`
	Kode          string    `json:"kode"`
	Nama          string    `json:"nama"`
	Deskripsi     *string   `json:"deskripsi"`
	Harga         int       `json:"harga"`
	Stok          int       `json:"stok"`
	GambarMakanan *string   `json:"gambar_makanan"`
	Status        string    `json:"status"`
	IsAvailable   bool      `json:"is_available"`
	TipeMenu      string    `json:"tipe_menu"`
	IsCouponUse   bool      `json:"is_coupon_use"`
	TambahanBayar int       `json:"tambahan_bayar"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
