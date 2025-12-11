package controllers

import (
	"kantin_api/config"
	"kantin_api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var req struct {
		Nama     string `json:"nama" binding:"required"`
		NIK      string `json:"nik" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	var exists int
	err := config.DB.QueryRow("SELECT COUNT(*) FROM users WHERE nik = ? OR email = ?", req.NIK, req.Email).Scan(&exists)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	if exists > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NIK or Email already exists"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	res, err := config.DB.Exec(
		`INSERT INTO users (nama, nik, email, password) VALUES (?, ?, ?, ?)`,
		req.Nama, req.NIK, req.Email, string(hashedPassword),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	id, _ := res.LastInsertId()
	c.JSON(http.StatusOK, gin.H{
		"message": "Register success",
		"user": gin.H{
			"id":    id,
			"nama":  req.Nama,
			"nik":   req.NIK,
			"email": req.Email,
		},
	})
}

func Login(c *gin.Context) {
	var req struct {
		NIKOrEmail string `json:"nik_or_email"`
		Password   string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	var user models.User
	err := config.DB.QueryRow(
		`SELECT id, nama, nik, email, password, nomor_telepon, alamat, gambar_profil, tanggal_lahir, tanggal_bergabung, jabatan_id, bagian_id, penempatan_id, keterangan, role, jumlah_kupon_tersedia, jumlah_kupon_terpakai, jumlah_kupon_dibatalkan, created_at, updated_at 
        FROM users WHERE nik = ? OR email = ?`,
		req.NIKOrEmail, req.NIKOrEmail,
	).Scan(
		&user.ID, &user.Nama, &user.NIK, &user.Email, &user.Password, &user.NomorTelepon, &user.Alamat, &user.GambarProfil, &user.TanggalLahir, &user.TanggalBergabung, &user.JabatanID, &user.BagianID, &user.PenempatanID, &user.Keterangan, &user.Role, &user.JumlahKuponTersedia, &user.JumlahKuponTerpakai, &user.JumlahKuponDibatalkan, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "NIK/Email or password incorrect"})
		return
	}
	// Verifikasi password hash
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "NIK/Email or password incorrect"})
		return
	}
	user.Password = "" // jangan kirim password ke frontend
	c.JSON(http.StatusOK, gin.H{"message": "Login success", "user": user})
}

func GetUsers(c *gin.Context) {
	rows, err := config.DB.Query(`
        SELECT id, nama, nik, email, password, nomor_telepon, alamat, gambar_profil, tanggal_lahir, tanggal_bergabung, jabatan_id, bagian_id, penempatan_id, keterangan, role, jumlah_kupon_tersedia, jumlah_kupon_terpakai, jumlah_kupon_dibatalkan, created_at, updated_at 
        FROM users`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(
			&u.ID, &u.Nama, &u.NIK, &u.Email, &u.Password, &u.NomorTelepon, &u.Alamat, &u.GambarProfil, &u.TanggalLahir, &u.TanggalBergabung, &u.JabatanID, &u.BagianID, &u.PenempatanID, &u.Keterangan, &u.Role, &u.JumlahKuponTersedia, &u.JumlahKuponTerpakai, &u.JumlahKuponDibatalkan, &u.CreatedAt, &u.UpdatedAt,
		)
		if err != nil {
			continue
		}
		users = append(users, u)
	}
	c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var u models.User
	err := config.DB.QueryRow(`
        SELECT id, nama, nik, email, password, nomor_telepon, alamat, gambar_profil, tanggal_lahir, tanggal_bergabung, jabatan_id, bagian_id, penempatan_id, keterangan, role, jumlah_kupon_tersedia, jumlah_kupon_terpakai, jumlah_kupon_dibatalkan, created_at, updated_at 
        FROM users WHERE id = ?`, id).
		Scan(
			&u.ID, &u.Nama, &u.NIK, &u.Email, &u.Password, &u.NomorTelepon, &u.Alamat, &u.GambarProfil, &u.TanggalLahir, &u.TanggalBergabung, &u.JabatanID, &u.BagianID, &u.PenempatanID, &u.Keterangan, &u.Role, &u.JumlahKuponTersedia, &u.JumlahKuponTerpakai, &u.JumlahKuponDibatalkan, &u.CreatedAt, &u.UpdatedAt,
		)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, u)
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	var OldPassword string
	err := config.DB.QueryRow("SELECT password FROM users WHERE id = ?", id).Scan(&OldPassword)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	_, err = config.DB.Exec(`
        UPDATE users SET nama=?, nik=?, email=?, password=?, nomor_telepon=?, alamat=?, gambar_profil=?, tanggal_lahir=?, tanggal_bergabung=?, jabatan_id=?, bagian_id=?, penempatan_id=?, keterangan=?, role=?, jumlah_kupon_tersedia=?, jumlah_kupon_terpakai=?, jumlah_kupon_dibatalkan=? WHERE id=?`,
		req.Nama, req.NIK, req.Email, OldPassword, req.NomorTelepon, req.Alamat, req.GambarProfil, req.TanggalLahir, req.TanggalBergabung, req.JabatanID, req.BagianID, req.PenempatanID, req.Keterangan, req.Role, req.JumlahKuponTersedia, req.JumlahKuponTerpakai, req.JumlahKuponDibatalkan, id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := config.DB.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func UpdatePassword(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req struct {
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid input"})
		return
	}
	var hashed string
	err := config.DB.QueryRow("SELECT password FROM users WHERE id = ?", id).Scan(&hashed)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(hashed), []byte(req.OldPassword)) != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Old Password incorrect"})
		return
	}
	newHashed, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	_, err = config.DB.Exec("UPDATE users SET password=? WHERE id=?", string(newHashed), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Password updated"})
}
