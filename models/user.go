package models

import "time"

type User struct {
	ID                    int        `json:"id"`
	Nama                  string     `json:"nama"`
	NIK                   string     `json:"nik"`
	Email                 string     `json:"email"`
	Password              string     `json:"password"`
	NomorTelepon          *string    `json:"nomor_telepon"`
	Alamat                *string    `json:"alamat"`
	GambarProfil          *string    `json:"gambar_profil"`
	TanggalLahir          *time.Time `json:"tanggal_lahir"`
	TanggalBergabung      *time.Time `json:"tanggal_bergabung"`
	JabatanID             *int       `json:"jabatan_id"`
	BagianID              *int       `json:"bagian_id"`
	PenempatanID          *int       `json:"penempatan_id"`
	Keterangan            *string    `json:"keterangan"`
	Role                  string     `json:"role"`
	JumlahKuponTersedia   int        `json:"jumlah_kupon_tersedia"`
	JumlahKuponTerpakai   int        `json:"jumlah_kupon_terpakai"`
	JumlahKuponDibatalkan int        `json:"jumlah_kupon_dibatalkan"`
	CreatedAt             time.Time  `json:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at"`
}
