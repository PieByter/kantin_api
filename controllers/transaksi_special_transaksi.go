package controllers

import (
	"kantin_api/config"
	"kantin_api/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Create TransaksiSpecial
func CreateTransaksiSpecial(c *gin.Context) {
	var req models.TransaksiSpecial
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	now := time.Now()
	res, err := config.DB.Exec(
		`INSERT INTO transaksi_special (order_kupon_item_id, total_bayar, metode_pembayaran, tanggal_bayar, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		req.OrderKuponItemID, req.TotalBayar, req.MetodePembayaran, req.TanggalBayar, req.Status, now, now,
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

// Get all TransaksiSpecial
func GetTransaksiSpecials(c *gin.Context) {
	rows, err := config.DB.Query(
		`SELECT id, order_kupon_item_id, total_bayar, metode_pembayaran, tanggal_bayar, status, created_at, updated_at FROM transaksi_special`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()
	var transaksis []models.TransaksiSpecial
	for rows.Next() {
		var t models.TransaksiSpecial
		err := rows.Scan(
			&t.ID, &t.OrderKuponItemID, &t.TotalBayar, &t.MetodePembayaran, &t.TanggalBayar, &t.Status, &t.CreatedAt, &t.UpdatedAt,
		)
		if err == nil {
			transaksis = append(transaksis, t)
		}
	}
	c.JSON(http.StatusOK, transaksis)
}

// Get TransaksiSpecial by ID
func GetTransaksiSpecialByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var t models.TransaksiSpecial
	err := config.DB.QueryRow(
		`SELECT id, order_kupon_item_id, total_bayar, metode_pembayaran, tanggal_bayar, status, created_at, updated_at FROM transaksi_special WHERE id = ?`, id).
		Scan(&t.ID, &t.OrderKuponItemID, &t.TotalBayar, &t.MetodePembayaran, &t.TanggalBayar, &t.Status, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "TransaksiSpecial not found"})
		return
	}
	c.JSON(http.StatusOK, t)
}

// Update TransaksiSpecial
func UpdateTransaksiSpecial(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req models.TransaksiSpecial
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	now := time.Now()
	_, err := config.DB.Exec(
		`UPDATE transaksi_special SET order_kupon_item_id=?, total_bayar=?, metode_pembayaran=?, tanggal_bayar=?, status=?, updated_at=? WHERE id=?`,
		req.OrderKuponItemID, req.TotalBayar, req.MetodePembayaran, req.TanggalBayar, req.Status, now, id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "TransaksiSpecial updated"})
}

// Delete TransaksiSpecial
func DeleteTransaksiSpecial(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := config.DB.Exec("DELETE FROM transaksi_special WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "TransaksiSpecial deleted"})
}
