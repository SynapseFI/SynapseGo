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

// Authenticate returns an oauth key and sets it to the user object
func (u *User) Authenticate(data string) *Response {
	var response Response

	url := buildURL(authURL, u.UserID)

	_, err := request.Post(url, data, nil, &response)

	if err != nil {
		panic(err)
	}

	return &response
}

/********** NODE **********/

// AnswerMFA submits an answer to a MFA question from bank login attempt
func (u *User) AnswerMFA(data string) *Nodes {
	var nodes Nodes

	url := buildURL(usersURL, u.UserID, path["nodes"])

	_, err := request.Post(url, data, nil, &nodes)

	if err != nil {
		panic(err)
	}

	return &nodes
}

// CreateNode creates a node depending on the type of node specified
func (u *User) CreateNode(data string) *Nodes {
	var nodes Nodes

	url := buildURL(usersURL, u.UserID, path["nodes"])

	_, err := request.Post(url, data, nil, &nodes)

	if err != nil {
		panic(err)
	}

	return &nodes
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

// GetApplePayToken generates tokenized info for Apple Wallet
func (u *User) GetApplePayToken(nodeID, data string) *Response {
	var response Response

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, "applepay")

	_, err := request.Patch(url, data, nil, &response)

	if err != nil {
		panic(err)
	}

	return &response
}

// GetNode returns a single node object
func (u *User) GetNode(nodeID string, queryParams ...string) *Node {
	var node Node

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID)

	_, err := request.Get(url, nil, &node)

	if err != nil {
		panic(err)
	}

	return &node
}

// GetNodes returns all of the nodes associated with a user
func (u *User) GetNodes(queryParams ...string) *Nodes {
	var nodes Nodes

	url := buildURL(usersURL, u.UserID, path["nodes"])

	_, err := request.Get(url, nil, &nodes)

	if err != nil {
		panic(err)
	}

	return &nodes
}

// ReintiateMicroDeposit reinitiates micro-deposits for an ACH-US node with AC/RT
func (u *User) ReintiateMicroDeposit(nodeID string) *Node {
	var node Node

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID) + "?resend_micro=YES"

	_, err := request.Patch(url, "", nil, &node)

	if err != nil {
		panic(err)
	}

	return &node
}

// ResetDebitCard resets the debit card number, card cvv, and expiration date
func (u *User) ResetDebitCard(nodeID string) *Response {
	var response Response

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID) + "?reset=YES"

	_, err := request.Patch(url, "", nil, &response)

	if err != nil {
		panic(err)
	}

	return &response
}

// ShipDebitCard ships a physical debit card out to the user
func (u *User) ShipDebitCard(nodeID, data string) *Response {
	var response Response

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID) + "?ship=YES"

	_, err := request.Patch(url, data, nil, &response)

	if err != nil {
		panic(err)
	}

	return &response
}

// TriggerDummyTransactions triggers external dummy transactions on deposit or card accounts
func (u *User) TriggerDummyTransactions(nodeID string, credit bool) *Response {
	var response Response

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID) + "/dummy-tran"

	if credit == true {
		url += "?is_credit=YES"
	}

	_, err := request.Get(url, nil, &response)

	if err != nil {
		panic(err)
	}

	return &response
}

// UpdateNode updates a node
func (u *User) UpdateNode(nodeID, data string) *Node {
	var node Node

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID)

	_, err := request.Patch(url, data, nil, &node)

	if err != nil {
		panic(err)
	}

	return &node
}

// VerifyMicroDeposit verifies micro-deposit amounts for a node
func (u *User) VerifyMicroDeposit(nodeID, data string) *Node {
	var node Node

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID)

	body, err := request.Patch(url, data, nil, &node)

	if err != nil {
		panic(err)
	}

	node.Response = read(body)

	return &node
}

/********** STATEMENT **********/

// GetNodeStatements gets all of the node statements
func (u *User) GetNodeStatements(nodeID string, queryParams ...string) *Statements {
	var statements Statements

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, path["statements"])

	_, err := request.Get(url, nil, &statements)

	if err != nil {
		panic(err)
	}

	return &statements
}

