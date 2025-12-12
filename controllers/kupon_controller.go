package controllers

import (
	"database/sql"
	"kantin_api/config"
	"kantin_api/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Create Kupon
func CreateKupon(c *gin.Context) {
	var req models.Kupon
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	now := time.Now()
	res, err := config.DB.Exec(
		`INSERT INTO kupon (pembelian_kupon_id, user_id, jenis, kode, tipe_kupon, nilai_diskon, harga_pembelian, status, is_used, tanggal_berlaku, created_at, updated_at) 
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		req.PembelianKuponID, req.UserID, req.Jenis, req.Kode, req.TipeKupon, req.NilaiDiskon, req.HargaPembelian, req.Status, req.IsUsed, req.TanggalBerlaku, now, now,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	id, _ := res.LastInsertId()
	req.ID = int(id)
	req.CreatedAt = sql.NullTime{Time: now, Valid: true}
	req.UpdatedAt = sql.NullTime{Time: now, Valid: true}
	c.JSON(http.StatusOK, req)
}

// Get all Kupon
func GetKupons(c *gin.Context) {
	rows, err := config.DB.Query(
		`SELECT id, pembelian_kupon_id, user_id, jenis, kode, tipe_kupon, nilai_diskon, harga_pembelian, status, is_used, tanggal_berlaku,  created_at, updated_at FROM kupon`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()
	var kupons []models.Kupon
	for rows.Next() {
		var k models.Kupon
		err := rows.Scan(
			&k.ID, &k.PembelianKuponID, &k.UserID, &k.Jenis, &k.Kode, &k.TipeKupon, &k.NilaiDiskon, &k.HargaPembelian, &k.Status, &k.IsUsed, &k.TanggalBerlaku, &k.CreatedAt, &k.UpdatedAt,
		)
		if err == nil {
			kupons = append(kupons, k)
		}
	}
	c.JSON(http.StatusOK, kupons)
}

// Get Kupon by ID
func GetKuponByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var k models.Kupon
	err := config.DB.QueryRow(
		`SELECT id, pembelian_kupon_id, user_id, jenis, kode, tipe_kupon, nilai_diskon, harga_pembelian, status, is_used, tanggal_berlaku,created_at, updated_at FROM kupon WHERE id = ?`, id).
		Scan(&k.ID, &k.PembelianKuponID, &k.UserID, &k.Jenis, &k.Kode, &k.TipeKupon, &k.NilaiDiskon, &k.HargaPembelian, &k.Status, &k.IsUsed, &k.TanggalBerlaku, &k.CreatedAt, &k.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kupon not found"})
		return
	}
	c.JSON(http.StatusOK, k)
}

// Update Kupon
func UpdateKupon(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req models.Kupon
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	now := time.Now()
	_, err := config.DB.Exec(
		`UPDATE kupon SET pembelian_kupon_id=?, user_id=?, jenis=?, kode=?, tipe_kupon=?, nilai_diskon=?, harga_pembelian=?, status=?, is_used=?, tanggal_berlaku=?, updated_at=? WHERE id=?`,
		req.PembelianKuponID, req.UserID, req.Jenis, req.Kode, req.TipeKupon, req.NilaiDiskon, req.HargaPembelian, req.Status, req.IsUsed, req.TanggalBerlaku, now, id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Kupon updated"})
}

// Delete Kupon
func DeleteKupon(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := config.DB.Exec("DELETE FROM kupon WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Kupon deleted"})
}

func GetKuponsByStatus(c *gin.Context) {
	status := c.Param("status")
	rows, err := config.DB.Query(
		`SELECT id, pembelian_kupon_id, user_id, jenis, kode, tipe_kupon, nilai_diskon, harga_pembelian, status, is_used, tanggal_berlaku, created_at, updated_at FROM kupon WHERE status = ?`, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()
	var kupons []models.Kupon
	for rows.Next() {
		var k models.Kupon
		err := rows.Scan(
			&k.ID, &k.PembelianKuponID, &k.UserID, &k.Jenis, &k.Kode, &k.TipeKupon, &k.NilaiDiskon, &k.HargaPembelian, &k.Status, &k.IsUsed, &k.TanggalBerlaku, &k.CreatedAt, &k.UpdatedAt,
		)
		if err == nil {
			kupons = append(kupons, k)
		}
	}
	c.JSON(http.StatusOK, kupons)
}
