// handler/Stunting_handler.go
package handler

import (
	"fmt"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CreateStunting(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bind JSON data from the request to the input variable
		var input model.Stunting
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

	

		// Save the Stunting to the database
		if err := db.Create(&input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Stunting"})
			return
		}

		// Respond with HTTP 201 Created status and the created Stunting data
		c.JSON(http.StatusCreated, model.ApiResponse{
			Message: "Stunting berhasil dibuat",
			Data:    input,
		})
	}
}

func GetStunting(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var result []model.Stunting

		// Mengambil data dari tabel "Stunting" dengan join ke tabel "departemen"
		if err := db.Table("Stuntings").
			Find(&result).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Memeriksa apakah slice "result" kosong
		if len(result) == 0 {
			c.JSON(404, gin.H{"message": "Tidak ada data Stunting"})
			return
		}

		c.JSON(200, result)
	}
}

func GetStuntingByUUID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("uuid")
		var Stunting struct {
			model.Stunting
		}
		if err := db.Table("Stuntings").
			Where("uuid = ?", uuid).
			First(&Stunting).
			Error; err != nil {
			c.JSON(404, gin.H{"error": "Stunting not found"})
			return
		}
		c.JSON(200, Stunting)
	}
}

func UpdateStunting(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mendapatkan UUID dari URL parameter
		uuid := c.Param("uuid")

		// Mencari Stunting berdasarkan UUID
		var existingStunting model.Stunting
		if err := db.Where("uuid = ?", uuid).First(&existingStunting).Preload("Doctor").Preload("Poly").Preload("Patient").Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Stunting not found"})
			return
		}

		// Binding input JSON ke struct Stunting
		var updatedStunting model.Stunting
		if err := c.ShouldBindJSON(&updatedStunting); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Memperbarui atribut Stunting
		db.Model(&existingStunting).Updates(&updatedStunting)

		// Mengambil data Stunting yang diperbarui dari database
		var updatedStuntingFromDB model.Stunting
		if err := db.Where("uuid = ?", uuid).First(&updatedStuntingFromDB).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated Stunting"})
			return
		}

		// Menyertakan pesan sukses dalam respons JSON
		c.JSON(http.StatusOK, model.ApiResponse{
			Message: "Stunting successfully updated",
			Data:    updatedStuntingFromDB,
		})

	}
}

func DeleteStunting(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("uuid")
		var Stunting model.Stunting
		if err := db.Where("uuid = ?", uuid).First(&Stunting).Error; err != nil {
			c.JSON(404, gin.H{"error": "Stunting not found"})
			return
		}

		StuntingName := Stunting.UUID

		db.Delete(&Stunting)
		c.JSON(200, gin.H{"message": fmt.Sprintf("%s deleted", StuntingName)})
	}
}

func PermanentDeleteStunting(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("uuid")
		var Stunting model.Stunting
		if err := db.Unscoped().Where("UUID = ?", uuid).First(&Stunting).Error; err != nil {
			c.JSON(404, gin.H{"error": "Stunting not found"})
			return
		}

		StuntingName := Stunting.UUID

		db.Unscoped().Delete(&Stunting)
		c.JSON(200, gin.H{"message": fmt.Sprintf("%s permanently deleted", StuntingName)})
	}
}
