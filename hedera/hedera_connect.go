package hedera_connect

import (
	"github.com/gin-gonic/gin"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

//initializeClient will create a client and set operator using accountID and private key
func InitializeClient(accountID, privateKey string, mainnetEnabled bool) (*hedera.Client, error) {

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
