// main.go
package main

import (
	"fmt"
	"main/database"
	"main/router"
	"net/url"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	db := database.InitDB()
	defer db.Close()

	db.LogMode(true)

	// migration.Migrate(db)
	// seeder.Seed(db)

	r := gin.Default()

	// Aktifkan middleware CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:8000/"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			// Check if the origin is in the range 3000 to 3999
			originURL, err := url.Parse(origin)
			if err != nil {
				return false
			}

			port, err := strconv.Atoi(originURL.Port())
			if err != nil {
				return false
			}

			return port >= 8000 && port <= 8999
		},
	}))
	// Group router for authentication
	authRouter := r.Group("/")
	router.SetAuthRoutes(authRouter, db)

	// Group router untuk barang
	stuntingRouter := r.Group("/dinkes")
	router.SetStuntingRoutes(stuntingRouter, db)

	// Group router untuk barang
	ProgramRouter := r.Group("/dinkes/")
	router.SetProgramRoutes(ProgramRouter, db)

	r.Run(":8080")
}
