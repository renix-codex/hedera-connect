package topicid_handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/renix-codex/hedera-connect/internal/topicid"
)

var CreateTopicID = func(client *hedera.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		topic, err := topicid.CreateTopicID(c, client)
		if err != nil {
			log.Print("hedera client initialization failed")
			return
		}
		c.JSON(200, gin.H{
			"topicID": topic,
		})
	}
}
