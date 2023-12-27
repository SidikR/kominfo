package router

import (
	"main/handler"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetProgramRoutes(r *gin.RouterGroup, db *gorm.DB) {
	// authMiddleware := middleware.AuthMiddleware(db)
	// adminMiddleware := middleware.RoleMiddleware("admin")
	programGroup := r.Group("/")
	{
		programGroup.GET("program", handler.GetProgram(db))
		programGroup.GET("program/:id", handler.GetProgramByUUID(db))
		programGroup.POST("program", handler.CreateProgram(db))
		programGroup.PUT("program/:id", handler.UpdateProgram(db))
		programGroup.DELETE("program/:id", handler.DeleteProgram(db))
		programGroup.DELETE("permanent-program/:id", handler.PermanentDeleteProgram(db))
	}
}
