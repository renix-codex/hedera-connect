package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/joho/godotenv"
)

//initializeClient will create a client and set operator using accountID and private key
func initializeClient(accountID, privateKey string, mainnetEnabled bool) (*hedera.Client, error) {

	accID, err := hedera.AccountIDFromString(accountID)
	if err != nil {
		return nil, gin.Error{Err: err}
	}

	pKey, err := hedera.PrivateKeyFromString(privateKey)
	if err != nil {
		return nil, gin.Error{Err: err}
	}

	client := &hedera.Client{}
	if mainnetEnabled {
		client = hedera.ClientForMainnet()
	} else {
		client = hedera.ClientForTestnet()
	}
	client.SetOperator(accID, pKey)

	return client, nil
}

func createTopicID(c *gin.Context, client *hedera.Client) (string, error) {
	transactionResponse, err := hedera.NewTopicCreateTransaction().Execute(client)
	if err != nil {
		return "", gin.Error{Err: err}
	}

	//Get the topic create transaction receipt
	transactionReceipt, err := transactionResponse.GetReceipt(client)
	if err != nil {
		return "", gin.Error{Err: err}
	}

	return transactionReceipt.TopicID.String(), nil
}

func main() {
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

	client, err := initializeClient(accountID, privateKey, mainnetEnabledBool)
	if err != nil {
		log.Print("hedera client initialization failed")
		return
	}

	router := gin.Default()
	router.GET("/topic-id", func(c *gin.Context) {
		topic, err := createTopicID(c, client)
		if err != nil {
			log.Print("hedera client initialization failed")
			return
		}
		c.JSON(200, gin.H{
			"topicID": topic,
		})
	})
	router.Run(":6000")
}
