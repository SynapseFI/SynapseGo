# Table of Contents
- [General](general-examples)
  * [Query Parameters](#query-parameters)
- [Client Examples](#client)
  * [Initialize Client](#initialize-client)
  * [Get Client Nodes](#get-client-nodes)
  * [Issue Public Key](#issue-public-key)
  * [Get Client Subscriptions](#get-client-subscriptions)
  * [Get Subscription](#get-subscription)
  * [Create Subscription](#create-subscription)
  * [Update Subscription](#update-subscription)
  * [Get Client Transactions](#get-client-transactions)
  * [Get Client Users](#get-client-users)
  * [Get User](#get-user)
  * [Create User](#create-user)
- [User Examples](#user)
  + [Authentication](#authentication)
    * [Get New Oauth](#get-new-oauth)
    * [Register Fingerprint](#register-fingerprint)
  + [Nodes](#nodes)
    * [Get User Nodes](#get-user-nodes)
    * [Get Node](#get-node)
    + [Create Node](#create-node)
        * [Bank Login w/ MFA](#bank-login)
    * [Update Node](#update-node)
    * [Delete Node](#delete-node)
    * [Get Node Subnets](#get-node-subnets)
    * [Get Node Transactions](#get-node-transactions)
    * [Ship Card Node](#ship-card-node)
    * [Reset Card Node](#reset-card-node)
    * [Verify Micro Deposit](#verify-micro-deposit)
    * [Reinitiate Micro Deposit](#reinitiate-micro-deposit)
    * [Get Apple Pay Token](#get-apple-pay-token)
  + [Subnets](#subnets)
    * [Get Subnet](#get-subnet)
    * [Create Subnet](#create-subnet)
    * [Update Subnet](#update-subnet)
    * [Ship Card](#ship-card)
  + [Transactions](#transactions)
    * [Get Transactions](#get-transactions)
    * [Get Transaction](#get-transaction)
    * [Create Transaction](#create-transaction)
    * [Comment on Transaction Status](#comment-on-transaction-status)
    * [Dispute Transaction](#dispute-transaction)
    * [Cancel Transaction](#cancel-transaction)
  + [Users](#users)
    * [Update User or Update/Add Documents](#update-user-or-update-add-documents)
    * [Generate UBO](#generate-ubo)

## General Examples

#### Query Parameters
Query parameters must be of type `string` and follow the following pattern:
`key=value&key=value&key=value`

```go
// Get Users with query parameters
data, err := client.GetUsers("per_page=3&page=2")

// Get User with full dehydrate
user, err := client.GetUser("5bec6ebebaabfc00ab168fa0", client.IP, client.Fingerprint, "full_dehydrate=yes")
```

## Client Examples

#### Initialize Client
```go
// credentials used to set headers for each method request
var client = synapse.New(
"client_id_1239ABCdefghijk1092312309",
"client_secret_1239ABCdefghijk1092312309",
"1.2.3.132",
"1023918209480asdf8341098",
)
```

Enable logging & turn off developer mode (developer mode is true by default)

```go
var client = synapse.New(
	var client = synapse.New(
	"CLIENT_ID",
	"CLIENT_SECRET",
	"IP_ADDRESS",
	"FINGERPRINT",
	true,
	false,
	)
)
```

#### Get Client Nodes
```go
data, err := client.GetNodes()
```

#### Issue Public Key
```go
scope := "OAUTH|POST,USERS|POST,USERS|GET,USER|GET,USER|PATCH"

data, err := client.GetPublicKey(scope)
```

```go
data, err := client.GetCryptoMarketData()
data, err := client.GetCryptoQuotes()
data, err := client.GetInstitutions()
data, err := client.LocateATMs()
```

#### Get Client Subscriptions
```go
data, err := client.GetSubscriptions()
```

#### Get Subscription
```go
subsID := "589b6adec83e17002122196c"

data, err := client.GetSubscription(subsID)
```

#### Create Subscription 
```go
body := `{
  "scope": [
    "USERS|POST",
    "USER|PATCH",
    "NODES|POST",
    "NODE|PATCH",
    "TRANS|POST",
    "TRAN|PATCH"
  ],
  "url": "https://requestb.in/zp216zzp"
}`

idempotencykey := "123456789"

data, err := client.CreateSubscription(body, idempotencyKey)
```

#### Update Subscription
```go
subID := "589b6adec83e17002122196c"
body := `{
    "scope": [
        "USERS|POST",
        "USER|PATCH",
        "NODES|POST",
        ...
      ]
    }`

data, err := client.UpdateSubscription(subID, body)
```

#### Get Client Transactions
```go
data, err := client.GetTransactions()
```

#### Get Client Users 
```go
data, err := client.GetUsers()
```

#### Get User
```go
// set FullDehydrate to true
userID = "594e0fa2838454002ea317a0"
userIP = "127.0.0.1" // or client.IP
userFingerprint = "TEST_FINGERPRINT" // or client.Fingerprint

user, err := client.GetUser(userID, userIP, userFingerprint)
```

#### Create User
```go
body := `{
  "logins": [
    {
      "email": "test@synapsefi.com"
    }
  ],
  "phone_numbers": [
    "901.111.1111",
    "test@synapsefi.com"
  ],
  "legal_names": [
    "Test User"
  ],
  ...
}`

userIP = "127.0.0.1" // or client.IP
userFingerprint = "TEST_FINGERPRINT" // or client.Fingerprint

user, err := client.CreateUser(body, userIP, userFingerprint)
```

#### Get Webhook Logs
```go
data, err := client.GetWebhookLogs()
```

## User Examples

### Authentication

#### Get New Oauth
```go
body := `{
    "refresh_token":"refresh_Y5beJdBLtgvply3KIzrh72UxWMEqiTNoVAfDs98G",
    "scope":[
        "USER|PATCH",
        "USER|GET",
        ...
    ]
}`

data, err := user.Authenticate(body)
```

#### Register Fingerprint
```go
/*
{
	"error": {
		en": "Fingerprint not registered. Please perform the MFA flow."
	},
	"error_code": "10",
	"http_code": "202",
	"phone_numbers": [
		"developer@email.com",
		"901-111-2222"
	],
	"success": false
}
*/

// Submit a valid email address or phone number from "phone_numbers" list
res, err := user.Select2FA("developer@email.com")

// MFA sent to developer@email.com
res, err := user.VerifyPIN("123456")

```

### Nodes

#### Get User Nodes
```go
data, err := user.GetNodes()
```

#### Get Node
```go
nodeID := "594e606212e17a002f2e3251"

data, err := user.GetNode(nodeID)
```

#### Create Node
Refer to the following docs for how to setup the payload for a specific Node type:
- [Deposit Accounts](https://docs.synapsefi.com/v3.1/docs/deposit-accounts)
- [Card Issuance](https://docs.synapsefi.com/v3.1/docs/card-issuance)
- [ACH-US with Logins](https://docs.synapsefi.com/v3.1/docs/add-ach-us-node)
- [ACH-US MFA](https://docs.synapsefi.com/v3.1/docs/add-ach-us-node-via-bank-logins-mfa)
- [ACH-US with AC/RT](https://docs.synapsefi.com/v3.1/docs/add-ach-us-node-via-acrt-s)
- [INTERCHANGE-US](https://docs.synapsefi.com/v3.1/docs/interchange-us)
- [CHECK-US](https://docs.synapsefi.com/v3.1/docs/check-us)
- [CRYPTO-US](https://docs.synapsefi.com/v3.1/docs/crypto-us)
- [WIRE-US](https://docs.synapsefi.com/v3.1/docs/add-wire-us-node)
- [WIRE-INT](https://docs.synapsefi.com/v3.1/docs/add-wire-int-node)
- [IOU](https://docs.synapsefi.com/v3.1/docs/add-iou-node)

```go
body := `{
  "type": "DEPOSIT-US",
  "info":{
      "nickname":"My Checking"
  }
}`

data, err := user.CreateNode(body)
```

##### Bank Login w/ MFA
```go
body := `{
  "type": "ACH-US",
  "info":{
    "bank_id":"synapse_good",
    "bank_pw":"test1234",
    "bank_name":"fake"
  }
}`

data, err := user.CreateNode(body)

// parse `access_token` from `data`

// create MFA answer body

mfaBody := `{
  "access_token":"fake_cd60680b9addc013ca7fb25b2b704ba82d3",
  "mfa_answer":"test_answer"
}`

achData, achErr := user.SubmitMFA(mfaBody)
```

#### Update Node
```go
nodeID := "5ba05ed620b3aa005882c52a"
body := `{
  "supp_id":"new_supp_id_1234"
}`

data, err := user.UpdateNode(nodeID, body)
```

#### Delete Node
```go
nodeID := "594e606212e17a002f2e3251"

data, err := user.DeleteNode(nodeID)
```

#### Get Node Subnets
```go
nodeID := "594e606212e17a002f2e3251"

data, err := user.GetNodeSubnets(nodeID, "page=4&per_page=10")
```
#### Get Node Transactions
```go
nodeID := "594e606212e17a002f2e3251"

data, err := user.GetNodeTransactions(nodeID, "page=4&per_page=10")
```

#### Ship Card Node
```go
nodeID := "5ba05ed620b3aa005882c52a"
body := `{
  "fee_node_id":"5ba05e7920b3aa006482c5ad",
  "expedite": True
}`

data, err := user.ShipCardNode(nodeID, body)
```

#### Reset Card Node
```go
nodeID := "5ba05ed620b3aa005882c52a"

data, err := user.ResetCardNode(nodeID)
```

#### Verify Micro Deposit
```go
nodeID := "5ba05ed620b3aa005882c52a"
body := `{
  "micro":[0.1,0.1]
}`

data, err := user.VerifyMicroDeposit(nodeID, body)
```

#### Reinitiate Micro Deposit
```go
nodeID := "5ba05ed620b3aa005882c52a"

data, err := user.ReinitiateMicroDeposit(nodeID)
```

#### Get Apple Pay Token
```go
nodeID := "5ba05ed620b3aa005882c52a"
body = `{
  "certificate": "your applepay cert",
  "nonce": "9c02xxx2",
  "nonce_signature": "4082f883ae62d0700c283e225ee9d286713ef74"
}`

data, err := user.GenerateApplePayToken(nodeID, body)
```

### Subnets

#### Get Subnet
```go
nodeID := "594e606212e17a002f2e3251"
subID := "59c9f77cd412960028b99d2b"

data, err := user.GetSubnet(nodeID, subID)
```

#### Create Subnet
```go
nodeID := "594e606212e17a002f2e3251"
body := `{
  "nickname":"Test AC/RT"
}`

data, err := user.CreateSubnet(nodeID, body)
```

#### Update Subnet
```go
nodeID := "594e606212e17a002f2e3251"
subnetID := "5bc920f2fff373002bf0d51b"
body := `{
  "preferences": {
    "allow_foreign_transactions":true,
    "daily_atm_withdrawal_limit":10,
    "daily_transaction_limit":1000
  }
}`

data, err := user.UpdateSubnet(nodeID, subnetID, body)
```

#### Ship Card
```go
nodeID := "594e606212e17a002f2e3251"
subnetID := "5bc920f2fff373002bf0d51b"
body := `{
  "fee_node_id":"5bba781485411800991b606b",
  "expedite":false,
  "card_style_id":"555"
}`

data, err := user.ShipCard(nodeID, subnetID, body)
```

### Transactions

#### Get Transactions
```go
nodeID := "594e606212e17a002f2e3251"

data, err := user.GetTransactions(nodeID)
```

#### Get Transaction
```go
nodeID := "594e606212e17a002f2e3251"
transID := "594e72124599e8002fe62e4f"

data, err := user.GetTransactions(nodeID, transID)
```

#### Create Transaction
```go
nodeID := "594e606212e17a002f2e3251"
body := `{
  "to": {
    "type": "ACH-US",
    "id": "594e6e6c12e17a002f2e39e4"
  },
  "amount": {
    "amount": 20.1,
    "currency": "USD"
  },
  "extra": {
    "ip": "192.168.0.1"
  }
}`

data, err := user.CreateTransaction(nodeID, body)
```

#### Comment on Transaction Status
```go
nodeID := "594e606212e17a002f2e3251"
transID := "594e72124599e8002fe62e4f"

data, err := user.CommentOnTransactionStatus(nodeID, transID, "Pending verification...")
```

#### Dispute Transaction
```go
nodeID := "594e606212e17a002f2e3251"
transID := "594e72124599e8002fe62e4f"
body := `{
  "dispute_reason": "CHARGE_BACK"
}`

data, err := user.DisputeTransaction(nodeID, transID, body)
```
#### Cancel Transaction
```go
nodeID := "594e606212e17a002f2e3251"
transID := "594e72124599e8002fe62e4f"

data, err := user.CancelTransaction(nodeID, transID)
```

### Users

#### Update User or Update/Add Documents
```go
body := `{
  "update":{
    "login":{
      "email":"test2@synapsefi.com"
    },
    "remove_login":{
      "email":"test@synapsefi.com"
    },
    "phone_number":"901-111-2222",
    "remove_phone_number":"901.111.1111"
    }
}`

data, err := user.Update(body)
```

#### Generate UBO
```go
body := `{
   "entity_info": {
      "cryptocurrency": True,
      "msb": {
         "federal": True,
         "states": ["AL"]
      },
      "public_company": False,
      "majority_owned_by_listed": False,
      "registered_SEC": False,
      "regulated_financial": False,
      "gambling": False,
      "document_id": "2a4a5957a3a62aaac1a0dd0edcae96ea2cdee688ec6337b20745eed8869e3ac8"
   ...
}`

data, err := user.CreateUBO(body)
```