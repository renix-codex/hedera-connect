package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/renix-codex/hedera-connect/handlers/hcs_handlers"
)

func InitializeRoutes(e *gin.Engine, client *hedera.Client) {
	e.POST("/topic-id", hcs_handlers.CreateTopicID(client))
	e.POST("/submit-message/topic-id/:topicid", hcs_handlers.SubmitMessage(client))
}
