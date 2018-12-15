# SynapseFI-Go
Simple API wrapper for SynapseFi v3 REST API.

# Installation
```bash
$ go get github.com/synapsefi/synapsefi-go
```

```go
import github.com/synapsefi/synapsefi-go
```

# Code Examples

See main.go for each method in use

*queryParams* and *scope* are optional parameters

## CLIENT METHODS

```go
// credentials used to set headers for each method request
var client = wrapper.New(
"ClientID":     "CLIENT_ID",
"ClientSecret": "CLIENT_SECRET",
"IP":           "IP_ADDRESS",
"Fingerprint":  "FINGERPRINT"
)
```

To enable logging (development mode):

```go
var client = wrapper.New(
"ClientID":     "CLIENT_ID",
"ClientSecret": "CLIENT_SECRET",
"IP":           "IP_ADDRESS",
"Fingerprint":  "FINGERPRINT",
"devMode":      true
)
```

#### Node

```go
data, err := client.GetNodes(queryParams ...string)
```

##### Other

```go
data, err := client.GetCryptoMarketData()
data, err := client.GetCryptoQuotes(queryParams ...string)
data, err := client.GetInstitutions()
data, err := client.LocateATMs(queryParams ...string)
data, err := client.GetPublicKey(scope ...string)
```

#### Subscription

```go
data, err := client.GetSubscriptions(queryParams ...string)
data, err := client.GetSubscription(subscriptionID string, queryParams ...string)
data, err := client.CreateSubscription(data string, queryParams ...string)
data, err := client.UpdateSubscription(subscriptionID string, queryParams ...string)

```

#### Transaction

```go
data, err := client.GetTransactions(queryParams ...string)
```

#### User

```go
data, err := client.GetUsers(queryParams ...string)
user, err := client.GetUser(userID string, fullDehydrate bool, queryParams ...string)
user, err := client.CreateUser(userID string, queryParams ...string)
```

## USER METHODS

#### Authentication

```go
data, err := user.Auth(data string)
data, err := user.GetRefreshToken()
data, err := user.Select2FA(device string)
data, err := user.SubmitMFA(data string)
data, err := user.VerifyPIN(pin string)
```

#### Node

```go
data, err := user.GetNodes(queryParams ...string)
data, err := user.GetNode(nodeID string, queryParams ..string)
data, err := user.CreateNode(data string)
data, err := user.UpdateNode(nodeID, data string)
data, err := user.DeleteNode(nodeID string)
```

#### Node (Other)

```go
data, err := user.GetApplePayToken(nodeID, data string)
data, err := user.ReinitiateMicroDeposit(nodeID, string)
data, err := user.ResetDebitCard(nodeID string)
data, err := user.ShipDebitCard(nodeID, data string)
data, err := user.TriggerDummyTransactions(nodeID string, credit bool)
data, err := user.VerifyMicroDeposit(nodeID, data string)
```

#### Statement

```go
data, err := user.GetNodeStatements(nodeID string, queryParams ...string)
data, err := user.GetStatements(queryParams ...string)
```

#### Subnet

```go
data, err := user.GetSubnets(nodeID string)
data, err := user.GetSubnet(nodeID, subnetID string)
data, err := user.CreateSubnet(nodeID, data string)
```

### Transaction

```go
data, err := user.GetTransactions(nodeID, transactionID string)
data, err := user.GetTransaction(nodeID, transactionID string)
data, err := user.CreateTransaction(nodeID, transactionID, data string)
data, err := user.DeleteTransaction(nodeID, transactionID string)
data, err := user.CommentOnTransactionStatus(nodeID, transactionID, data string)
data, err := user.DisputeTransaction(nodeID, transactionID string)
```

#### User

```go
user, err := user.Update(data string, queryParams ...string)

data, err := user.CreateUBO(data string)
```

