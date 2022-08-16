package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/renix-codex/hedera-connect/handlers/topicid_handlers"
)

func InitializeRoutes(e *gin.Engine, client *hedera.Client) {
	e.POST("/topic-id", topicid_handlers.CreateTopicID(client))
}
