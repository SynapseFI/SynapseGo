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

Examples

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
var developerMode = true
var logMode = false

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
	if len(modes) > 0 {
		if modes[0] == true {
			logMode = true
		}

		if len(modes) > 1 && modes[1] == false {
			developerMode = false
		}
	}

	request := Request{
		authKey:      "",
		clientID:     clientID,
		clientSecret: clientSecret,
		fingerprint:  fingerprint,
		ipAddress:    ipAddress,
	}

	log.info("Building new client...")
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
	url := buildURL(path["client"])
	defaultScope := "OAUTH|POST,USERS|POST,USERS|GET,USER|GET,USER|PATCH,SUBSCRIPTIONS|GET,SUBSCRIPTIONS|POST,SUBSCRIPTION|GET,SUBSCRIPTION|PATCH,CLIENT|REPORTS,CLIENT|CONTROLS"

	if len(scope) > 0 {
		defaultScope = scope[0]
	}

	url += defaultScope
	qp := []string{"?issue_public_key=YES&scope=" + defaultScope}

	log.info("Getting public key...")
	return c.do("GET", url, "", qp)
}

/********** NODE **********/

// GetNodes returns all of the nodes
func (c *Client) GetNodes(queryParams ...string) (map[string]interface{}, error) {
	url := buildURL(path["nodes"])

	log.info("Getting list of nodes...")
	return c.do("GET", url, "", queryParams)
}

// GetTradeMarketData returns data on a stock based on its ticker symbol
func (c *Client) GetTradeMarketData(tickerSymbol string) (map[string]interface{}, error) {
	url := buildURL(path["nodes"], "trade-market-watch")

	ts := []string{tickerSymbol}

	log.info("Getting trade market data for " + tickerSymbol)
	return c.do("GET", url, "", ts)
}

/********** OTHER **********/

// GetCryptoMarketData returns market data for cryptocurrencies
func (c *Client) GetCryptoMarketData() (map[string]interface{}, error) {
	url := buildURL(path["nodes"], "crypto-market-watch")

	log.info("Getting crypto market data...")
	return c.do("GET", url, "", nil)
}

// GetCryptoQuotes returns all of the quotes for crypto currencies
func (c *Client) GetCryptoQuotes(queryParams ...string) (map[string]interface{}, error) {
	url := buildURL(path["nodes"], "crypto-quotes")

	log.info("Getting crypto market quotes...")
	return c.do("GET", url, "", queryParams)
}

// GetInstitutions returns all of the nodes associated with a user
func (c *Client) GetInstitutions() (map[string]interface{}, error) {
	url := buildURL(path["institutions"])

	log.info("Getting list of institutions...")
	return c.do("GET", url, "", nil)
}

// LocateATMs returns a list of nearby ATMs
func (c *Client) LocateATMs(queryParams ...string) (map[string]interface{}, error) {
	url := buildURL(path["nodes"], "atms")

	log.info("Getting list of ATMs...")
	return c.do("GET", url, "", queryParams)
}

// VerifyAddress checks if an address if valid
func (c *Client) VerifyAddress(data string) (map[string]interface{}, error) {
	url := buildURL("address-verification")

	log.info("Verifying address...")
	return c.do("POST", url, data, nil)
}

// VerifyRoutingNumber checks and returns the bank details of a routing number
func (c *Client) VerifyRoutingNumber(data string) (map[string]interface{}, error) {
	url := buildURL("routing-number-verification")

	log.info("Retrieving routing number details")
	return c.do("GET", url, data, nil)
}

/********** SUBSCRIPTION **********/

// GetSubscriptions returns all of the nodes associated with a user
func (c *Client) GetSubscriptions(queryParams ...string) (map[string]interface{}, error) {
	url := buildURL(path["subscriptions"])

	log.info("Getting list of subscriptions...")
	return c.do("GET", url, "", queryParams)
}

// GetSubscription returns a single subscription
func (c *Client) GetSubscription(subscriptionID string) (map[string]interface{}, error) {
	url := buildURL(path["subscriptions"], subscriptionID)

	log.info("Getting subscription...")
	return c.do("GET", url, "", nil)
}

// CreateSubscription creates a subscription and returns the subscription data
func (c *Client) CreateSubscription(data string, idempotencyKey ...string) (map[string]interface{}, error) {
	url := buildURL(path["subscriptions"])

	log.info("Creating subscription...")
	return c.do("POST", url, data, idempotencyKey)
}

// UpdateSubscription updates an existing subscription
func (c *Client) UpdateSubscription(subscriptionID string, data string) (map[string]interface{}, error) {
	url := buildURL(path["subscriptions"], subscriptionID)

	log.info("Updating subscription...")
	return c.do("PATCH", url, data, nil)
}

// GetWebhookLogs returns all of the webhooks sent to a specific client
func (c *Client) GetWebhookLogs() (map[string]interface{}, error) {
	url := buildURL(path["subscriptions"], "logs")

	log.info("Getting webhook logs...")
	return c.do("GET", url, "", nil)
}

/********** TRANSACTION **********/

// GetTransactions returns all client transactions
func (c *Client) GetTransactions(queryParams ...string) (map[string]interface{}, error) {
	url := buildURL(path["transactions"])

	log.info("Getting list of transactions...")
	return c.do("GET", url, "", queryParams)
}

/********** USER **********/

// GetUsers returns a list of users
func (c *Client) GetUsers(queryParams ...string) (map[string]interface{}, error) {
	url := buildURL(path["users"])

	log.info("Getting list of users...")
	return c.do("GET", url, "", queryParams)
}

// GetUser returns a single user
func (c *Client) GetUser(userID string, fingerprint, ipAddress string, queryParams ...string) (*User, error) {
	url := buildURL(path["users"], userID)
	res, err := c.do("GET", url, "", queryParams)

	var user User
	mapstructure.Decode(res, &user)
	user.Response = res
	user.request = Request{
		authKey:      "",
		clientID:     c.ClientID,
		clientSecret: c.ClientSecret,
		fingerprint:  fingerprint,
		ipAddress:    ipAddress,
	}

	log.info("Getting user...")
	return &user, err
}

// CreateUser creates a single user and returns the new user data
func (c *Client) CreateUser(data, fingerprint, ipAddress string, idempotencyKey ...string) (*User, error) {
	var user User
	user.request = Request{
		"",
		c.ClientID,
		c.ClientSecret,
		fingerprint,
		ipAddress,
	}
	url := buildURL(path["users"])
	res, err := user.do("POST", url, data, idempotencyKey)

	mapstructure.Decode(res, &user)
	user.Response = res

	log.info("Creating user...")
	return &user, err
}
