package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

	app.Post("/api/bo/users", userHandler.Register)
	app.Post("/login", userHandler.Login)

	// secure group
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET must be set")
	}
	secure := app.Group("/api", JWTMiddleware())

	secure.Post("/protections", firewallHandler.Receive)

	secure.Get("/bo/limiter-conditions", limiterConditionHandler.GetLimiterConditions)
	secure.Get("/bo/limiter-conditions/:id", limiterConditionHandler.GetLimiterCondition)
	secure.Post("/bo/limiter-conditions", limiterConditionHandler.AddLimiterCondition)
	secure.Put("/bo/limiter-conditions/:id", limiterConditionHandler.EditLimiterCondition)
	secure.Put("/bo/limiter-conditions/state", limiterConditionHandler.ReviewLimiterCondition)
	secure.Delete("/bo/limiter-conditions/:id", limiterConditionHandler.DeleteLimiterCondition)
	secure.Delete("/bo/limiter-conditions/:id/sms", limiterConditionHandler.GetCaughtSms)

	secure.Get("/bo/sms", smsHandler.GetSms)

	secure.Get("/bo/suspicious-categories", suspiciousCategoryHandler.GetSuspiciousCategories)
	secure.Post("/bo/suspicious-categories", suspiciousCategoryHandler.AddSuspiciousCategory)
	secure.Put("/bo/suspicious-categories/:id", suspiciousCategoryHandler.EditSuspiciousCategory)
	secure.Delete("/bo/suspicious-categories/:id", suspiciousCategoryHandler.DeleteSuspiciousCategory)

	secure.Get("/bo/suspicious-words", suspiciousWordHandler.GetSuspiciousWords)
	secure.Post("/bo/suspicious-words", suspiciousWordHandler.AddSuspiciousWords)
	secure.Delete("/bo/suspicious-words/:id", suspiciousWordHandler.DeleteSuspiciousWords)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}
