package synapse

import "github.com/mitchellh/mapstructure"

/*********** GLOBAL VARIABLES ***********/

/********** TYPES **********/

type (
	// User represents a single user object
	User struct {
		AuthKey       string
		FullDehydrate bool
		UserID        string `mapstructure:"_id"`
		RefreshToken  string `mapstructure:"refresh_token"`
		Response      interface{}
		request       Request
	}

	// Users represents a collection of user objects
	Users struct {
		Limit      int64  `mapstructure:"limit"`
		Page       int64  `mapstructure:"page"`
		PageCount  int64  `mapstructure:"page_count"`
		UsersCount int64  `mapstructure:"users_count"`
		Users      []User `mapstructure:"users"`
	}
)

/********** METHODS **********/

func (u *User) do(method, url, data string, params []string) (map[string]interface{}, error) {
	var response []byte
	var err error

	u.request = u.request.updateRequest(u.request.clientID, u.request.clientSecret, u.request.fingerprint, u.request.ipAddress, u.AuthKey)

	switch method {
	case "GET":
		response, err = u.request.Get(url, params)

	case "POST":
		response, err = u.request.Post(url, data, params)

	case "PATCH":
		response, err = u.request.Patch(url, data, params)

	case "DELETE":
		response, err = u.request.Delete(url)
	}

	switch err.(type) {
	case *IncorrectUserCredentials:
		_, err = u.Authenticate(`{ "refresh_token": "` + u.RefreshToken + `" }`)

		if err != nil {
			return nil, err
		}

		u.request.authKey = u.AuthKey

		return u.do(method, url, data, params)

	case *IncorrectValues:
		_, err := u.GetRefreshToken()

		if err != nil {
			return nil, err
		}

		_, err = u.Authenticate(`{ "refresh_token": "` + u.RefreshToken + `" }`)

		if err != nil {
			return nil, err
		}

		u.request.authKey = u.AuthKey

		return u.do(method, url, data, params)
	}

	return readStream(response), err
}

/********** AUTHENTICATION **********/

// Authenticate returns an oauth key and sets it to the user object
func (u *User) Authenticate(data string) (map[string]interface{}, error) {
	url := buildURL(path["auth"], u.UserID)

	res, err := u.do("POST", url, data, nil)

	if res["oauth_key"] != nil {
		u.AuthKey = res["oauth_key"].(string)
		u.request.authKey = res["oauth_key"].(string)
	}

	return res, err
}

// GetRefreshToken performs a GET request and returns a new refresh token
func (u *User) GetRefreshToken() (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID)

	res, err := u.do("GET", url, "", nil)

	if res["refresh_token"] != nil {
		u.RefreshToken = res["refresh_token"].(string)
	}

	return res, err
}

// Select2FA sends the 2FA device selection to the system
func (u *User) Select2FA(device string) (map[string]interface{}, error) {
	url := buildURL(path["auth"], u.UserID)

	data := `{ "refresh_token": "` + u.RefreshToken + `", "phone_number": "` + device + `" }`

	res, err := u.do("POST", url, data, nil)

	return res, err
}

// SubmitMFA submits the access token and mfa answer
func (u *User) SubmitMFA(data string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"])

	return u.do("POST", url, data, nil)
}

// VerifyPIN sends the requested pin to the system to complete the 2FA process
func (u *User) VerifyPIN(pin string) (map[string]interface{}, error) {
	url := buildURL(path["auth"], u.UserID)

	data := `{ "refresh_token": "` + u.RefreshToken + `", "validation_pin": "` + pin + `" }`

	res, err := u.do("POST", url, data, nil)

	return res, err
}

/********** NODE **********/

// GetNodes returns all of the nodes associated with a user
func (u *User) GetNodes(queryParams ...string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"])

	return u.do("GET", url, "", queryParams)
}

// GetNode returns a single node object
func (u *User) GetNode(nodeID string) (map[string]interface{}, error) {

	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID)

	res, err := u.do("GET", url, "", nil)

	return res, err
}

// CreateNode creates a node depending on the type of node specified
func (u *User) CreateNode(data string, idempotencyKey ...string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"])

	return u.do("POST", url, data, idempotencyKey)
}

// UpdateNode updates a node
func (u *User) UpdateNode(nodeID, data string) (map[string]interface{}, error) {

	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID)

	return u.do("PATCH", url, data, nil)
}

// DeleteNode deletes a node
func (u *User) DeleteNode(nodeID string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID)

	return u.do("DELETE", url, "", nil)
}

