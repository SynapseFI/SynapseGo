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

func (u *User) request(method, url, data string, queryParams []string, result interface{}) ([]byte, error) {
	var body []byte
	var err error

	switch method {
	case "GET":
		body, err = request.Get(url, queryParams, result)

	case "POST":
		body, err = request.Post(url, data, queryParams, result)

	case "PATCH":
		body, err = request.Patch(url, data, queryParams, result)

	case "DELETE":
		body, err = request.Delete(url, result)
	}

	switch err.(type) {
	case *IncorrectUserCredentials:
		var b map[string]interface{}

		rt := `{ "refresh_token": "` + u.RefreshToken + `" }`
		b, err = u.Authenticate(rt)
		u.AuthKey = b["oauth_key"].(string)

		if err != nil {
			return nil, err
		}

		return u.request(method, url, data, queryParams, result)

	case *IncorrectValues:
		var b map[string]interface{}
		var user *User

		u.request("GET", usersURL, "", nil, &user)
		u.RefreshToken = user.RefreshToken
		rt := `{ "refresh_token": "` + u.RefreshToken + `" }`
		b, err = u.Authenticate(rt)
		u.AuthKey = b["oauth_key"].(string)

		if err != nil {
			return nil, err
		}

		return u.request(method, url, data, queryParams, result)
	}

	return body, err
}

/********** AUTHENTICATION **********/

// Authenticate returns an oauth key and sets it to the user object
func (u *User) Authenticate(data string) (map[string]interface{}, error) {
	var response map[string]interface{}

	url := buildURL(authURL, u.UserID)

	_, err := request.Post(url, data, nil, &response)

	return response, err
}

/********** NODE **********/

// AnswerMFA submits an answer to a MFA question from bank login attempt
func (u *User) AnswerMFA(data string) (*Nodes, error) {
	var nodes Nodes

	url := buildURL(usersURL, u.UserID, path["nodes"])

	_, err := request.Post(url, data, nil, &nodes)

	return &nodes, err
}

// CreateNode creates a node depending on the type of node specified
func (u *User) CreateNode(data string) (*Nodes, error) {
	var nodes Nodes

	url := buildURL(usersURL, u.UserID, path["nodes"])

	_, err := request.Post(url, data, nil, &nodes)

	return &nodes, err
}

// DeleteNode deletes a node
func (u *User) DeleteNode(nodeID string) (Response, error) {
	var response Response

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID)

	_, err := request.Delete(url, &response)

	return response, err
}

// GetApplePayToken generates tokenized info for Apple Wallet
func (u *User) GetApplePayToken(nodeID, data string) (Response, error) {
	var response Response

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, "applepay")

	_, err := request.Patch(url, data, nil, &response)

	return response, err
}

// GetNode returns a single node object
func (u *User) GetNode(nodeID string, queryParams ...string) (*Node, error) {
	var node Node

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID)

	_, err := request.Get(url, nil, &node)

	return &node, err
}

// GetNodes returns all of the nodes associated with a user
func (u *User) GetNodes(queryParams ...string) (*Nodes, error) {
	var nodes Nodes

	url := buildURL(usersURL, u.UserID, path["nodes"])

	_, err := request.Get(url, nil, &nodes)

	return &nodes, err
}

// ReintiateMicroDeposit reinitiates micro-deposits for an ACH-US node with AC/RT
func (u *User) ReintiateMicroDeposit(nodeID string) (*Node, error) {
	var node Node

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID) + "?resend_micro=YES"

	_, err := request.Patch(url, "", nil, &node)

	return &node, err
}

// ResetDebitCard resets the debit card number, card cvv, and expiration date
func (u *User) ResetDebitCard(nodeID string) (, error) {
	var response Response

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID) + "?reset=YES"

	_, err := request.Patch(url, "", nil, &response)

	return response, err
}

// ShipDebitCard ships a physical debit card out to the user
func (u *User) ShipDebitCard(nodeID, data string) (Response, error) {
	var response Response

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID) + "?ship=YES"

	_, err := request.Patch(url, data, nil, &response)

	return response, err
}

// TriggerDummyTransactions triggers external dummy transactions on deposit or card accounts
func (u *User) TriggerDummyTransactions(nodeID string, credit bool) (Response, error) {
	var response Response

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID) + "/dummy-tran"

	if credit == true {
		url += "?is_credit=YES"
	}

	_, err := request.Get(url, nil, &response)

	return response, err
}

