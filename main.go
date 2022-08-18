package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/renix-codex/hedera-connect/handlers"
	hedera_connect "github.com/renix-codex/hedera-connect/hedera"
)

func main() {
	//load env variables
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
		return
	}

	accountID, _ := os.LookupEnv("HEDERA_ACCOUNT_ID")
	privateKey, _ := os.LookupEnv("HEDERA_PRIVATE_KEY")
	mainnetEnabled, _ := os.LookupEnv("HEDERA_MAINNET_ENABLED")

	mainnetEnabledBool := false
	if mainnetEnabled == "true" {
		mainnetEnabledBool = true
	}

	//initialize Hedera client
	client, err := hedera_connect.InitializeClient(accountID, privateKey, mainnetEnabledBool)
	if err != nil {
		log.Print("hedera client initialization failed")
		return
	}

	//initialize gin gonic
	router := gin.Default()
	router.Use(gin.Recovery())
	handlers.InitializeRoutes(router, client)
	router.Run(":6000")
}
