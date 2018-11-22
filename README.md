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
var credentials = map[string]interface{}{
"clientID":     "CLIENT_ID",
"clientSecret": "CLIENT_SECRET",
"ipAddress":    "IP_ADDRESS",
"userID":       "USER_ID",
}
var client = wrapper.Client(credentials)
```

To enable logging (development mode):

```go
var credentials = map[string]interface{}{
"clientID":     "CLIENT_ID",
"clientSecret": "CLIENT_SECRET",
"ipAddress":    "IP_ADDRESS",
"userID":       "USER_ID",
"devMode":      true,
}
var client = wrapper.Client(credentials)
```

### Misc

#### Institutions

```go
data := client.GetInstitutions()
```

### Nodes

```go
data := client.GetNodes(queryParams map[string]interface{})
```

### Subscriptions

```go
data := client.GetSubscriptions(queryParams map[string]interface{})
data := client.GetSubscription(subID string, queryParams map[string]interface{})
data := client.CreateSubscription(data string, queryParams map[string]interface{})
data := client.UpdateSubscription(subID string, query)
```

### Transactions

```go
data := client.GetClientTransactions(queryParams map[string]interface{})
```

### User Interface

#### Issue Public Key

```go
data := client.GetPublicKey(scope string)
```

### Users

```go
data := client.GetUsers(queryParams map[string]interface{})

// instantiate User object
user := client.GetUser(userID string, fullDehydrate bool, queryParams map[string]interface{})
user := client.CreateUser(data string, queryParams map[string]interface{})
```

## User

```go
data := user.Update(data string, bodyParams map[string]interface{})
```

### Authentication

```go
data := user.Auth(data string)
```

### Documents

```go
data := user.AddNewDocuments(data string)
data := user.UpdateExistingDocument(data string)
data := user.DeleteExistingDocument(data string)
```

### Nodes

```go
data := user.GetNodes(queryParams ...map[string]interface{})
data := user.CreateDepositeNode(data string)
```
