/*
Package synapse is a wrapper library for the Synapse API (https://docs.synapsefi.com)

Instantiate client

	// credentials used to set headers for each method request
	var client = synapse.New(
	"CLIENT_ID",
	"CLIENT_SECRET",
	"IP_ADDRESS",
	"FINGERPRINT",
	)

# Examples

Enable logging & turn off developer mode (developer mode is true by default)

	var client = synapse.New(
	"CLIENT_ID",
	"CLIENT_SECRET",
	"IP_ADDRESS",
	"FINGERPRINT",
	true,
	false,
	)

Register Fingerprint

	// payload response
	{
		"error": {
				"en": "Fingerprint not registered. Please perform the MFA flow."
		},
		"error_code": "10",
		"http_code": "202",
		"phone_numbers": [
				"developer@email.com",
				"901-111-2222"
		],
		"success": false
	}

	// Submit a valid email address or phone number from "phone_numbers" list
	res, err := user.Select2FA("developer@email.com")

	// MFA sent to developer@email.com
	res, err := user.VerifyPIN("123456")

Set an `IDEMPOTENCY_KEY` (for `POST` requests only)

	scopeSettings := `{
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

	idempotencyKey := `1234567890`

	data, err := client.CreateSubscription(scopeSettings, idempotencyKey)

Submit optional query parameters

	params := "per_page=3&page=2"

	data, err := client.GetUsers(params)
*/
package synapse

import (
	"github.com/mitchellh/mapstructure"
)

/********** GLOBAL VARIABLES **********/
var logMode = false
var developerMode = true

/********** TYPES **********/

type (
	// Client represents the credentials used by the developer to instantiate a client
	Client struct {
		ClientID     string
		ClientSecret string
		Fingerprint  string
		IP           string
		request      Request
	}
)

/********** METHODS **********/

func (c *Client) do(method, url, data string, queryParams []string) (map[string]interface{}, error) {
	var body []byte
	var err error

	switch method {
	case "GET":
		body, err = c.request.Get(url, queryParams)

	case "POST":
		body, err = c.request.Post(url, data, queryParams)

	case "PATCH":
		body, err = c.request.Patch(url, data, queryParams)

	case "DELETE":
		body, err = c.request.Delete(url)
	}

	return readStream(body), err
}

/********** CLIENT **********/

// New creates a client object
func New(clientID, clientSecret, fingerprint, ipAddress string, modes ...bool) *Client {
	log.info("========== CREATING CLIENT INSTANCE ==========")
	if len(modes) > 0 {
		if modes[0] == true {
			logMode = true
		}

		if len(modes) > 1 && modes[1] == false {
			developerMode = false
		}
	}

	request := Request{
		clientID:     clientID,
		clientSecret: clientSecret,
		fingerprint:  fingerprint,
		ipAddress:    ipAddress,
	}

	return &Client{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Fingerprint:  fingerprint,
		IP:           ipAddress,
		request:      request,
	}
}

/********** AUTHENTICATION **********/

// GetPublicKey returns a public key as a token representing client credentials
func (c *Client) GetPublicKey(scope ...string) (map[string]interface{}, error) {
	log.info("========== GET PUBLIC KEY ==========")
	url := buildURL(path["client"])
	defaultScope := "OAUTH|POST,USERS|POST,USERS|GET,USER|GET,USER|PATCH,SUBSCRIPTIONS|GET,SUBSCRIPTIONS|POST,SUBSCRIPTION|GET,SUBSCRIPTION|PATCH,CLIENT|REPORTS,CLIENT|CONTROLS"

	if len(scope) > 0 {
		defaultScope = scope[0]
	}

	qp := []string{"issue_public_key=YES&scope=" + defaultScope}

	if len(scope) > 1 {
		userId := scope[1]
		qp[0] += "&user_id=" + userId
	}

	return c.do("GET", url, "", qp)
}

/********** NODE **********/

// GetNodes returns all of the nodes
func (c *Client) GetNodes(queryParams ...string) (map[string]interface{}, error) {
	log.info("========== GET CLIENT NODES ==========")
	url := buildURL(path["nodes"])

	return c.do("GET", url, "", queryParams)
}

// GetTradeMarketData returns data on a stock based on its ticker symbol
func (c *Client) GetTradeMarketData(tickerSymbol string) (map[string]interface{}, error) {
	log.info("========== GET TRADE MARKET DATA ==========")
	url := buildURL(path["nodes"], "trade-market-watch")

	ts := []string{tickerSymbol}

	return c.do("GET", url, "", ts)
}

// GetNodeTypes returns available node types
func (c *Client) GetNodeTypes() (map[string]interface{}, error) {
	log.info("========== GET NODE TYPES ==========")
	url := buildURL(path["nodes"], "types")

	return c.do("GET", url, "", nil)
}

/********** OTHER **********/

// GetCryptoMarketData returns market data for cryptocurrencies
func (c *Client) GetCryptoMarketData() (map[string]interface{}, error) {
	log.info("========== GET CRYPTO MARKET DATA ==========")
	url := buildURL(path["nodes"], "crypto-market-watch")

	return c.do("GET", url, "", nil)
}

// GetCryptoQuotes returns all of the quotes for crypto currencies
func (c *Client) GetCryptoQuotes(queryParams ...string) (map[string]interface{}, error) {
	log.info("========== GET CRYPTO QUOTES ==========")
	url := buildURL(path["nodes"], "crypto-quotes")

	return c.do("GET", url, "", queryParams)
}

