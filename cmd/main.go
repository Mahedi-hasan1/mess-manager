package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"mess-manager/internal/db"
	"mess-manager/internal/handler"
	"mess-manager/internal/middleware"
)

func init() {
	db.ConnectPostgresDB()
	db.AutoMigrateModels()
	//db.InitRedis()
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://soulvent-frontend.vercel.app", "https://soulvent.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	//unprotected
	r.GET("/", health)
	r.HEAD("/", health)
	r.POST("/users", handler.CreateUser)
	r.POST("/login", handler.Login)
	r.POST("/signup", handler.SignUP)

	//protected
	protected := r.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		//user routes
		protected.GET("/users", handler.GetUserByUsername)
		//post routes
		protected.POST("/meals", handler.AddMeal)
		protected.GET("/meals", handler.GetMealsByUsername)

	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Mass manager main service running on :" + port)
	if err := r.Run("0.0.0.0:" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
