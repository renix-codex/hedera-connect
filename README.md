# Hedera-Connect #

Hedera-Connect is developed using Go-SDK helps to connect with Hedera Hashgraph Network.
## Microservices ##
Microservices architecture is an approach in which a single application is composed of many loosely coupled and independently deployable smaller services.
## Gin ##
Gin is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter. If you need performance and good productivity. Gin is fast, lightweight and open source.
## Hedera Hashgraph ## 
Hashgraph is a distributed ledger technology that has been described as an alternative to blockchains. The hashgraph technology is currently patented, and the only authorised ledger is Hedera Hashgraph. The native cryptocurrency of the Hedera Hashgraph system is HBAR.
Hedera goes beyond blockchain for developers to create the next era of fast, fair, and secure applications.
### Hedera Consensus Service ###
Hedera Consensus Service (HCS) is a purpose-built tool for creating decentralized, auditable logs of immutable and timestamped events for web2 and web3 applications. Messages are submitted to the Hedera network for consensus, given a trusted timestamp, and fairly ordered. HCS is used by applications in production to track provenance across supply chains, log asset transfers between blockchain networks, count votes in a DAO, monitor IoT devices, and more.
## Application ##
Hedera client initialise once the service is up and running. Client initialisation defined in main.go. 
```
client, err := hedera_connect.InitializeClient(accountID, privateKey, mainnetEnabledBool)
	if err != nil {
		log.Print("hedera client initialization failed")
		return
	}
  ```
 ### APIs ###
 Create topic ID on HCS - ```/topic-id``` 
 No payload 
 Response ```{
	"topicID": "0.0.xxxxxxx"
}```
Status:```200```
 
 Submit Message on HCS - ```/submit-message/topic-id/0.0.xxxxxxx```
 Payload ```{
	"message":"hello world!"
}```
Response ```
	"topicID": "0.0.xxxxxxx",
	"topicSequenceNumber": 2,
	"transactionStatus": "SUCCESS"
}```
Status:```200```
