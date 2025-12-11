package controllers

import (
	"kantin_api/config"
	"kantin_api/models"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("your_secret_key")

func GenerateJWT(userID int, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // expired 24 jam
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

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
	jwtToken, err := GenerateJWT(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	user.Password = "" // jangan kirim password ke frontend
	c.JSON(http.StatusOK, gin.H{"message": "Login success", "user": user, "token": jwtToken})
}