// GetStatements gets all of the user statements
func (u *User) GetStatements(queryParams ...string) *Statements {
	var statements Statements

	url := buildURL(usersURL, u.UserID, path["statements"])

	_, err := request.Get(url, nil, &statements)

	if err != nil {
		panic(err)
	}

	return &statements
}

/********** SUBNET **********/

// CreateSubnet creates a subnet object
func (u *User) CreateSubnet(nodeID, data string) *Subnet {
	var subnet Subnet

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, path["subnets"])

	body, err := request.Patch(url, data, nil, &subnet)

	if err != nil {
		panic(err)
	}

	subnet.Response = read(body)

	return &subnet
}

// GetSubnet gets a single subnet object
func (u *User) GetSubnet(nodeID, subnetID string) *Subnet {
	var subnet Subnet

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, path["subnets"], subnetID)

	body, err := request.Get(url, nil, &subnet)

	if err != nil {
		panic(err)
	}

	subnet.Response = read(body)

	return &subnet
}

// GetSubnets gets a single subnet object
func (u *User) GetSubnets(nodeID string) *Subnets {
	var subnets Subnets

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, path["subnets"])

	_, err := request.Get(url, nil, &subnets)

	if err != nil {
		panic(err)
	}

	return &subnets
}

/********** TRANSACTION **********/

// CancelTransaction cancels a transaction
func (u *User) CancelTransaction(nodeID, transactionID, data string) *Transaction {
	var transaction Transaction

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, path["transactions"], transactionID)

	_, err := request.Delete(url, &transaction)

	if err != nil {
		panic(err)
	}

	return &transaction
}

// CommentOnTransactionStatus adds comment to the transaction status
func (u *User) CommentOnTransactionStatus(nodeID, transactionID, data string) *Transaction {
	var transaction Transaction

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, path["transactions"], transactionID)

	_, err := request.Post(url, data, nil, &transaction)

	if err != nil {
		panic(err)
	}

	return &transaction
}

// CreateTransaction creates a transaction for the specified node
func (u *User) CreateTransaction(nodeID, transactionID, data string) *Transaction {
	var transaction Transaction

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, path["trans"], transactionID)

	body, err := request.Post(url, data, nil, &transaction)

	if err != nil {
		panic(err)
	}

	transaction.Response = read(body)

	return &transaction
}

// DisputeTransaction disputes a transaction for a user
func (u *User) DisputeTransaction(nodeID, transactionID string) *Node {
	var node Node

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, path["transactions"], transactionID, "dispute")

	data := string(`{
		"dispute_reason":"CHARGE_BACK"
	}`)

	_, err := request.Patch(url, data, nil, &node)

	if err != nil {
		panic(err)
	}

	return &node
}

// GetTransaction returns a specific transaction associated with a node
func (u *User) GetTransaction(nodeID, transactionID string) *Transaction {
	var transaction Transaction

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, path["trans"], transactionID)

	_, err := request.Get(url, nil, &transaction)

	if err != nil {
		panic(err)
	}

	return &transaction
}

// GetTransactions returns transactions associated with a node
func (u *User) GetTransactions(nodeID, transactionID string) *Transactions {
	var transactions Transactions

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, path["trans"])

	_, err := request.Get(url, nil, &transactions)

	if err != nil {
		panic(err)
	}

	return &transactions
}

/********** USER **********/

// CreateUBO creates and uploads an Ultimate Beneficial Ownership (UBO) and REG GG form as a physical document under the Business’s base document
func (u *User) CreateUBO(data string) *User {
	var user User

	url := buildURL(usersURL, u.UserID, "ubo")

	body, err := request.Patch(url, data, nil, &user)

	if err != nil {
		panic(err)
	}

	user.Response = read(body)

	return &user
}

// Update updates a single user and returns the updated user information
func (u *User) Update(data string, queryParams ...string) *User {
	var user User

	url := buildURL(usersURL, u.UserID)

	body, err := request.Patch(url, data, nil, &user)

	if err != nil {
		panic(err)
	}

	user.Response = read(body)

	return &user
}
