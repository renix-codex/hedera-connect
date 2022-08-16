package topicid

import (
	"github.com/gin-gonic/gin"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

//createTopicID return new topicID
func CreateTopicID(c *gin.Context, client *hedera.Client) (string, error) {
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
