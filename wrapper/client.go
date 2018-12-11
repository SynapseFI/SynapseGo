package wrapper

import (
	"strings"

	"github.com/mitchellh/mapstructure"
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

	return read(body), err
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
func (c *Client) GetNodes(queryParams ...string) (map[string]interface{}, error) {
	return c.do("GET", nodesURL, "", queryParams)
}

/********** OTHER **********/

// GetInstitutions returns all of the nodes associated with a user
func (c *Client) GetInstitutions() (map[string]interface{}, error) {
	return c.do("GET", institutionsURL, "", nil)
}

// GetPublicKey returns a public key as a token representing client credentials
func (c *Client) GetPublicKey(scope ...string) (map[string]interface{}, error) {
	url := clientURL + "?issue_public_key=YES&amp;scope="

	for i := 0; i < len(scope); i++ {
		url += scope[i] + ","
	}

	url = strings.TrimSuffix(url, ",")

	return c.do("GET", url, "", nil)
}

/********** SUBSCRIPTION **********/

// GetSubscriptions returns all of the nodes associated with a user
func (c *Client) GetSubscriptions(queryParams ...string) (map[string]interface{}, error) {
	return c.do("GET", subscriptionsURL, "", queryParams)
}

// GetSubscription returns a single subscription
func (c *Client) GetSubscription(subscriptionID string, queryParams ...string) (map[string]interface{}, error) {
	url := buildURL(subscriptionsURL, subscriptionID)

	return c.do("GET", url, "", queryParams)
}

// CreateSubscription creates a subscription and returns the subscription data
func (c *Client) CreateSubscription(data string, queryParams ...string) (map[string]interface{}, error) {
	return c.do("POST", subscriptionsURL, data, queryParams)
}

// UpdateSubscription updates an existing subscription
func (c *Client) UpdateSubscription(subscriptionID string, data string, queryParams ...string) (map[string]interface{}, error) {
	url := buildURL(subscriptionsURL, subscriptionID)

	return c.do("PATCH", url, data, queryParams)
}

/********** TRANSACTION **********/

// GetTransactions returns all client transactions
func (c *Client) GetTransactions(queryParams ...string) (map[string]interface{}, error) {
	return c.do("GET", transactionsURL, "", queryParams)
}

/********** USER **********/

// GetUsers returns a list of users
func (c *Client) GetUsers(queryParams ...string) (map[string]interface{}, error) {
	return c.do("GET", usersURL, "", queryParams)
}

// GetUser returns a single user
func (c *Client) GetUser(UserID string, fullDehydrate bool, queryParams ...string) (*User, error) {
	var user User

	url := buildURL(usersURL, UserID)

	if fullDehydrate != true {
		url += "?full_dehydrate=yes"
	}

	res, err := c.do("GET", url, "", queryParams)

	mapstructure.Decode(res, &user)

	user.request = user.request.updateRequest(c.ClientID, c.ClientSecret, c.Fingerprint, c.IP)
	user.FullDehydrate = fullDehydrate
	user.Response = res

	return &user, err
}

// CreateUser creates a single user and returns the new user data
func (c *Client) CreateUser(data string, queryParams ...string) (*User, error) {
	var user User

	res, err := c.do("POST", usersURL, data, queryParams)

	mapstructure.Decode(res, &user)

	user.request = user.request.updateRequest(c.ClientID, c.ClientSecret, c.Fingerprint, c.IP)
	user.Response = res

	return &user, err
}
