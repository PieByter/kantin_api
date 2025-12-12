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

// Create PembelianKupon
func CreatePembelianKupon(c *gin.Context) {
	var req models.PembelianKupon
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	now := time.Now()
	res, err := config.DB.Exec(
		`INSERT INTO pembelian_kupon (user_id, tanggal_pembelian, jumlah, total_bayar, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		req.UserID, req.TanggalPembelian, req.Jumlah, req.TotalBayar, req.Status, now, now,
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

// Get all PembelianKupon
func GetPembelianKupons(c *gin.Context) {
	rows, err := config.DB.Query(`SELECT id, user_id, tanggal_pembelian, jumlah, total_bayar, status, created_at, updated_at FROM pembelian_kupon`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()
	var kupons []models.PembelianKupon
	for rows.Next() {
		var p models.PembelianKupon
		err := rows.Scan(&p.ID, &p.UserID, &p.TanggalPembelian, &p.Jumlah, &p.TotalBayar, &p.Status, &p.CreatedAt, &p.UpdatedAt)
		if err == nil {
			kupons = append(kupons, p)
		}
	}
	c.JSON(http.StatusOK, kupons)
}

// Get PembelianKupon by ID
func GetPembelianKuponByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var p models.PembelianKupon
	err := config.DB.QueryRow(`SELECT id, user_id, tanggal_pembelian, jumlah, total_bayar, status, created_at, updated_at FROM pembelian_kupon WHERE id = ?`, id).
		Scan(&p.ID, &p.UserID, &p.TanggalPembelian, &p.Jumlah, &p.TotalBayar, &p.Status, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "PembelianKupon not found"})
		return
	}
	c.JSON(http.StatusOK, p)
}

// Update PembelianKupon
func UpdatePembelianKupon(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req models.PembelianKupon
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	now := time.Now()
	_, err := config.DB.Exec(
		`UPDATE pembelian_kupon SET user_id=?, tanggal_pembelian=?, jumlah=?, total_bayar=?, status=?, updated_at=? WHERE id=?`,
		req.UserID, req.TanggalPembelian, req.Jumlah, req.TotalBayar, req.Status, now, id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "PembelianKupon updated"})
}

// Delete PembelianKupon
func DeletePembelianKupon(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := config.DB.Exec("DELETE FROM pembelian_kupon WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "PembelianKupon deleted"})
}