// GetInstitutions returns a list of all available banking institutions
func (c *Client) GetInstitutions() (map[string]interface{}, error) {
	log.info("========== GET INSTITUTIONS ==========")
	url := buildURL(path["institutions"])

	return c.do("GET", url, "", nil)
}

// LocateATMs returns a list of nearby ATMs
func (c *Client) LocateATMs(queryParams ...string) (map[string]interface{}, error) {
	log.info("========== LOCATE ATMS ==========")
	url := buildURL(path["nodes"], "atms")

	return c.do("GET", url, "", queryParams)
}

// VerifyAddress checks if an address if valid
func (c *Client) VerifyAddress(data string) (map[string]interface{}, error) {
	log.info("========== VERIFY ADDRESS ==========")
	url := buildURL("address-verification")

	return c.do("POST", url, data, nil)
}

// VerifyRoutingNumber checks and returns the bank details of a routing number
func (c *Client) VerifyRoutingNumber(data string) (map[string]interface{}, error) {
	log.info("========== VERIFY ROUTING NUMBER ==========")
	url := buildURL("routing-number-verification")

	return c.do("POST", url, data, nil)
}

/********** SUBSCRIPTION **********/

// GetSubscriptions returns all of the nodes associated with a user
func (c *Client) GetSubscriptions(queryParams ...string) (map[string]interface{}, error) {
	log.info("========== GET SUBSCRIPTIONS ==========")
	url := buildURL(path["subscriptions"])

	return c.do("GET", url, "", queryParams)
}

// GetSubscription returns a single subscription
func (c *Client) GetSubscription(subscriptionID string) (map[string]interface{}, error) {
	log.info("========== GET SUBSCRIPTION ==========")
	url := buildURL(path["subscriptions"], subscriptionID)

	return c.do("GET", url, "", nil)
}

// CreateSubscription creates a subscription and returns the subscription data
func (c *Client) CreateSubscription(data string, idempotencyKey ...string) (map[string]interface{}, error) {
	log.info("========== CREATE SUBSCRIPTION ==========")
	url := buildURL(path["subscriptions"])

	return c.do("POST", url, data, idempotencyKey)
}

// UpdateSubscription updates an existing subscription
func (c *Client) UpdateSubscription(subscriptionID string, data string) (map[string]interface{}, error) {
	log.info("========== UPDATE SUBSCRIPTION ==========")
	url := buildURL(path["subscriptions"], subscriptionID)

	return c.do("PATCH", url, data, nil)
}

// GetWebhookLogs returns all of the webhooks sent to a specific client
func (c *Client) GetWebhookLogs() (map[string]interface{}, error) {
	log.info("========== GET WEBHOOK LOGS ==========")
	url := buildURL(path["subscriptions"], "logs")

	return c.do("GET", url, "", nil)
}

/********** TRANSACTION **********/

// GetTransactions returns all client transactions
func (c *Client) GetTransactions(queryParams ...string) (map[string]interface{}, error) {
	log.info("========== GET CLIENT TRANSACTIONS ==========")
	url := buildURL(path["transactions"])

	return c.do("GET", url, "", queryParams)
}

/********** USER **********/

// GetUsers returns a list of users
func (c *Client) GetUsers(queryParams ...string) (map[string]interface{}, error) {
	log.info("========== GET CLIENT USERS ==========")
	url := buildURL(path["users"])

	return c.do("GET", url, "", queryParams)
}

// GetUser returns a single user
func (c *Client) GetUser(userID, fingerprint, ipAddress string, queryParams ...string) (*User, error) {
	log.info("========== GET USER ==========")
	url := buildURL(path["users"], userID)
	res, err := c.do("GET", url, "", queryParams)

	var user User
	mapstructure.Decode(res, &user)
	user.Response = res
	request := Request{
		clientID:     c.ClientID,
		clientSecret: c.ClientSecret,
		fingerprint:  fingerprint,
		ipAddress:    ipAddress,
	}
	user.request = request

	return &user, err
}

// CreateUser creates a single user and returns the new user data
func (c *Client) CreateUser(data, fingerprint, ipAddress string, idempotencyKey ...string) (*User, error) {
	log.info("========== CREATE USER ==========")
	var user User
	user.request = Request{
		clientID:     c.ClientID,
		clientSecret: c.ClientSecret,
		fingerprint:  fingerprint,
		ipAddress:    ipAddress,
	}

	url := buildURL(path["users"])
	res, err := user.do("POST", url, data, idempotencyKey)
	mapstructure.Decode(res, &user)
	user.Response = res

	return &user, err
}

// GetUserDocumentTypes returns available user document types
func (c *Client) GetUserDocumentTypes() (map[string]interface{}, error) {
	log.info("========== GET USER DOCUMENT TYPES ==========")
	url := buildURL(path["users"], "document-types")

	return c.do("GET", url, "", nil)
}

// GetUserEntityTypes returns available user entity types
func (c *Client) GetUserEntityTypes() (map[string]interface{}, error) {
	log.info("========== GET USER ENTITY TYPES ==========")
	url := buildURL(path["users"], "entity-types")

	return c.do("GET", url, "", nil)
}

// GetUserEntityScopes returns available user entity scopes
func (c *Client) GetUserEntityScopes() (map[string]interface{}, error) {
	log.info("========== GET USER ENTITY TYPES ==========")
	url := buildURL(path["users"], "entity-scopes")

	return c.do("GET", url, "", nil)
}
