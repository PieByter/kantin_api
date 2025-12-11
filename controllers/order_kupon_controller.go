package controllers

import (
	"kantin_api/config"
	"kantin_api/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Create OrderKupon
func CreateOrderKupon(c *gin.Context) {
	var req models.OrderKupon
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	now := time.Now()
	res, err := config.DB.Exec(
		`INSERT INTO order_kupon (user_id, tanggal_order, total_hari, total_kupon_terpakai, total_uang_tambahan, status, keterangan, created_at, updated_at) 
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		req.UserID, req.TanggalOrder, req.TotalHari, req.TotalKuponTerpakai, req.TotalUangTambahan, req.Status, req.Keterangan, now, now,
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

// Get all OrderKupon
func GetOrderKupons(c *gin.Context) {
	rows, err := config.DB.Query(
		`SELECT id, user_id, tanggal_order, total_hari, total_kupon_terpakai, total_uang_tambahan, status, keterangan, created_at, updated_at FROM order_kupon`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()
	var kupons []models.OrderKupon
	for rows.Next() {
		var o models.OrderKupon
		err := rows.Scan(
			&o.ID, &o.UserID, &o.TanggalOrder, &o.TotalHari, &o.TotalKuponTerpakai, &o.TotalUangTambahan, &o.Status, &o.Keterangan, &o.CreatedAt, &o.UpdatedAt,
		)
		if err == nil {
			kupons = append(kupons, o)
		}
	}
	c.JSON(http.StatusOK, kupons)
}

// Get OrderKupon by ID
func GetOrderKuponByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var o models.OrderKupon
	err := config.DB.QueryRow(
		`SELECT id, user_id, tanggal_order, total_hari, total_kupon_terpakai, total_uang_tambahan, status, keterangan, created_at, updated_at FROM order_kupon WHERE id = ?`, id).
		Scan(&o.ID, &o.UserID, &o.TanggalOrder, &o.TotalHari, &o.TotalKuponTerpakai, &o.TotalUangTambahan, &o.Status, &o.Keterangan, &o.CreatedAt, &o.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "OrderKupon not found"})
		return
	}
	c.JSON(http.StatusOK, o)
}

// Update OrderKupon
func UpdateOrderKupon(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req models.OrderKupon
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	now := time.Now()
	_, err := config.DB.Exec(
		`UPDATE order_kupon SET user_id=?, tanggal_order=?, total_hari=?, total_kupon_terpakai=?, total_uang_tambahan=?, status=?, keterangan=?, updated_at=? WHERE id=?`,
		req.UserID, req.TanggalOrder, req.TotalHari, req.TotalKuponTerpakai, req.TotalUangTambahan, req.Status, req.Keterangan, now, id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OrderKupon updated"})
}

// Delete OrderKupon
func DeleteOrderKupon(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := config.DB.Exec("DELETE FROM order_kupon WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OrderKupon deleted"})
}

// CRUD untuk OrderKuponItem

// Create OrderKuponItem
func CreateOrderKuponItem(c *gin.Context) {
	var req models.OrderKuponItem
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	now := time.Now()
	res, err := config.DB.Exec(
		`INSERT INTO order_kupon_items (order_kupon_id, kupon_id, makanan_id, qty, kupon_terpakai, tambahan_bayar, status, tanggal_konsumsi, cancelled_at, used_at, created_at, updated_at) 
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		req.OrderKuponID, req.KuponID, req.MakananID, req.Qty, req.KuponTerpakai, req.TambahanBayar, req.Status, req.TanggalKonsumsi, req.CancelledAt, req.UsedAt, now, now,
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

// Get all OrderKuponItems
func GetOrderKuponItems(c *gin.Context) {
	rows, err := config.DB.Query(
		`SELECT id, order_kupon_id, kupon_id, makanan_id, qty, kupon_terpakai, tambahan_bayar, status, tanggal_konsumsi, cancelled_at, used_at, created_at, updated_at FROM order_kupon_items`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()
	var items []models.OrderKuponItem
	for rows.Next() {
		var i models.OrderKuponItem
		err := rows.Scan(
			&i.ID, &i.OrderKuponID, &i.KuponID, &i.MakananID, &i.Qty, &i.KuponTerpakai, &i.TambahanBayar, &i.Status, &i.TanggalKonsumsi, &i.CancelledAt, &i.UsedAt, &i.CreatedAt, &i.UpdatedAt,
		)
		if err == nil {
			items = append(items, i)
		}
	}
	c.JSON(http.StatusOK, items)
}

// Get OrderKuponItem by ID
func GetOrderKuponItemByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var i models.OrderKuponItem
	err := config.DB.QueryRow(
		`SELECT id, order_kupon_id, kupon_id, makanan_id, qty, kupon_terpakai, tambahan_bayar, status, tanggal_konsumsi, cancelled_at, used_at, created_at, updated_at FROM order_kupon_items WHERE id = ?`, id).
		Scan(&i.ID, &i.OrderKuponID, &i.KuponID, &i.MakananID, &i.Qty, &i.KuponTerpakai, &i.TambahanBayar, &i.Status, &i.TanggalKonsumsi, &i.CancelledAt, &i.UsedAt, &i.CreatedAt, &i.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "OrderKuponItem not found"})
		return
	}
	c.JSON(http.StatusOK, i)
}

// Update OrderKuponItem
func UpdateOrderKuponItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req models.OrderKuponItem
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	now := time.Now()
	_, err := config.DB.Exec(
		`UPDATE order_kupon_items SET order_kupon_id=?, kupon_id=?, makanan_id=?, qty=?, kupon_terpakai=?, tambahan_bayar=?, status=?, tanggal_konsumsi=?, cancelled_at=?, used_at=?, updated_at=? WHERE id=?`,
		req.OrderKuponID, req.KuponID, req.MakananID, req.Qty, req.KuponTerpakai, req.TambahanBayar, req.Status, req.TanggalKonsumsi, req.CancelledAt, req.UsedAt, now, id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OrderKuponItem updated"})
}

// Delete OrderKuponItem
func DeleteOrderKuponItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := config.DB.Exec("DELETE FROM order_kupon_items WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OrderKuponItem deleted"})
}
