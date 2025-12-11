package controllers

import (
	"kantin_api/config"
	"kantin_api/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Create Makanan
func CreateMakanan(c *gin.Context) {
	var req models.Makanan
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	now := time.Now()
	res, err := config.DB.Exec(
		`INSERT INTO makanan (kode, nama, deskripsi, harga, stok, gambar_makanan, is_available, tipe_menu, tambahan_bayar, created_at, updated_at) 
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		req.Kode, req.Nama, req.Deskripsi, req.Harga, req.Stok, req.GambarMakanan, req.IsAvailable, req.TipeMenu, req.TambahanBayar, now, now,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	id, _ := res.LastInsertId()
	req.ID = int(id)
	req.CreatedAt = now
	req.UpdatedAt = now
	c.JSON(http.StatusOK, req)
}

// Get all Makanan
func GetMakanans(c *gin.Context) {
	rows, err := config.DB.Query(
		`SELECT id, kode, nama, deskripsi, harga, stok, gambar_makanan, is_available, tipe_menu,  tambahan_bayar, created_at, updated_at FROM makanan`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()
	var makanans []models.Makanan
	for rows.Next() {
		var m models.Makanan
		err := rows.Scan(
			&m.ID, &m.Kode, &m.Nama, &m.Deskripsi, &m.Harga, &m.Stok, &m.GambarMakanan, &m.IsAvailable, &m.TipeMenu, &m.TambahanBayar, &m.CreatedAt, &m.UpdatedAt,
		)
		if err == nil {
			makanans = append(makanans, m)
		}
	}
	c.JSON(http.StatusOK, makanans)
}

// Get Makanan by ID
func GetMakananByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var m models.Makanan
	err := config.DB.QueryRow(
		`SELECT id, kode, nama, deskripsi, harga, stok, gambar_makanan, is_available, tipe_menu, tambahan_bayar, created_at, updated_at FROM makanan WHERE id = ?`, id).
		Scan(&m.ID, &m.Kode, &m.Nama, &m.Deskripsi, &m.Harga, &m.Stok, &m.GambarMakanan, &m.IsAvailable, &m.TipeMenu, &m.TambahanBayar, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Makanan not found"})
		return
	}
	c.JSON(http.StatusOK, m)
}

// Update Makanan
func UpdateMakanan(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req models.Makanan
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	now := time.Now()
	_, err := config.DB.Exec(
		`UPDATE makanan SET kode=?, nama=?, deskripsi=?, harga=?, stok=?, gambar_makanan=?, is_available=?, tipe_menu=?, tambahan_bayar=?, updated_at=? WHERE id=?`,
		req.Kode, req.Nama, req.Deskripsi, req.Harga, req.Stok, req.GambarMakanan, req.IsAvailable, req.TipeMenu, req.TambahanBayar, now, id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Makanan updated"})
}

// Delete Makanan
func DeleteMakanan(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := config.DB.Exec("DELETE FROM makanan WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Makanan deleted"})
}
