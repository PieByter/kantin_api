package controllers

import (
	"kantin_api/config"
	"kantin_api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Jabatan
func CreateJabatan(c *gin.Context) {
	var req models.Jabatan
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	res, err := config.DB.Exec("INSERT INTO jabatan (kode, nama) VALUES (?, ?)", req.Kode, req.Nama)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	id, _ := res.LastInsertId()
	req.ID = int(id)
	c.JSON(http.StatusOK, req)
}

// Get all Jabatan
func GetJabatans(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, kode, nama FROM jabatan")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()
	var jabatans []models.Jabatan
	for rows.Next() {
		var j models.Jabatan
		if err := rows.Scan(&j.ID, &j.Kode, &j.Nama); err == nil {
			jabatans = append(jabatans, j)
		}
	}
	c.JSON(http.StatusOK, jabatans)
}

// Get Jabatan by ID
func GetJabatanByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var j models.Jabatan
	err := config.DB.QueryRow("SELECT id, kode, nama FROM jabatan WHERE id = ?", id).Scan(&j.ID, &j.Kode, &j.Nama)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Jabatan not found"})
		return
	}
	c.JSON(http.StatusOK, j)
}

// Update Jabatan
func UpdateJabatan(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req models.Jabatan
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	_, err := config.DB.Exec("UPDATE jabatan SET kode=?, nama=? WHERE id=?", req.Kode, req.Nama, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Jabatan updated"})
}

// Delete Jabatan
func DeleteJabatan(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := config.DB.Exec("DELETE FROM jabatan WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Jabatan deleted"})
}
