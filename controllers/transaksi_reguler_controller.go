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

// Create TransaksiReguler
func CreateTransaksiReguler(c *gin.Context) {
	var req models.TransaksiReguler
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	now := time.Now()
	res, err := config.DB.Exec(
		`INSERT INTO transaksi_reguler (order_reguler_id, total_harga, total_diskon, total_bayar, tanggal_bayar, status, keterangan, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		req.OrderRegulerID, req.TotalHarga, req.TotalDiskon, req.TotalBayar, req.TanggalBayar, req.Status, req.Keterangan, now, now,
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

// Get all TransaksiReguler
func GetTransaksiRegulers(c *gin.Context) {
	rows, err := config.DB.Query(
		`SELECT id, order_reguler_id, total_harga, total_diskon, total_bayar, tanggal_bayar, status, keterangan, created_at, updated_at FROM transaksi_reguler`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()
	var transaksis []models.TransaksiReguler
	for rows.Next() {
		var t models.TransaksiReguler
		err := rows.Scan(
			&t.ID, &t.OrderRegulerID, &t.TotalHarga, &t.TotalDiskon, &t.TotalBayar, &t.TanggalBayar, &t.Status, &t.Keterangan, &t.CreatedAt, &t.UpdatedAt,
		)
		if err == nil {
			transaksis = append(transaksis, t)
		}
	}
	c.JSON(http.StatusOK, transaksis)
}

// Get TransaksiReguler by ID
func GetTransaksiRegulerByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var t models.TransaksiReguler
	err := config.DB.QueryRow(
		`SELECT id, order_reguler_id, total_harga, total_diskon, total_bayar, tanggal_bayar, status, keterangan, created_at, updated_at FROM transaksi_reguler WHERE id = ?`, id).
		Scan(&t.ID, &t.OrderRegulerID, &t.TotalHarga, &t.TotalDiskon, &t.TotalBayar, &t.TanggalBayar, &t.Status, &t.Keterangan, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "TransaksiReguler not found"})
		return
	}
	c.JSON(http.StatusOK, t)
}

// Update TransaksiReguler
func UpdateTransaksiReguler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req models.TransaksiReguler
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	now := time.Now()
	_, err := config.DB.Exec(
		`UPDATE transaksi_reguler SET order_reguler_id=?, total_harga=?, total_diskon=?, total_bayar=?, tanggal_bayar=?, status=?, keterangan=?, updated_at=? WHERE id=?`,
		req.OrderRegulerID, req.TotalHarga, req.TotalDiskon, req.TotalBayar, req.TanggalBayar, req.Status, req.Keterangan, now, id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "TransaksiReguler updated"})
}

// Delete TransaksiReguler
func DeleteTransaksiReguler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := config.DB.Exec("DELETE FROM transaksi_reguler WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "TransaksiReguler deleted"})
}

// CRUD untuk TransaksiRegulerDetail

// Create TransaksiRegulerDetail
func CreateTransaksiRegulerDetail(c *gin.Context) {
	var req models.TransaksiRegulerDetail
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	now := time.Now()
	res, err := config.DB.Exec(
		`INSERT INTO transaksi_reguler_detail (transaksi_reguler_id, makanan_id, jumlah, harga_satuan, diskon_kupon, subtotal, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		req.TransaksiID, req.MakananID, req.Jumlah, req.HargaSatuan, req.DiskonKupon, req.Subtotal, now, now,
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

// Get all TransaksiRegulerDetail
func GetTransaksiRegulerDetails(c *gin.Context) {
	rows, err := config.DB.Query(
		`SELECT id, transaksi_reguler_id, makanan_id, jumlah, harga_satuan, diskon_kupon, subtotal, created_at, updated_at FROM transaksi_reguler_detail`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()
	var details []models.TransaksiRegulerDetail
	for rows.Next() {
		var d models.TransaksiRegulerDetail
		err := rows.Scan(
			&d.ID, &d.TransaksiID, &d.MakananID, &d.Jumlah, &d.HargaSatuan, &d.DiskonKupon, &d.Subtotal, &d.CreatedAt, &d.UpdatedAt,
		)
		if err == nil {
			details = append(details, d)
		}
	}
	c.JSON(http.StatusOK, details)
}

// Get TransaksiRegulerDetail by ID
func GetTransaksiRegulerDetailByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var d models.TransaksiRegulerDetail
	err := config.DB.QueryRow(
		`SELECT id, transaksi_reguler_id, makanan_id, jumlah, harga_satuan, diskon_kupon, subtotal, created_at, updated_at FROM transaksi_reguler_detail WHERE id = ?`, id).
		Scan(&d.ID, &d.TransaksiID, &d.MakananID, &d.Jumlah, &d.HargaSatuan, &d.DiskonKupon, &d.Subtotal, &d.CreatedAt, &d.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "TransaksiRegulerDetail not found"})
		return
	}
	c.JSON(http.StatusOK, d)
}

// Update TransaksiRegulerDetail
func UpdateTransaksiRegulerDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req models.TransaksiRegulerDetail
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	now := time.Now()
	_, err := config.DB.Exec(
		`UPDATE transaksi_reguler_detail SET transaksi_reguler_id=?, makanan_id=?, jumlah=?, harga_satuan=?, diskon_kupon=?, subtotal=?, updated_at=? WHERE id=?`,
		req.TransaksiID, req.MakananID, req.Jumlah, req.HargaSatuan, req.DiskonKupon, req.Subtotal, now, id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "TransaksiRegulerDetail updated"})
}

// Delete TransaksiRegulerDetail
func DeleteTransaksiRegulerDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := config.DB.Exec("DELETE FROM transaksi_reguler_detail WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "TransaksiRegulerDetail deleted"})
}
