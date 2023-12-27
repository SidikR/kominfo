// handler/Program_handler.go
package handler

import (
	"fmt"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CreateProgram(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bind JSON data from the request to the input variable
		var input model.Program
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

	

		// Save the Program to the database
		if err := db.Create(&input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Program"})
			return
		}

		// Respond with HTTP 201 Created status and the created Program data
		c.JSON(http.StatusCreated, model.ApiResponse{
			Message: "Program berhasil dibuat",
			Data:    input,
		})
	}
}

func GetProgram(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var result []model.Program

		// Mengambil data dari tabel "Program" dengan join ke tabel "departemen"
		if err := db.Table("Programs").
			Find(&result).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Memeriksa apakah slice "result" kosong
		if len(result) == 0 {
			c.JSON(404, gin.H{"message": "Tidak ada data Program"})
			return
		}

		c.JSON(200, result)
	}
}

func GetProgramByUUID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var Program struct {
			model.Program
		}
		if err := db.Table("Programs").
			Where("id = ?", id).
			First(&Program).
			Error; err != nil {
			c.JSON(404, gin.H{"error": "Program not found"})
			return
		}
		c.JSON(200, Program)
	}
}

func UpdateProgram(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mendapatkan UUID dari URL parameter
		id := c.Param("id")

		// Mencari Program berdasarkan UUID
		var existingProgram model.Program
		if err := db.Where("id = ?", id).First(&existingProgram).Preload("Doctor").Preload("Poly").Preload("Patient").Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Program not found"})
			return
		}

		// Binding input JSON ke struct Program
		var updatedProgram model.Program
		if err := c.ShouldBindJSON(&updatedProgram); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Memperbarui atribut Program
		db.Model(&existingProgram).Updates(&updatedProgram)

		// Mengambil data Program yang diperbarui dari database
		var updatedProgramFromDB model.Program
		if err := db.Where("id = ?", id).First(&updatedProgramFromDB).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated Program"})
			return
		}

		// Menyertakan pesan sukses dalam respons JSON
		c.JSON(http.StatusOK, model.ApiResponse{
			Message: "Program successfully updated",
			Data:    updatedProgramFromDB,
		})

	}
}

func DeleteProgram(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var Program model.Program
		if err := db.Where("id = ?", id).First(&Program).Error; err != nil {
			c.JSON(404, gin.H{"error": "Program not found"})
			return
		}

		ProgramName := Program.Name

		db.Delete(&Program)
		c.JSON(200, gin.H{"message": fmt.Sprintf("%s deleted", ProgramName)})
	}
}

func PermanentDeleteProgram(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var Program model.Program
		if err := db.Unscoped().Where("UUID = ?", id).First(&Program).Error; err != nil {
			c.JSON(404, gin.H{"error": "Program not found"})
			return
		}

		ProgramName := Program.Name

		db.Unscoped().Delete(&Program)
		c.JSON(200, gin.H{"message": fmt.Sprintf("%s permanently deleted", ProgramName)})
	}
}
