# Table of Contents
- [General Examples](#general-examples)
  * [Query Parameters](#query-parameters)
- [Client Examples](#client-examples)
  + [Client](#client)
    * [Initialize Client](#initialize-client)
  + [Authentication](#authentication-client)
    * [Get Public Key](#get-public-key)
  + [Nodes](#nodes-client)
    * [Get Client Nodes](#get-client-nodes)
    * [Get Trade Market Data](#get-trade-market-data)
  + [Other](#other-client)
    * [Get Crypto Market Data](#get-crypto-market-data)
    * [Get Crypto Quotes](#get-crypto-quotes)
    * [Get Institutions](#get-institutions)
    * [Locate ATMs](#locate-atms)
    * [Verify Routing Number](#verify-routing-number)
    * [Verify Address](#verify-address)
  + [Subscriptions](#subscriptions-client)
    * [Get Client Subscriptions](#get-client-subscriptions)
    * [Get Subscription](#get-subscription)
    * [Create Subscription](#create-subscription)
    * [Update Subscription](#update-subscription)
    * [Get Subscription Logs](#get-subscription-logs)
  + [Transactions](#transactions-client)
    * [Get Client Transactions](#get-client-transactions)
  + [Users](#users-client)
    * [Get Users](#get-users)
    * [Get User](#get-user)
    * [Create User](#create-user)
- [User Examples](#user)
  + [Authentication](#authentication)
    * [Authenticate](#authenticate)
    * [Get Refresh Token](#get-refresh-token)
    * [Register Fingerprint](#register-fingerprint)
    * [Select2FA](#select-2fa)
    * [SubmitMFA](#submit-mfa)
    * [VerifyPIN](#verify-pin)
  + [Nodes](#nodes)
    * [Get Nodes](#get-nodes)
    * [Get Node](#get-node)
    + [Create Node](#create-node)
        * [Bank Login w/ MFA](#bank-login)
    * [Update Node](#update-node)
    * [Delete Node](#delete-node)
    * [Verify Micro Deposit](#verify-micro-deposit)    
    * [Reinitiate Micro Deposits](#reinitiate-micro-deposits)
    * [Ship Card Node](#ship-card-node)
    * [Reset Card Node](#reset-card-node)   
    * [Get Apple Pay Token](#get-apple-pay-token)
  + [Statements](#statements)
    * [Get Statements](#get-statements)
    * [Get Node Statements](#get-node-statements)
    * [Create Node Statements](#create-node-statements)
  + [Subnets](#subnets)
    * [Get Subnets](#get-subnets)
    * [Get Node Subnets](#get-node-subnets)
    * [Get Subnet](#get-subnet)
    * [Create Subnet](#create-subnet)
    * [Update Subnet](#update-subnet)
    * [Ship Card](#ship-card)
  + [Transactions](#transactions)
    * [Get Transactions](#get-transactions)
    * [Get Node Transactions](#get-node-transactions)
    * [Get Transaction](#get-transaction)
    * [Create Transaction](#create-transaction)
    * [Cancel Transaction](#cancel-transaction)  
    * [Comment on Transaction Status](#comment-on-transaction-status)
    * [Dispute Transaction](#dispute-transaction)
    * [Create Dummy Transactions](#create-dummy-transactions)
  + [Users](#users)
    * [Update User or Update/Add Documents](#update-user-or-update-add-documents)
    * [Create UBO](#create-ubo)

## General Examples

#### Query Parameters
Query parameters must be of type `string` and follow the following pattern:
`key=value&key=value&key=value`

```go
// Get Users with query parameters
data, err := client.GetUsers("per_page=3&page=2")

// Get User with full dehydrate
user, err := client.GetUser("5bec6ebebaabfc00ab168fa0", "TEST_FINGERPRINT", "127.0.0.1", "full_dehydrate=yes")
```

## Client Examples

### Client

#### Initialize Client
```go
// credentials used to set headers for each method request
var client = synapse.New(
"client_id_1239ABCdefghijk1092312309",
"client_secret_1239ABCdefghijk1092312309",
"1023918209480asdf8341098",
"1.2.3.132",
)
```

Enable logging & turn off developer mode (developer mode is true by default)

```go
	var client = synapse.New(
	"CLIENT_ID",
  "CLIENT_SECRET",
	"FINGERPRINT",  
	"IP_ADDRESS",
	true,     // set to `false` to disable logging
	false,    // set to `true` to enable developer mode
	)
```

### Authentication (Client)

#### Get Public Key
```go
scope := "OAUTH|POST,USERS|POST,USERS|GET,USER|GET,USER|PATCH"

data, err := client.GetPublicKey(scope)
```

### Nodes (Client)

#### Get Client Nodes
```go
data, err := client.GetNodes()
```

#### Get Trade Market Data
```go
data, err := client.GetTradeMarketData("AAPL")
```

### Other (Client)

#### Get Crypto Market Data
```go
data, err := client.GetCryptoMarketData()
```

#### GetCryptoQuotes
```go
data, err := client.GetCryptoQuotes()
```

#### Get Institutions
```go
data, err := client.GetInstitutions()
```

#### Locate ATMs
```go
data, err = client.LocateATMs()
```

#### Verify Routing Number
```go
body := `{
  "routing_num": "084008426",
  "type": "ACH-US"
}`

data, err := client.VerifyRoutingNumber(body)
```

#### Verify Address
```go
address := `{
  "address_street": "170 St Germain St",
  "address_city": "SF",
  "address_subdivision": "CA",
  "address_country_code": "US",
  "address_postal_code": "94404"
}`
data, err := client.VerifyAddress(address)
```

### Subscriptions (Client)

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

#### Get Subscription Logs
```go
data, err := client.GetWebhookLogs()
```

### Transactions (Client)

#### Get Client Transactions
```go
data, err := client.GetTransactions()
```

### Users (Client)

#### Get Users 
```go
data, err := client.GetUsers()
```

#### Get User
```go
// set FullDehydrate to true
userID = "594e0fa2838454002ea317a0"
userFingerprint = "TEST_FINGERPRINT" // or client.Fingerprint
userIP = "127.0.0.1" // or client.IP

user, err := client.GetUser(userID, userFingerprint, userIP)
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

userFingerprint = "TEST_FINGERPRINT" // or client.Fingerprint
userIP = "127.0.0.1" // or client.IP

user, err := client.CreateUser(body, userFingerprint, userIP)
```

## User Examples

### Authentication

**To perform user actions, users must be authenticated**
- If performing user actions with a **new user** (i.e. created during the current session), further authentication is not necessary
- If performing user actions with any other user, please authenticate the user
- User authentication is only required at the start of a new session (`user.AuthKey == nil`), the wrapper will handle future authentication sessions

#### Authenticate
```go
body := `{
  "refresh_token":"refresh_Y5beJdBLtgvply3KIzrh72UxWMEqiTNoVAfDs98G",
  "scope":[
      "USER|PATCH",
      "USER|GET",
      ...
  ]
}`

userFingerprint = "TEST_FINGERPRINT" // or client.Fingerprint
userIP = "127.0.0.1" // or client.IP

data, err := user.Authenticate(body, userFingerprint, userIP)
```

#### Get Refresh Token
```go
data, err := user.GetRefreshToken()
```

#### Register Fingerprint
```go
// Submit a new fingerprint to be registered for the user
data, err := user.RegisterFingerprint("NEW_FINGERPRINT")
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
```

#### Select 2FA
```go
// Submit a valid email address or phone number from "phone_numbers" list
data, err := user.Select2FA("developer@email.com")
```

#### Submit MFA
```go
body := `{
  "access_token":"fake_cd60680b9addc013ca7fb25b2b704ba82d3",
  "mfa_answer":"test_answer"
}`
data, err := user.SubmitMFA(body)
```

#### Verify PIN
```go
// MFA sent to developer@email.com
data, err := user.VerifyPIN("123456")
```

### Nodes

#### Get Nodes
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

#### Verify Micro Deposit
```go
nodeID := "5ba05ed620b3aa005882c52a"
body := `{
  "micro":[0.1,0.1]
}`

data, err := user.VerifyMicroDeposit(nodeID, body)
```

#### Reinitiate Micro Deposits
```go
nodeID := "5ba05ed620b3aa005882c52a"

data, err := user.ReinitiateMicroDeposits(nodeID)
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

### Statements

#### Get Statements
```go
data, err := user.GetStatements()
```

#### Get Node Statements
```go
nodeID := "594e606212e17a002f2e3251"

data, err := user.GetNodeStatements(nodeID)
```

#### Create Node Statements
```go
nodeID := "5b4b2df145d1cc006d088f2e"
body := `{
  "date_start": 1525132800000,
  "date_end": 1525132800000,
  "webhook": "https://wh.synapsefi.com/gen_me_statement_001"
}`

data, err := user.CreateNodeStatements(nodeID, body)
```

### Subnets

#### Get Subnets
```go
data, err := user.GetSubnets()
```

#### Get Node Subnets
```go
nodeID := "594e606212e17a002f2e3251"

data, err := user.GetNodeSubnets(nodeID, "page=4&per_page=10")
```

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
data, err := user.GetTransactions()
```

#### Get Node Transactions
```go
nodeID := "594e606212e17a002f2e3251"

data, err := user.GetNodeTransactions(nodeID)
```

#### Get Transaction
```go
nodeID := "594e606212e17a002f2e3251"
transID := "594e72124599e8002fe62e4f"

data, err := user.GetTransaction(nodeID, transID)
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

#### Cancel Transaction
```go
nodeID := "594e606212e17a002f2e3251"
transID := "594e72124599e8002fe62e4f"

data, err := user.CancelTransaction(nodeID, transID)
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

#### Create Dummy Transaction
```go
nodeID := "594e606212e17a002f2e3251"
queryParams := "is_credit=YES&type=INTERCHANGE&subnetid=5caac735e1232a0029ee649c&foreign_transaction=NO"

data, err := user.CreateDummyTransaction(nodeID, queryParams)
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