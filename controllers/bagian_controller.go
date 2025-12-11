package controllers

import (
	"kantin_api/config"
	"kantin_api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Bagian
func CreateBagian(c *gin.Context) {
	var req models.Bagian
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	res, err := config.DB.Exec("INSERT INTO bagian (kode, nama) VALUES (?, ?)", req.Kode, req.Nama)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	id, _ := res.LastInsertId()
	req.ID = int(id)
	c.JSON(http.StatusOK, req)
}

// Get all Bagian
func GetBagians(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, kode, nama FROM bagian")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()
	var bagians []models.Bagian
	for rows.Next() {
		var b models.Bagian
		if err := rows.Scan(&b.ID, &b.Kode, &b.Nama); err == nil {
			bagians = append(bagians, b)
		}
	}
	c.JSON(http.StatusOK, bagians)
}

// Get Bagian by ID
func GetBagianByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var b models.Bagian
	err := config.DB.QueryRow("SELECT id, kode, nama FROM bagian WHERE id = ?", id).Scan(&b.ID, &b.Kode, &b.Nama)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bagian not found"})
		return
	}
	c.JSON(http.StatusOK, b)
}

// Update Bagian
func UpdateBagian(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req models.Bagian
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	_, err := config.DB.Exec("UPDATE bagian SET kode=?, nama=? WHERE id=?", req.Kode, req.Nama, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Bagian updated"})
}

// Delete Bagian
func DeleteBagian(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := config.DB.Exec("DELETE FROM bagian WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Bagian deleted"})
}
