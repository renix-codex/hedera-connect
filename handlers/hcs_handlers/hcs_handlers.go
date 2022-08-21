package hcs_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/renix-codex/hedera-connect/hedera/hcs"
)

var CreateTopicID = func(client *hedera.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		topic, err := hcs.CreateTopicID(c, client)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed to create topic-id"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"topicID": topic,
		})
	}
}

var SubmitMessage = func(client *hedera.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		topicid := c.Param("topicid")
		htx := &hcs.HCSTx{}
		err := c.ShouldBindJSON(&htx)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "cannot bind the request"})
			return
		}

		tx, err := hcs.HCSSubmitMessage(c, client, topicid, htx.Message)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "hedera consenses service failed to submit message"})
			return
		}
		c.JSON(http.StatusOK, tx)
	}
}
