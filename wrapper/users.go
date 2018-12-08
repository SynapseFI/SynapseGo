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
		request       Request
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

func (u *User) do(method, url, data string, queryParams []string, result interface{}) ([]byte, error) {
	var body []byte
	var err error

	u.request = u.request.updateRequest(request.clientID, request.clientSecret, request.fingerprint, request.ipAddress, u.AuthKey)

	switch method {
	case "GET":
		body, err = u.request.Get(url, queryParams, result)

	case "POST":
		body, err = u.request.Post(url, data, queryParams, result)

	case "PATCH":
		body, err = u.request.Patch(url, data, queryParams, result)

	case "DELETE":
		body, err = u.request.Delete(url, result)
	}

	switch err.(type) {
	case *IncorrectUserCredentials:
		_, err = u.Authenticate(`{ "refresh_token": "` + u.RefreshToken + `" }`)

		if err != nil {
			return nil, err
		}

		request.authKey = u.AuthKey

		return u.do(method, url, data, queryParams, result)

	case *IncorrectValues:
		_, err := u.GetRefreshToken()

		if err != nil {
			return nil, err
		}

		_, err = u.Authenticate(`{ "refresh_token": "` + u.RefreshToken + `" }`)

		if err != nil {
			return nil, err
		}

		request.authKey = u.AuthKey

		return u.do(method, url, data, queryParams, result)
	}

	return body, err
}

/********** AUTHENTICATION **********/

// Authenticate returns an oauth key and sets it to the user object
func (u *User) Authenticate(data string) (map[string]interface{}, error) {
	var response map[string]interface{}

	url := buildURL(authURL, u.UserID)

	_, err := u.do("POST", url, data, nil, &response)

	if err != nil {
		return nil, err
	}

	u.AuthKey = response["oauth_key"].(string)
	request.authKey = response["oauth_key"].(string)

	return response, err
}

// GetRefreshToken performs a GET request and returns a new refresh token
func (u *User) GetRefreshToken() (*Refresh, error) {
	var refresh Refresh

	url := buildURL(usersURL, u.UserID)

	_, err := u.do("GET", url, "", nil, &refresh)

	if err != nil {
		return nil, err
	}

	u.RefreshToken = refresh.Token

	return &refresh, err
}

/********** NODE **********/

// GetNodes returns all of the nodes associated with a user
func (u *User) GetNodes(queryParams ...string) (*Nodes, error) {
	var nodes Nodes

	url := buildURL(usersURL, u.UserID, path["nodes"])

	_, err := u.do("GET", url, "", nil, &nodes)

	return &nodes, err
}

// GetNode returns a single node object
func (u *User) GetNode(nodeID string, queryParams ...string) (*Node, error) {
	var node Node

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID)

	_, err := u.do("GET", url, "", nil, &node)

	return &node, err
}

// CreateNode creates a node depending on the type of node specified
func (u *User) CreateNode(data string) (*Nodes, error) {
	var nodes Nodes

	url := buildURL(usersURL, u.UserID, path["nodes"])

	_, err := u.do("POST", url, data, nil, &nodes)

	return &nodes, err
}

// UpdateNode updates a node
func (u *User) UpdateNode(nodeID, data string) (*Node, error) {
	var node Node

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID)

	_, err := u.do("PATCH", url, data, nil, &node)

	return &node, err
}

// DeleteNode deletes a node
func (u *User) DeleteNode(nodeID string) (Response, error) {
	var response Response

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID)

	_, err := u.do("DELETE", url, "", nil, &response)

	return response, err
}

/********** NODE (OTHER) **********/

// AnswerMFA submits an answer to a MFA question from bank login attempt
func (u *User) AnswerMFA(data string) (*Nodes, error) {
	var nodes Nodes

	url := buildURL(usersURL, u.UserID, path["nodes"])

	_, err := u.do("POST", url, data, nil, &nodes)

	return &nodes, err
}

// GetApplePayToken generates tokenized info for Apple Wallet
func (u *User) GetApplePayToken(nodeID, data string) (map[string]interface{}, error) {
	var response map[string]interface{}

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, "applepay")

	_, err := u.do("PATCH", url, data, nil, &response)

	return response, err
}

// ReinitiateMicroDeposit reinitiates micro-deposits for an ACH-US node with AC/RT
func (u *User) ReinitiateMicroDeposit(nodeID string) (*Node, error) {
	var node Node

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID) + "?resend_micro=YES"

	_, err := u.do("PATCH", url, "", nil, &node)

	return &node, err
}

// ResetDebitCard resets the debit card number, card cvv, and expiration date
func (u *User) ResetDebitCard(nodeID string) (map[string]interface{}, error) {
	var response map[string]interface{}

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID) + "?reset=YES"

	_, err := u.do("PATCH", url, "", nil, &response)

	return response, err
}

// ShipDebitCard ships a physical debit card out to the user
func (u *User) ShipDebitCard(nodeID, data string) (map[string]interface{}, error) {
	var response map[string]interface{}

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID) + "?ship=YES"

	_, err := u.do("PATCH", url, data, nil, &response)

	return response, err
}

