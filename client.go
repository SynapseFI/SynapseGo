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

	data, err := client.CreateSubscription(scopeSettings, "", idempotencyKey)

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

	c.request = c.request.updateRequest(c.ClientID, c.ClientSecret, c.Fingerprint, c.IP)

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
func New(clientID, clientSecret, ipAddress, fingerprint string, modes ...bool) *Client {
	if len(modes) > 0 {
		if modes[0] == true {
			logMode = true
		}

		if len(modes) > 1 && modes[1] == false {
			developerMode = false
		}
	}

	return &Client{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Fingerprint:  fingerprint,
		IP:           ipAddress,
	}
}

/********** NODE **********/

// GetNodes returns all of the nodes
func (c *Client) GetNodes(queryParams ...string) (map[string]interface{}, error) {
	return c.do("GET", path["nodes"], "", queryParams)
}

/********** OTHER **********/

// GetCryptoMarketData returns market data for cryptocurrencies
func (c *Client) GetCryptoMarketData() (map[string]interface{}, error) {
	url := buildURL(path["nodes"], "crypto-market-watch")

	return c.do("GET", url, "", nil)
}

// GetCryptoQuotes returns all of the quotes for crypto currencies
func (c *Client) GetCryptoQuotes(queryParams ...string) (map[string]interface{}, error) {
	url := buildURL(path["nodes"], "crypto-quotes")

	return c.do("GET", url, "", queryParams)
}

// GetInstitutions returns all of the nodes associated with a user
func (c *Client) GetInstitutions() (map[string]interface{}, error) {
	return c.do("GET", path["institutions"], "", nil)
}

// LocateATMs returns a list of nearby ATMs
func (c *Client) LocateATMs(queryParams ...string) (map[string]interface{}, error) {
	url := buildURL(path["nodes"], "atms")

	return c.do("GET", url, "", queryParams)
}

// GetPublicKey returns a public key as a token representing client credentials
func (c *Client) GetPublicKey(scope ...string) (map[string]interface{}, error) {
	url := path["client"] + "?issue_public_key=YES&scope="
	defaultScope := "OAUTH|POST,USERS|POST,USERS|GET,USER|GET,USER|PATCH,SUBSCRIPTIONS|GET,SUBSCRIPTIONS|POST,SUBSCRIPTION|GET,SUBSCRIPTION|PATCH,CLIENT|REPORTS,CLIENT|CONTROLS"

	if len(scope) > 0 {
		defaultScope = scope[0]
	}

	url += defaultScope

	return c.do("GET", url, "", nil)
}

/********** SUBSCRIPTION **********/

// GetSubscriptions returns all of the nodes associated with a user
func (c *Client) GetSubscriptions(queryParams ...string) (map[string]interface{}, error) {
	return c.do("GET", path["subscriptions"], "", queryParams)
}

// GetSubscription returns a single subscription
func (c *Client) GetSubscription(subscriptionID string) (map[string]interface{}, error) {
	url := buildURL(path["subscriptions"], subscriptionID)

	return c.do("GET", url, "", nil)
}

// CreateSubscription creates a subscription and returns the subscription data
func (c *Client) CreateSubscription(data string, idempotencyKey ...string) (map[string]interface{}, error) {
	return c.do("POST", path["subscriptions"], data, idempotencyKey)
}

// UpdateSubscription updates an existing subscription
func (c *Client) UpdateSubscription(subscriptionID string, data string) (map[string]interface{}, error) {
	url := buildURL(path["subscriptions"], subscriptionID)

	return c.do("PATCH", url, data, nil)
}

/********** TRANSACTION **********/

// GetTransactions returns all client transactions
func (c *Client) GetTransactions(queryParams ...string) (map[string]interface{}, error) {
	return c.do("GET", path["transactions"], "", queryParams)
}

/********** USER **********/

// GetUsers returns a list of users
func (c *Client) GetUsers(queryParams ...string) (map[string]interface{}, error) {
	return c.do("GET", path["users"], "", queryParams)
}

// GetUser returns a single user
func (c *Client) GetUser(UserID string, fullDehydrate bool) (*User, error) {
	var user User

	url := buildURL(path["users"], UserID)

	if fullDehydrate != true {
		url += "?full_dehydrate=yes"
	}

	res, err := c.do("GET", url, "", nil)

	mapstructure.Decode(res, &user)

	user.request = user.request.updateRequest(c.ClientID, c.ClientSecret, c.Fingerprint, c.IP)
	user.FullDehydrate = fullDehydrate
	user.Response = res

	return &user, err
}

// CreateUser creates a single user and returns the new user data
func (c *Client) CreateUser(data string, idempotencyKey ...string) (*User, error) {
	var user User

	res, err := c.do("POST", path["users"], data, nil)

	mapstructure.Decode(res, &user)

	user.request = user.request.updateRequest(c.ClientID, c.ClientSecret, c.Fingerprint, c.IP)
	user.Response = res

	return &user, err
}
