package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/handler"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/repository"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/service"
	"log"
	"os"
)

func InitFiber(config *CouchbaseConfig) {
	app := fiber.New()
	app.Use(logger.New())

	userRepository := repository.NewUserRepository(config)
	smsRepository := repository.NewSmsRepository(config)
	limiterConditionRepository := repository.NewLimiterConditionRepository(config)
	suspiciousWordRepository := repository.NewSuspiciousWordRepository(config)
	suspiciousCategoryRepository := repository.NewSuspiciousCategoryRepository(config)

	userService := service.NewUserService(userRepository)
	smsService := service.NewSmsService(smsRepository)
	limiterConditionService := service.NewLimiterConditionService(limiterConditionRepository)
	suspiciousWordService := service.NewSuspiciousWordService(suspiciousWordRepository)
	suspiciousCategoryService := service.NewSuspiciousCategoryService(suspiciousCategoryRepository)
	phoneNumberService := service.NewPhoneNumberService()
	transceiverService := service.NewTransceiverService(smsRepository, phoneNumberService)
	rateLimiterService := service.NewRateLimiterService(smsRepository)
	firewallService := service.NewFirewallService(smsRepository, rateLimiterService, limiterConditionService)

	userHandler := handler.NewUserHandler(userService)
	smsHandler := handler.NewSmsHandler(smsService)
	limiterConditionHandler := handler.NewLimiterConditionHandler(limiterConditionService)
	suspiciousWordHandler := handler.NewSuspiciousWordHandler(suspiciousWordService)
	suspiciousCategoryHandler := handler.NewSuspiciousCategoryHandler(suspiciousCategoryService)
	firewallHandler := handler.NewFirewallHandler(firewallService, transceiverService)

	app.Post("/login", userHandler.LoginHandler)

	// secure group
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET must be set")
	}
	secure := app.Group("/api", jwtware.New(jwtware.Config{SigningKey: []byte(jwtSecret)}))

	secure.Post("/sms", firewallHandler.Receive)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}
