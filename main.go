package main

import (
	"github.com/joho/godotenv"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/config"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, relying on environment variables")
	}

	couchbaseConfig, err := config.InitCouchbase()
	if err != nil {
		log.Fatalf("couchbase init: %v", err)
	}

	config.InitFiber(couchbaseConfig)
}
