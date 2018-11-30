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

## Client (developer)

```go
// credentials used to set headers for each method request
var client = wrapper.Client(
"clientID":     "CLIENT_ID",
"clientSecret": "CLIENT_SECRET",
"ipAddress":    "IP_ADDRESS",
"userID":       "USER_ID"
)
```

To enable logging (development mode):

```go
var client = wrapper.Client(
"clientID":     "CLIENT_ID",
"clientSecret": "CLIENT_SECRET",
"ipAddress":    "IP_ADDRESS",
"userID":       "USER_ID",
"devMode":      true
)
```

### Misc

#### Institutions

```go
data := client.GetInstitutions()
```

### Nodes

```go
data := client.GetNodes(queryParams ...string)
```

### Subscriptions

```go
data := client.GetSubscriptions(queryParams ...string)
data := client.GetSubscription(subID string, queryParams ...string)
data := client.CreateSubscription(data string, queryParams ...string)
data := client.UpdateSubscription(subID string, queryParams ...string)
```

### Transactions

```go
data := client.GetClientTransactions(queryParams ...string)
```

### User Interface

#### Issue Public Key

```go
data := client.GetPublicKey(scope ...string)
```

### Users

```go
data := client.GetUsers(queryParams ...string)

// instantiate User object
user := client.GetUser(userID string, fullDehydrate bool, queryParams ...string)
user := client.CreateUser(data string, queryParams ...string)
```

## User

```go
user := user.Update(data string)
```

### Authentication

```go
ak := user.Auth(data string)
```

### Documents

```go
user := user.AddNewDocuments(data string)
user := user.UpdateExistingDocument(data string)
user := user.DeleteExistingDocument(data string)
```

### Nodes

```go
user := user.GetNodes(queryParams ...string)
user := user.CreateDepositeNode(data string)
```

## Node

```go
node := node.ShipDebitCard(data string)
node := node.ResetDebitCard()

res:= node.DummyTransactions(credit bool)

tran := node.GetTransaction(transactionID string)
tran := node.CreateTranscation(transactionID data string)
```

## Transactions

```go
tran := tran.CommentOnStatus(data string)
tran := tran.CancelTransaction(data string)
```