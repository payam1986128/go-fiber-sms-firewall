package main

import (
	"github.com/payam1986128/go-fiber-sms-firewall/internal/db"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/handlers"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
)

func main() {
	// load .env
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, relying on environment variables")
	}

	app := fiber.New()
	app.Use(logger.New())

	// init couchbase
	if err := db.InitCouchbase(); err != nil {
		log.Fatalf("couchbase init: %v", err)
	}
	defer db.CloseCouchbase()

	// public route to log in (returns JWT for demo)
	app.Post("/login", handlers.LoginHandler)

	// secure group
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET must be set")
	}
	secure := app.Group("/api", jwtware.New(jwtware.Config{SigningKey: []byte(jwtSecret)}))

	// sms endpoints
	secure.Post("/sms", handlers.CreateSMS)
	secure.Get("/sms/:id", handlers.GetSMS)
	secure.Get("/sms", handlers.ListSMS)

	// rule endpoints
	secure.Post("/rules", handlers.CreateRule)
	secure.Get("/rules/:id", handlers.GetRule)
	secure.Get("/rules", handlers.ListRules)

	// evaluate an sms against rules (returns allow/deny and matched rules)
	secure.Post("/evaluate", handlers.EvaluateSMS)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}
