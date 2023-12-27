package router

import (
	"main/handler"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetStuntingRoutes(r *gin.RouterGroup, db *gorm.DB) {
	// authMiddleware := middleware.AuthMiddleware(db)
	// adminMiddleware := middleware.RoleMiddleware("admin")
	stuntingGroup := r.Group("/")
	{
		stuntingGroup.GET("stunting", handler.GetStunting(db))
		stuntingGroup.GET("stunting/:uuid", handler.GetStuntingByUUID(db))
		stuntingGroup.POST("stunting", handler.CreateStunting(db))
		stuntingGroup.PUT("stunting/:uuid", handler.UpdateStunting(db))
		stuntingGroup.DELETE("stunting/:uuid", handler.DeleteStunting(db))
		stuntingGroup.DELETE("permanent-stunting/:uuid", handler.PermanentDeleteStunting(db))
	}
}
