package controllers

import (
	"kantin_api/config"
	"kantin_api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Penempatan
func CreatePenempatan(c *gin.Context) {
	var req models.Penempatan
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	res, err := config.DB.Exec("INSERT INTO penempatan (kode, nama) VALUES (?, ?)", req.Kode, req.Nama)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	id, _ := res.LastInsertId()
	req.ID = int(id)
	c.JSON(http.StatusOK, req)
}

// Get all Penempatan
func GetPenempatans(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, kode, nama FROM penempatan")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()
	var penempatans []models.Penempatan
	for rows.Next() {
		var p models.Penempatan
		if err := rows.Scan(&p.ID, &p.Kode, &p.Nama); err == nil {
			penempatans = append(penempatans, p)
		}
	}
	c.JSON(http.StatusOK, penempatans)
}

// Get Penempatan by ID
func GetPenempatanByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var p models.Penempatan
	err := config.DB.QueryRow("SELECT id, kode, nama FROM penempatan WHERE id = ?", id).Scan(&p.ID, &p.Kode, &p.Nama)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Penempatan not found"})
		return
	}
	c.JSON(http.StatusOK, p)
}

// Update Penempatan
func UpdatePenempatan(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req models.Penempatan
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	_, err := config.DB.Exec("UPDATE penempatan SET kode=?, nama=? WHERE id=?", req.Kode, req.Nama, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Penempatan updated"})
}

// Delete Penempatan
func DeletePenempatan(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := config.DB.Exec("DELETE FROM penempatan WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Penempatan deleted"})
}
