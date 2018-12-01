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

## Client

```go
// credentials used to set headers for each method request
var client = wrapper.New(
"clientID":     "CLIENT_ID",
"clientSecret": "CLIENT_SECRET",
"ipAddress":    "IP_ADDRESS",
"userID":       "USER_ID"
)
```

To enable logging (development mode):

```go
var client = wrapper.New(
"clientID":     "CLIENT_ID",
"clientSecret": "CLIENT_SECRET",
"ipAddress":    "IP_ADDRESS",
"userID":       "USER_ID",
"devMode":      true
)
```

#### Get Institutions

```go
data := client.GetInstitutions()
```

#### Get Nodes

```go
data := client.GetNodes(queryParams ...string)
```

#### Get Public Key

```go
key := client.GetPublicKey(scope ...string)
```

#### Get Subscriptions

```go
data := client.GetSubscriptions(queryParams ...string)
```

#### Get Subscription

```go
subscription := client.GetSubscription(subscriptionID string, queryParams ...string)
```

#### Create Subscription

```go
subscription := client.CreateSubscription(data string, queryParams ...string)
```

#### Update Subscription

```go
subscription := client.UpdateSubscription(subscriptionID string, queryParams ...string)
```

#### Get Transactions

```go
data := client.GetTransactions(queryParams ...string)
```

#### Get Users

```go
data := client.GetUsers(queryParams ...string)
```

#### Get User

```go
user := client.GetUser(userID string, fullDehydrate bool, queryParams ...string)
```

#### Create User

```go
user := client.CreateUser(userID string, queryParams ...string)
```

## User

#### Authentication

```go
authKey := user.Auth(data string)
```

#### Get Nodes

```go
data := user.GetNodes(queryParams ...string)
```

#### Create Deposit Account

```go
node := user.CreateDepositAccount(data string)
```

#### Update User

```go
user := user.Update(data string, queryParams ...string)
```

## Transaction

#### Comment On Status

```go
transaction := transaction.CommentOnStatus(data string)
```

#### Cancel Transaction

```go
transaction := transaction.CancelTransaction(data string)
```
