package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/nvlhnn/url-shortener/internal/config"
	"github.com/nvlhnn/url-shortener/internal/db"
	"github.com/nvlhnn/url-shortener/internal/db/mysql"
	"github.com/nvlhnn/url-shortener/internal/handler"
	"github.com/nvlhnn/url-shortener/internal/middleware"
	"gorm.io/gorm"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}


func setupRoutes(app *fiber.App, db *gorm.DB, config *config.Config) {

	// middleware
	rateLimiter := middleware.NewLimiterMiddleware(&config.Limiter, &config.MemStore)
	app.Use(rateLimiter.LimiterMiddleware())


	// handler
	shortenHandler := &handler.ShortenHandler{
		DB: mysql.NewURLMysql(db),
	}
	redirectHandler := &handler.RedirectHandler{
		DB: mysql.NewURLMysql(db),
	}


	app.Post("/api/v1", shortenHandler.ShortenURL)
	app.Get("/:shortUrl", redirectHandler.RedirectUrl)

}


func main() {

	config := config.NewConfig()

	// init postgres db
	database := db.NewDatabase(&config.Database)
	db := database.Connect() 


	app := fiber.New()

	setupRoutes(app, db, config)

	app.Use(logger.New())
	log.Fatal(app.Listen(":8080"))

}