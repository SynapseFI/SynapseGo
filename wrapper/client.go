package wrapper

import "strings"

/********** GLOBAL VARIABLES **********/
var developerMode = false

const institutionsURL = _url + "/institutions"
const publicKeyURL = _url + "/client?issue_public_key=YES&amp;scope="

/********** TYPES **********/

type (
	// Client represents the credentials used by the developer to instantiate a client
	Client struct {
		Fingerprint string
		Gateway     string
		IP          string
	}

	// Institutions represents a list of Synapse institutions
	Institutions struct {
		Banks interface{} `json:"banks"`
	}

	// PublicKey represents the structure of a public key object
	PublicKey struct {
		Response interface{} `json:"public_key_obj"`
	}
)

/********** METHODS **********/

/********** CLIENT **********/

// NewClient creates a client object
func NewClient(clientID, clientSecret, ipAddress, fingerprint string, devMode ...bool) *Client {
	if len(devMode) > 0 && devMode[0] == true {
		developerMode = true
	}

	request = newRequest(clientID, clientSecret, fingerprint, ipAddress)

	return &Client{
		Fingerprint: "|" + fingerprint,
		Gateway:     clientID + "|" + clientSecret,
		IP:          ipAddress,
	}

}

/********** NODE **********/

// GetAllNodes returns all of the nodes
func (c *Client) GetAllNodes(queryParams ...string) *Nodes {
	var nodes Nodes

	_, err := request.Get(nodesURL, queryParams[0], &nodes)

	if err != nil {
		panic(err)
	}

	return &nodes
}

/********** OTHER **********/

// GetInstitutions returns all of the nodes associated with a user
func (c *Client) GetInstitutions() *Institutions {
	var institutions Institutions

	request.Get(institutionsURL, "", &institutions)

	return &institutions
}

// GetPublicKey returns a public key as a token representing client credentials
func (c *Client) GetPublicKey(scope ...string) *PublicKey {
	var publicKey PublicKey
	var urlParams = publicKeyURL

	for i := 0; i < len(scope); i++ {
		urlParams += scope[i] + ","
	}

	urlParams = strings.TrimSuffix(urlParams, ",")

	_, err := request.Get(urlParams, "", &publicKey)

	if err != nil {
		panic(err)
	}

	return &publicKey
}

/********** SUBSCRIPTION **********/

// GetSubscriptions returns all of the nodes associated with a user
func (c *Client) GetSubscriptions(queryParams ...string) *Subscriptions {
	var subscriptions Subscriptions

	_, err := request.Get(subscriptionsURL, "", &subscriptions)

	if err != nil {
		panic(err)
	}

	return &subscriptions
}

// GetSubscription returns a single subscription
func (c *Client) GetSubscription(subID string, queryParams ...string) *Subscription {
	var subscription Subscription

	url := subscriptionsURL + "/" + subID

	body, err := request.Get(url, "", &subscription)

	if err != nil {
		panic(err)
	}

	subscription.Response = read(body)

	return &subscription
}

// CreateSubscription creates a subscription and returns the subscription data
func (c *Client) CreateSubscription(data string, queryParams ...string) *Subscription {
	var subscription Subscription

	body, err := request.Post(subscriptionsURL, data, "", &subscription)

	if err != nil {
		panic(err)
	}

	subscription.Response = read(body)

	return &subscription
}

// UpdateSubscription updates an existing subscription
func (c *Client) UpdateSubscription(subID string, data string, queryParams ...string) *Subscription {
	var subscription Subscription

	body, err := request.Patch(subscriptionsURL, data, "", &subscription)

	if err != nil {
		panic(err)
	}

	subscription.Response = read(body)

	return &subscription
}

/********** TRANSACTION **********/

// GetClientTransactions returns all client transactions
func (c *Client) GetClientTransactions(queryParams ...string) *Transactions {
	var transactions Transactions

	_, err := request.Get(transactionsURL, queryParams[0], &transactions)

	if err != nil {
		panic(err)
	}

	return &transactions
}

/********** USER **********/

// GetUsers returns a list of users
func (c *Client) GetUsers(queryParams ...string) *Users {
	var users Users

	_, err := request.Get(usersURL, "", &users)

	if err != nil {
		panic(err)
	}

	return &users
}

// GetUser returns a single user
func (c *Client) GetUser(UserID string, fullDehydrate bool, queryParams ...string) *User {
	var user User

	url := usersURL + "/" + UserID

	if fullDehydrate != true {
		url += "?full_dehydrate=yes"
	}

	body, err := request.Get(url, "", &user)

	if err != nil {
		panic(err)
	}

	user.FullDehydrate = fullDehydrate
	user.Response = read(body)

	return &user
}

// CreateUser creates a single user and returns the new user data
func (c *Client) CreateUser(data string, queryParams ...string) *User {
	var user User

	body, err := request.Post(usersURL, data, "", &user)

	if err != nil {
		panic(err)
	}

	user.Response = read(body)

	return &user
}