// UpdateNode updates a node
func (u *User) UpdateNode(nodeID, data string) (*Node, error) {
	var node Node

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID)

	_, err := request.Patch(url, data, nil, &node)

	return &node, err
}

// VerifyMicroDeposit verifies micro-deposit amounts for a node
func (u *User) VerifyMicroDeposit(nodeID, data string) (*Node, error) {
	var node Node

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID)

	body, err := request.Patch(url, data, nil, &node)

	node.Response = read(body)

	return &node, err
}

/********** STATEMENT **********/

// GetNodeStatements gets all of the node statements
func (u *User) GetNodeStatements(nodeID string, queryParams ...string) (*Statements, error) {
	var statements Statements

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, path["statements"])

	_, err := request.Get(url, nil, &statements)

	return &statements, err
}

// GetStatements gets all of the user statements
func (u *User) GetStatements(queryParams ...string) (*Statements, error) {
	var statements Statements

	url := buildURL(usersURL, u.UserID, path["statements"])

	_, err := request.Get(url, nil, &statements)

	return &statements, err
}

/********** SUBNET **********/

// CreateSubnet creates a subnet object
func (u *User) CreateSubnet(nodeID, data string) (*Subnet, error) {
	var subnet Subnet

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, path["subnets"])

	body, err := request.Patch(url, data, nil, &subnet)

	subnet.Response = read(body)

	return &subnet, err
}

// GetSubnet gets a single subnet object
func (u *User) GetSubnet(nodeID, subnetID string) (*Subnet, error) {
	var subnet Subnet

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, path["subnets"], subnetID)

	body, err := request.Get(url, nil, &subnet)

	subnet.Response = read(body)

	return &subnet, err
}

// GetSubnets gets a single subnet object
func (u *User) GetSubnets(nodeID string) (*Subnets, error) {
	var subnets Subnets

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, path["subnets"])

	_, err := request.Get(url, nil, &subnets)

	return &subnets, err
}

/********** TRANSACTION **********/

// CancelTransaction cancels a transaction
func (u *User) CancelTransaction(nodeID, transactionID, data string) (*Transaction, error) {
	var transaction Transaction

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, path["transactions"], transactionID)

	_, err := request.Delete(url, &transaction)

	return &transaction, err
}

// CommentOnTransactionStatus adds comment to the transaction status
func (u *User) CommentOnTransactionStatus(nodeID, transactionID, data string) (*Transaction, error) {
	var transaction Transaction

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, path["transactions"], transactionID)

	_, err := request.Post(url, data, nil, &transaction)

	return &transaction, err
}

// CreateTransaction creates a transaction for the specified node
func (u *User) CreateTransaction(nodeID, transactionID, data string) (*Transaction, error) {
	var transaction Transaction

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, path["trans"], transactionID)

	body, err := request.Post(url, data, nil, &transaction)

	transaction.Response = read(body)

	return &transaction, err
}

// DisputeTransaction disputes a transaction for a user
func (u *User) DisputeTransaction(nodeID, transactionID string) (*Node, error) {
	var node Node

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, path["transactions"], transactionID, "dispute")

	data := string(`{
		"dispute_reason":"CHARGE_BACK"
	}`)

	_, err := request.Patch(url, data, nil, &node)

	return &node, err
}

// GetTransaction returns a specific transaction associated with a node
func (u *User) GetTransaction(nodeID, transactionID string) (*Transaction, error) {
	var transaction Transaction

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, path["trans"], transactionID)

	_, err := request.Get(url, nil, &transaction)

	return &transaction, err
}

// GetTransactions returns transactions associated with a node
func (u *User) GetTransactions(nodeID, transactionID string) (*Transactions, error) {
	var transactions Transactions

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, path["trans"])

	_, err := request.Get(url, nil, &transactions)

	return &transactions, err
}

/********** USER **********/

// CreateUBO creates and uploads an Ultimate Beneficial Ownership (UBO) and REG GG form as a physical document under the Business’s base document
func (u *User) CreateUBO(data string) (*User, error) {
	var user User

	url := buildURL(usersURL, u.UserID, "ubo")

	body, err := request.Patch(url, data, nil, &user)

	user.Response = read(body)

	return &user, err
}

// Update updates a single user and returns the updated user information
func (u *User) Update(data string, queryParams ...string) (*User, error) {
	var user User

	url := buildURL(usersURL, u.UserID)

	body, err := request.Patch(url, data, nil, &user)

	user.Response = read(body)

	return &user, err
}
