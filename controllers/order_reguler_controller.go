package controllers

import (
	"kantin_api/config"
	"kantin_api/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Create OrderReguler
func CreateOrderReguler(c *gin.Context) {
	var req models.OrderReguler
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	now := time.Now()
	res, err := config.DB.Exec(
		`INSERT INTO order_reguler (user_id, tanggal_order, total_bayar, status, payment_method, keterangan, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		req.UserID, req.TanggalOrder, req.TotalBayar, req.Status, req.PaymentMethod, req.Keterangan, now, now,
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

// Get all OrderReguler
func GetOrderRegulers(c *gin.Context) {
	rows, err := config.DB.Query(
		`SELECT id, user_id, tanggal_order, total_bayar, status, payment_method, keterangan, created_at, updated_at FROM order_reguler`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()
	var orders []models.OrderReguler
	for rows.Next() {
		var o models.OrderReguler
		err := rows.Scan(
			&o.ID, &o.UserID, &o.TanggalOrder, &o.TotalBayar, &o.Status, &o.PaymentMethod, &o.Keterangan, &o.CreatedAt, &o.UpdatedAt,
		)
		if err == nil {
			orders = append(orders, o)
		}
	}
	c.JSON(http.StatusOK, orders)
}

// Get OrderReguler by ID
func GetOrderRegulerByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var o models.OrderReguler
	err := config.DB.QueryRow(
		`SELECT id, user_id, tanggal_order, total_bayar, status, payment_method, keterangan, created_at, updated_at FROM order_reguler WHERE id = ?`, id).
		Scan(&o.ID, &o.UserID, &o.TanggalOrder, &o.TotalBayar, &o.Status, &o.PaymentMethod, &o.Keterangan, &o.CreatedAt, &o.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "OrderReguler not found"})
		return
	}
	c.JSON(http.StatusOK, o)
}

// Update OrderReguler
func UpdateOrderReguler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req models.OrderReguler
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	now := time.Now()
	_, err := config.DB.Exec(
		`UPDATE order_reguler SET user_id=?, tanggal_order=?, total_bayar=?, status=?, payment_method=?, keterangan=?, updated_at=? WHERE id=?`,
		req.UserID, req.TanggalOrder, req.TotalBayar, req.Status, req.PaymentMethod, req.Keterangan, now, id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OrderReguler updated"})
}

// Delete OrderReguler
func DeleteOrderReguler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := config.DB.Exec("DELETE FROM order_reguler WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OrderReguler deleted"})
}

// CRUD untuk OrderRegulerItem

// Create OrderRegulerItem
func CreateOrderRegulerItem(c *gin.Context) {
	var req models.OrderRegulerItem
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	now := time.Now()
	res, err := config.DB.Exec(
		`INSERT INTO order_reguler_items (order_reguler_id, makanan_id, qty, harga_satuan, subtotal, cancelled_at, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		req.OrderRegulerID, req.MakananID, req.Qty, req.HargaSatuan, req.Subtotal, req.CancelledAt, now, now,
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

// Get all OrderRegulerItems
func GetOrderRegulerItems(c *gin.Context) {
	rows, err := config.DB.Query(
		`SELECT id, order_reguler_id, makanan_id, qty, harga_satuan, subtotal, cancelled_at, created_at, updated_at FROM order_reguler_items`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()
	var items []models.OrderRegulerItem
	for rows.Next() {
		var i models.OrderRegulerItem
		err := rows.Scan(
			&i.ID, &i.OrderRegulerID, &i.MakananID, &i.Qty, &i.HargaSatuan, &i.Subtotal, &i.CancelledAt, &i.CreatedAt, &i.UpdatedAt,
		)
		if err == nil {
			items = append(items, i)
		}
	}
	c.JSON(http.StatusOK, items)
}

// Get OrderRegulerItem by ID
func GetOrderRegulerItemByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var i models.OrderRegulerItem
	err := config.DB.QueryRow(
		`SELECT id, order_reguler_id, makanan_id, qty, harga_satuan, subtotal, cancelled_at, created_at, updated_at FROM order_reguler_items WHERE id = ?`, id).
		Scan(&i.ID, &i.OrderRegulerID, &i.MakananID, &i.Qty, &i.HargaSatuan, &i.Subtotal, &i.CancelledAt, &i.CreatedAt, &i.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "OrderRegulerItem not found"})
		return
	}
	c.JSON(http.StatusOK, i)
}

// Update OrderRegulerItem
func UpdateOrderRegulerItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req models.OrderRegulerItem
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	now := time.Now()
	_, err := config.DB.Exec(
		`UPDATE order_reguler_items SET order_reguler_id=?, makanan_id=?, qty=?, harga_satuan=?, subtotal=?, cancelled_at=?, updated_at=? WHERE id=?`,
		req.OrderRegulerID, req.MakananID, req.Qty, req.HargaSatuan, req.Subtotal, req.CancelledAt, now, id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OrderRegulerItem updated"})
}

// Delete OrderRegulerItem
func DeleteOrderRegulerItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := config.DB.Exec("DELETE FROM order_reguler_items WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OrderRegulerItem deleted"})
}
