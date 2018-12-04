package wrapper

/*********** GLOBAL VARIABLES ***********/

/********** TYPES **********/

type (
	// Auth represents an oauth key
	Auth struct {
		Key string `json:"oauth_key"`
	}

	// MFA represents multi-factor authentication response
	MFA struct {
		AccessToken string `json:"access_token"`
		Message     string `json:"message"`
		Type        string `json:"type"`
	}

	// Refresh represents a refresh token
	Refresh struct {
		Token string `json:"refresh_token"`
	}

	// User represents a single user object
	User struct {
		AuthKey       string `json:"oauth_key"`
		FullDehydrate bool
		UserID        string `json:"_id"`
		RefreshToken  string `json:"refresh_token"`
		Response      interface{}
	}

	// Users represents a collection of user objects
	Users struct {
		Limit      int64  `json:"limit"`
		Page       int64  `json:"page"`
		PageCount  int64  `json:"page_count"`
		UsersCount int64  `json:"users_count"`
		Users      []User `json:"users"`
	}
)

/********** METHODS **********/

/********** AUTHENTICATION **********/

// Auth returns an oauth key and sets it to the user object
func (u *User) Auth(body ...string) *Auth {
	var data string
	if len(body) > 0 {
		data = body[0]
	}

	auth := request.authenticate(u.UserID, u.RefreshToken, data)
	request.authKey = auth.Key

	return auth
}

/********** NODE **********/

// GetNodes returns all of the nodes associated with a user
func (u *User) GetNodes(queryParams ...string) *Nodes {
	var nodes Nodes

	url := buildURL(usersURL, u.UserID, path["nodes"])

	_, err := request.Get(url, "", &nodes, u)

	if err != nil {
		panic(err)
	}

	return &nodes
}

// GetNode returns a single node object
func (u *User) GetNode(nodeID string, queryParams ...string) *Node {
	var node Node

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID)

	_, err := request.Get(url, "", &node)

	if err != nil {
		panic(err)
	}

	return &node
}

// CreateNode creates a node depending on the type of node specified
func (u *User) CreateNode(data string) *Nodes {
	var nodes Nodes

	url := buildURL(usersURL, u.UserID, path["nodes"])

	_, err := request.Post(url, data, "", &nodes, u)

	if err != nil {
		panic(err)
	}

	return &nodes
}

// UpdateNode updates a node
func (u *User) UpdateNode(nodeID, data string) *Node {
	var node Node

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID)

	_, err := request.Patch(url, data, "", &node)

	if err != nil {
		panic(err)

	}

	return &node
}

// DeleteNode deletes a node
func (u *User) DeleteNode(nodeID string) *Response {
	var response Response

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID)

	_, err := request.Delete(url, &response)

	if err != nil {
		panic(err)
	}

	return &response
}

/********** TRANSACTION **********/

// GetTransaction returns a specific transaction associated with a node
func (n *Node) GetTransaction(transactionID string) *Transaction {
	var transaction Transaction

	url := buildURL(usersURL, n.UserID, path["nodes"], n.NodeID, path["trans"], transactionID)

	_, err := request.Get(url, "", &transaction)

	if err != nil {
		panic(err)
	}

	return &transaction
}

// CreateTransaction creates a transaction for the specified node
func (n *Node) CreateTransaction(transactionID, data string) *Transaction {
	var transaction Transaction

	url := buildURL(usersURL, n.UserID, path["nodes"], n.NodeID, path["trans"], transactionID)

	_, err := request.Post(url, data, "", &transaction)

	if err != nil {
		panic(err)
	}

	return &transaction
}

/********** USER **********/

// Update updates a single user and returns the updated user information
func (u *User) Update(data string, queryParams ...string) *User {
	url := buildURL(usersURL, u.UserID)

	body, err := request.Patch(url, data, "", u, u)

	if err != nil {
		panic(err)
	}

	u.Response = read(body)

	return u
}
