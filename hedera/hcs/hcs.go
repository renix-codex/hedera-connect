package hcs

import (
	"github.com/gin-gonic/gin"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

type HCSTx struct {
	Message             string `json:"message,omitempty"`
	TopicID             string `json:"topicID,omitempty"`
	TopicSequenceNumber uint64 `json:"topicSequenceNumber,omitempty"`
	TransactionStatus   string `json:"transactionStatus,omitempty"`
}

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

//HCSSubmitMessage submits a message to topicID
func HCSSubmitMessage(c *gin.Context, client *hedera.Client, hcsTopicID string, hcsMessage string) (*HCSTx, error) {

	topicID, err := hedera.TopicIDFromString(hcsTopicID)
	if err != nil {
		return nil, gin.Error{Err: err}
	}
	//Submit transaction
	submitMessage, err := hedera.NewTopicMessageSubmitTransaction().
		SetMessage([]byte(hcsMessage)).
		SetTopicID(topicID).
		Execute(client)
	if err != nil {
		return nil, gin.Error{Err: err}
	}

	//Get the receipt of the transaction
	receipt, err := submitMessage.GetReceipt(client)
	if err != nil {
		return nil, gin.Error{Err: err}
	}

	tx := &HCSTx{}
	tx.TopicID = topicID.String()
	tx.TopicSequenceNumber = receipt.TopicSequenceNumber
	tx.TransactionStatus = receipt.Status.String()

	return tx, nil
}
