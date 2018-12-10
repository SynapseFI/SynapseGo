package wrapper

import (
	"strings"
)

/********** GLOBAL VARIABLES **********/
var developerMode = false

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

	// Institutions represents a list of Synapse institutions
	Institutions struct {
		Banks interface{} `json:"banks"`
	}

	// PublicKey represents the structure of a public key object
	PublicKey struct {
		Response map[string]interface{} `json:"public_key_obj"`
	}
)

/********** METHODS **********/

func (c *Client) do(method, url, data string, queryParams []string, result interface{}) ([]byte, error) {
	var body []byte
	var err error

	c.request = c.request.updateRequest(c.ClientID, c.ClientSecret, c.Fingerprint, c.IP)

	switch method {
	case "GET":
		body, err = c.request.Get(url, queryParams, result)

	case "POST":
		body, err = c.request.Post(url, data, queryParams, result)

	case "PATCH":
		body, err = c.request.Patch(url, data, queryParams, result)

	case "DELETE":
		body, err = c.request.Delete(url, result)
	}

	return body, err
}

/********** CLIENT **********/

// New creates a client object
func New(clientID, clientSecret, ipAddress, fingerprint string, devMode ...bool) *Client {
	if len(devMode) > 0 && devMode[0] == true {
		developerMode = true
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
func (c *Client) GetNodes(queryParams ...string) (*Nodes, error) {
	var nodes Nodes

	_, err := c.do("GET", nodesURL, "", queryParams, &nodes)

	return &nodes, err
}

/********** OTHER **********/

// GetInstitutions returns all of the nodes associated with a user
func (c *Client) GetInstitutions() (*Institutions, error) {
	var institutions Institutions

	_, err := c.do("GET", institutionsURL, "", nil, &institutions)

	return &institutions, err
}

// GetPublicKey returns a public key as a token representing client credentials
func (c *Client) GetPublicKey(scope ...string) (*PublicKey, error) {
	var publicKey PublicKey

	url := clientURL + "?issue_public_key=YES&amp;scope="

	for i := 0; i < len(scope); i++ {
		url += scope[i] + ","
	}

	url = strings.TrimSuffix(url, ",")

	_, err := c.do("GET", url, "", nil, &publicKey)

	return &publicKey, err
}

/********** SUBSCRIPTION **********/

// GetSubscriptions returns all of the nodes associated with a user
func (c *Client) GetSubscriptions(queryParams ...string) (*Subscriptions, error) {
	var subscriptions Subscriptions

	_, err := c.do("GET", subscriptionsURL, "", queryParams, &subscriptions)

	return &subscriptions, err
}

// GetSubscription returns a single subscription
func (c *Client) GetSubscription(subscriptionID string, queryParams ...string) (*Subscription, error) {
	var subscription Subscription

	url := buildURL(subscriptionsURL, subscriptionID)

	body, err := c.do("GET", url, "", queryParams, &subscription)

	subscription.Response = read(body)

	return &subscription, err
}

// CreateSubscription creates a subscription and returns the subscription data
func (c *Client) CreateSubscription(data string, queryParams ...string) (*Subscription, error) {
	var subscription Subscription

	body, err := c.do("POST", subscriptionsURL, data, queryParams, &subscription)

	subscription.Response = read(body)

	return &subscription, err
}

// UpdateSubscription updates an existing subscription
func (c *Client) UpdateSubscription(subscriptionID string, data string, queryParams ...string) (*Subscription, error) {
	var subscription Subscription

	url := buildURL(subscriptionsURL, subscriptionID)

	body, err := c.do("PATCH", url, data, queryParams, &subscription)

	subscription.Response = read(body)

	return &subscription, err
}

/********** TRANSACTION **********/

// GetTransactions returns all client transactions
func (c *Client) GetTransactions(queryParams ...string) (*Transactions, error) {
	var transactions Transactions

	_, err := c.do("GET", transactionsURL, "", queryParams, &transactions)

	return &transactions, err
}

/********** USER **********/

// GetUsers returns a list of users
func (c *Client) GetUsers(queryParams ...string) (*Users, error) {
	var users Users

	_, err := c.do("GET", usersURL, "", queryParams, &users)

	return &users, err
}

// GetUser returns a single user
func (c *Client) GetUser(UserID string, fullDehydrate bool, queryParams ...string) (*User, error) {
	var user User

	url := buildURL(usersURL, UserID)

	if fullDehydrate != true {
		url += "?full_dehydrate=yes"
	}

	body, err := c.do("GET", url, "", queryParams, &user)

	user.request = user.request.updateRequest(c.ClientID, c.ClientSecret, c.Fingerprint, c.IP)
	user.FullDehydrate = fullDehydrate
	user.Response = read(body)

	return &user, err
}

// CreateUser creates a single user and returns the new user data
func (c *Client) CreateUser(data string, queryParams ...string) (*User, error) {
	var user User

	body, err := c.do("POST", usersURL, data, queryParams, &user)

	user.request = user.request.updateRequest(c.ClientID, c.ClientSecret, c.Fingerprint, c.IP)
	user.Response = read(body)

	return &user, err
}
