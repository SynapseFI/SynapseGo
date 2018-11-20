# SynapseFI-Go
Simple API wrapper for SynapseFi v3 REST API.

# Installation
```bash
$ go get github.com/synapsefi/synapsefi-go
```

```go
import github.com/synapsefi/
```

# Code Examples

See main.go for each method in use

**queryParams** and **scope** are optional parameters

## Client

```go
// credentials used to set headers for each method request
var client = wrapper.Client(
  "CLIENT_ID|CLIENT_SECRET ",
  "IP_ADDRESS",
  "USER_ID"
)
```

To enable logging:

```go
var client = wrapper.Client(
  "CLIENT_ID|CLIENT_SECRET ",
  "IP_ADDRESS",
  "USER_ID",
  // pass true as fourth variable
  true
)
```

## Misc

### Institutions

```go
data := client.GetInstitutions()
```

## Nodes

```go
data := client.GetNodes(userID string, queryParams map[string]interface{})
```

## Subscriptions

```go
data := client.GetSubscriptions(queryParams map[string]interface{})
data := client.GetSubscription(subID string, queryParams map[string]interface{})
data := client.CreateSubscription(data string, queryParams map[string]interface{})
data := client.UpdateSubscription(subID string, query)
```

## Transactions

```go
data := client.GetClientTransactions(queryParams map[string]interface{})
```

## User Interface

### Issue Public Key

```go
data := client.GetPublicKey(scope string)
```

## Users

```go
data := client.GetUsers(queryParams map[string]interface{})
data := client.GetUser(userID string, fullDehydrate bool, queryParams map[string]interface{})
data := client.CreateUser(data string, queryParams map[string]interface{})
```