/********** NODE (OTHER) **********/

// GetApplePayToken generates tokenized info for Apple Wallet
func (u *User) GetApplePayToken(nodeID, data string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, "applepay")

	return u.do("PATCH", url, data, nil)
}

// ReinitiateMicroDeposit reinitiates micro-deposits for an ACH-US node with AC/RT
func (u *User) ReinitiateMicroDeposit(nodeID string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID) + "?resend_micro=YES"

	return u.do("PATCH", url, "", nil)
}

// ResetDebitCard resets the debit card number, card cvv, and expiration date
func (u *User) ResetDebitCard(nodeID string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID) + "?reset=YES"

	return u.do("PATCH", url, "", nil)
}

// ShipDebitCard ships a physical debit card out to the user
func (u *User) ShipDebitCard(nodeID, data string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID) + "?ship=YES"

	return u.do("PATCH", url, data, nil)
}

// TriggerDummyTransactions triggers external dummy transactions on deposit or card accounts
func (u *User) TriggerDummyTransactions(nodeID string, credit bool) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID) + "/dummy-tran"

	if credit == true {
		url += "?is_credit=YES"
	}

	return u.do("GET", url, "", nil)
}

// VerifyMicroDeposit verifies micro-deposit amounts for a node
func (u *User) VerifyMicroDeposit(nodeID, data string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID)

	return u.do("PATCH", url, data, nil)
}

/********** STATEMENT **********/

// GetNodeStatements gets all of the node statements
func (u *User) GetNodeStatements(nodeID string, queryParams ...string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["statements"])

	return u.do("GET", url, "", queryParams)
}

// GetStatements gets all of the user statements
func (u *User) GetStatements(queryParams ...string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["statements"])

	return u.do("GET", url, "", queryParams)
}

/********** SUBNET **********/

// GetNodeSubnets gets all subnets associated with a node
func (u *User) GetNodeSubnets(nodeID string, queryParams ...string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["subnets"])

	return u.do("GET", url, "", queryParams)
}

// GetSubnet gets a single subnet object
func (u *User) GetSubnet(nodeID, subnetID string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["subnets"], subnetID)

	return u.do("GET", url, "", nil)
}

// CreateSubnet creates a subnet object
func (u *User) CreateSubnet(nodeID, data string, idempotencyKey ...string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["subnets"])

	return u.do("PATCH", url, data, idempotencyKey)
}

/********** TRANSACTION **********/

// GetNodeTransactions returns transactions associated with a node
func (u *User) GetNodeTransactions(nodeID string, queryParams ...string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["trans"])

	return u.do("GET", url, "", queryParams)
}

// GetTransactions returns transactions associated with a user
func (u *User) GetTransactions(queryParams ...string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["trans"])

	return u.do("GET", url, "", queryParams)
}

// GetTransaction returns a specific transaction associated with a node
func (u *User) GetTransaction(nodeID, transactionID string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["trans"], transactionID)

	return u.do("GET", url, "", nil)
}

// CreateTransaction creates a transaction for the specified node
func (u *User) CreateTransaction(nodeID, transactionID, data string, idempotencyKey ...string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["trans"], transactionID)

	return u.do("POST", url, data, idempotencyKey)
}

// CancelTransaction deletes/cancels a transaction
func (u *User) CancelTransaction(nodeID, transactionID string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["transactions"], transactionID)

	return u.do("DELETE", url, "", nil)
}

// CommentOnTransactionStatus adds comment to the transaction status
func (u *User) CommentOnTransactionStatus(nodeID, transactionID, data string, idempotencyKey ...string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["transactions"], transactionID)

	return u.do("POST", url, data, idempotencyKey)
}

// DisputeTransaction disputes a transaction for a user
func (u *User) DisputeTransaction(nodeID, transactionID, data string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["transactions"], transactionID, "dispute")

	return u.do("PATCH", url, data, nil)
}

/********** USER **********/

// Update updates a single user and returns the updated user information
func (u *User) Update(data string) (*User, error) {
	url := buildURL(path["users"], u.UserID)

	res, err := u.do("PATCH", url, data, nil)

	mapstructure.Decode(res, u)

	u.Response = res

	return u, err
}

// CreateUBO creates and uploads an Ultimate Beneficial Ownership (UBO) and REG GG form as a physical document under the Business’s base document
func (u *User) CreateUBO(data string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, "ubo")

	return u.do("PATCH", url, data, nil)
}