// TriggerDummyTransactions triggers external dummy transactions on deposit or card accounts
func (u *User) TriggerDummyTransactions(nodeID string, credit bool) (map[string]interface{}, error) {
	var response map[string]interface{}

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID) + "/dummy-tran"

	if credit == true {
		url += "?is_credit=YES"
	}

	_, err := u.do("GET", url, "", nil, &response)

	return response, err
}

// VerifyMicroDeposit verifies micro-deposit amounts for a node
func (u *User) VerifyMicroDeposit(nodeID, data string) (*Node, error) {
	var node Node

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID)

	body, err := u.do("PATCH", url, data, nil, &node)

	node.Response = read(body)

	return &node, err
}

/********** STATEMENT **********/

// GetNodeStatements gets all of the node statements
func (u *User) GetNodeStatements(nodeID string, queryParams ...string) (*Statements, error) {
	var statements Statements

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, path["statements"])

	_, err := u.do("GET", url, "", nil, &statements)

	return &statements, err
}

// GetStatements gets all of the user statements
func (u *User) GetStatements(queryParams ...string) (*Statements, error) {
	var statements Statements

	url := buildURL(usersURL, u.UserID, path["statements"])

	_, err := u.do("GET", url, "", nil, &statements)

	return &statements, err
}

/********** SUBNET **********/

// GetSubnets gets a single subnet object
func (u *User) GetSubnets(nodeID string) (*Subnets, error) {
	var subnets Subnets

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, path["subnets"])

	_, err := u.do("GET", url, "", nil, &subnets)

	return &subnets, err
}

// GetSubnet gets a single subnet object
func (u *User) GetSubnet(nodeID, subnetID string) (*Subnet, error) {
	var subnet Subnet

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, path["subnets"], subnetID)

	body, err := u.do("GET", url, "", nil, &subnet)

	subnet.Response = read(body)

	return &subnet, err
}

// CreateSubnet creates a subnet object
func (u *User) CreateSubnet(nodeID, data string) (*Subnet, error) {
	var subnet Subnet

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, path["subnets"])

	body, err := u.do("PATCH", url, data, nil, &subnet)

	subnet.Response = read(body)

	return &subnet, err
}

/********** TRANSACTION **********/

// GetTransactions returns transactions associated with a node
func (u *User) GetTransactions(nodeID, transactionID string) (*Transactions, error) {
	var transactions Transactions

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, path["trans"])

	_, err := u.do("GET", url, "", nil, &transactions)

	return &transactions, err
}

// GetTransaction returns a specific transaction associated with a node
func (u *User) GetTransaction(nodeID, transactionID string) (*Transaction, error) {
	var transaction Transaction

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, path["trans"], transactionID)

	_, err := u.do("GET", url, "", nil, &transaction)

	return &transaction, err
}

// CreateTransaction creates a transaction for the specified node
func (u *User) CreateTransaction(nodeID, transactionID, data string) (*Transaction, error) {
	var transaction Transaction

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, path["trans"], transactionID)

	body, err := u.do("POST", url, data, nil, &transaction)

	transaction.Response = read(body)

	return &transaction, err
}

// DeleteTransaction deletes/cancels a transaction
func (u *User) DeleteTransaction(nodeID, transactionID string) (*Transaction, error) {
	var transaction Transaction

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, path["transactions"], transactionID)

	_, err := u.do("DELETE", url, "", nil, &transaction)

	return &transaction, err
}

// CommentOnTransactionStatus adds comment to the transaction status
func (u *User) CommentOnTransactionStatus(nodeID, transactionID, data string) (*Transaction, error) {
	var transaction Transaction

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, path["transactions"], transactionID)

	_, err := u.do("POST", url, data, nil, &transaction)

	return &transaction, err
}

// DisputeTransaction disputes a transaction for a user
func (u *User) DisputeTransaction(nodeID, transactionID string) (map[string]interface{}, error) {
	var response map[string]interface{}

	url := buildURL(usersURL, u.UserID, path["nodes"], nodeID, path["transactions"], transactionID, "dispute")

	data := string(`{
		"dispute_reason":"CHARGE_BACK"
	}`)

	_, err := u.do("PATCH", url, data, nil, &response)

	return response, err
}

/********** USER **********/

// Update updates a single user and returns the updated user information
func (u *User) Update(data string, queryParams ...string) (*User, error) {
	var user User

	url := buildURL(usersURL, u.UserID)

	body, err := u.do("PATCH", url, data, nil, &user)

	user.Response = read(body)

	return &user, err
}

// CreateUBO creates and uploads an Ultimate Beneficial Ownership (UBO) and REG GG form as a physical document under the Business’s base document
func (u *User) CreateUBO(data string) (map[string]interface{}, error) {
	var response map[string]interface{}

	url := buildURL(usersURL, u.UserID, "ubo")

	body, err := u.do("PATCH", url, data, nil, &response)

	user.Response = read(body)

	return response, err
}
