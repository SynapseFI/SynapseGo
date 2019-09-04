package synapse

import (
	"github.com/mitchellh/mapstructure"
)

/*********** GLOBAL VARIABLES ***********/

/********** TYPES **********/

type (
	// User represents a single user object
	User struct {
		AuthKey      string
		UserID       string `mapstructure:"_id"`
		RefreshToken string `mapstructure:"refresh_token"`
		Response     interface{}
		request      Request
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
	case *ActionPending:
		return readStream(response), err

	case *UnauthorizedAction:
		res, authErr := u.Authenticate(`{ "refresh_token": "`+u.RefreshToken+`" }`, u.request.fingerprint, u.request.ipAddress)

		if authErr != nil {
			return res, authErr
		}

		u.request.authKey = u.AuthKey
		return u.do(method, url, data, params)

		// case *IncorrectValues:
		// 	_, err := u.GetRefreshToken()

		// 	if err != nil {
		// 		return nil, err
		// 	}

		// 	_, err = u.Authenticate(`{ "refresh_token": "` + u.RefreshToken + `" }`)

		// 	if err != nil {
		// 		return nil, err
		// 	}

		// 	u.request.authKey = u.AuthKey

		// 	return u.do(method, url, data, params)
	}

	return readStream(response), err
}

/********** AUTHENTICATION **********/

// Authenticate returns an oauth key and sets it to the user object
// If the refresh token is expired the API will automatically send a new one
// Capture refresh token every time
func (u *User) Authenticate(data, fingerprint, ipAddress string) (map[string]interface{}, error) {
	url := buildURL(path["auth"], u.UserID)

	u.request.updateRequest(u.request.clientID, u.request.clientSecret, fingerprint, ipAddress)
	res, err := u.do("POST", url, data, nil)

	if res["refresh_token"] != nil {
		u.RefreshToken = res["refresh_token"].(string)
	}

	if res["oauth_key"] != nil {
		u.AuthKey = res["oauth_key"].(string)
		u.request.authKey = res["oauth_key"].(string)
	}

	log.info("Authenticating user...")
	return res, err
}

// GetRefreshToken performs a GET request and returns a new refresh token
func (u *User) GetRefreshToken() (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID)

	res, err := u.do("GET", url, "", nil)

	if res["refresh_token"] != nil {
		u.RefreshToken = res["refresh_token"].(string)
	}

	log.info("Getting user refresh token...")
	return res, err
}

// RegisterFingerprint submits a new fingerprint and triggers the MFA flow
func (u *User) RegisterFingerprint(fp string) (map[string]interface{}, error) {
	url := buildURL(path["auth"], u.UserID)

	u.request.fingerprint = fp

	data := `{ "refresh_token": "` + u.RefreshToken + `" }`

	res, err := u.do("POST", url, data, nil)

	log.info("Registering new fingerprint...")
	return res, err
}

// Select2FA sends the 2FA device selection to the system
func (u *User) Select2FA(device string) (map[string]interface{}, error) {
	url := buildURL(path["auth"], u.UserID)

	data := `{ "refresh_token": "` + u.RefreshToken + `", "phone_number": "` + device + `" }`

	res, err := u.do("POST", url, data, nil)

	log.info("Sending 2FA selection...")
	return res, err
}

// SubmitMFA submits the access token and mfa answer
func (u *User) SubmitMFA(data string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"])

	log.info("Submitting MFA...")
	return u.do("POST", url, data, nil)
}

// VerifyPIN sends the requested pin to the system to complete the 2FA process
func (u *User) VerifyPIN(pin string) (map[string]interface{}, error) {
	url := buildURL(path["auth"], u.UserID)

	data := `{ "refresh_token": "` + u.RefreshToken + `", "validation_pin": "` + pin + `" }`

	res, err := u.do("POST", url, data, nil)

	log.info("Sending pin verification...")
	return res, err
}

/********** NODE **********/

// GetNodes returns all of the nodes associated with a user
func (u *User) GetNodes(queryParams ...string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"])

	log.info("Getting list of user nodes...")
	return u.do("GET", url, "", queryParams)
}

// GetNode returns a single node object
func (u *User) GetNode(nodeID string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID)

	res, err := u.do("GET", url, "", nil)

	log.info("Getting user node...")
	return res, err
}

// CreateNode creates a node depending on the type of node specified
func (u *User) CreateNode(data string, idempotencyKey ...string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"])

	log.info("Creating user node...")
	return u.do("POST", url, data, idempotencyKey)
}

// UpdateNode updates a node
func (u *User) UpdateNode(nodeID, data string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID)

	log.info("Updating user node...")
	return u.do("PATCH", url, data, nil)
}

// DeleteNode deletes a node
func (u *User) DeleteNode(nodeID string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID)

	log.info("Deleting user node...")
	return u.do("DELETE", url, "", nil)
}

/********** NODE (OTHER) **********/

// VerifyMicroDeposit verifies micro-deposit amounts for a node
func (u *User) VerifyMicroDeposit(nodeID, data string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID)

	log.info("Verifying micro-deposit amounts...")
	return u.do("PATCH", url, data, nil)
}

// ReinitiateMicroDeposits reinitiates micro-deposits for an ACH-US node with AC/RT
func (u *User) ReinitiateMicroDeposits(nodeID string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID) + "?resend_micro=YES"

	log.info("Reinitiating mico-deposits...")
	return u.do("PATCH", url, "", nil)
}

// ResetCardNode resets the debit card number, card cvv, and expiration date
func (u *User) ResetCardNode(nodeID string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID) + "?reset=YES"

	log.info("Resetting card node...")
	return u.do("PATCH", url, "", nil)
}

// ShipCardNode ships a physical debit card out to the user
func (u *User) ShipCardNode(nodeID, data string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID) + "?ship=YES"

	log.info("Shipping card node...")
	return u.do("PATCH", url, data, nil)
}

// GetApplePayToken generates tokenized info for Apple Wallet
func (u *User) GetApplePayToken(nodeID, data string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, "applepay")

	log.info("Getting apple pay token...")
	return u.do("PATCH", url, data, nil)
}

/********** STATEMENT **********/

// GetStatements gets all of the user statements
func (u *User) GetStatements(queryParams ...string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["statements"])

	log.info("Getting list of user statements...")
	return u.do("GET", url, "", queryParams)
}

// GetNodeStatements gets all of the node statements
func (u *User) GetNodeStatements(nodeID string, queryParams ...string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["statements"])

	log.info("Getting list of node statements...")
	return u.do("GET", url, "", queryParams)
}

// CreateNodeStatements creates ad-hoc statements for the specified node
func (u *User) CreateNodeStatements(nodeID, data string, idempotencyKey ...string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["statements"])

	log.info("Creating ad-hoc statements")
	return u.do("POST", url, data, idempotencyKey)
}

/********** SUBNET **********/

// GetSubnets gets all subnets associated with a user
func (u *User) GetSubnets(queryParams ...string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["subnets"])

	log.info("Getting list of user subnets...")
	return u.do("GET", url, "", queryParams)
}

// GetNodeSubnets gets all subnets associated with a node
func (u *User) GetNodeSubnets(nodeID string, queryParams ...string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["subnets"])

	log.info("Getting list of node subnets...")
	return u.do("GET", url, "", queryParams)
}

// GetSubnet gets a single subnet object
func (u *User) GetSubnet(nodeID, subnetID string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["subnets"], subnetID)

	log.info("Getting node subnet...")
	return u.do("GET", url, "", nil)
}

// CreateSubnet creates a subnet object
func (u *User) CreateSubnet(nodeID, data string, idempotencyKey ...string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["subnets"])

	log.info("Creating node subnet...")
	return u.do("POST", url, data, idempotencyKey)
}

// UpdateSubnet updates a subnet object
func (u *User) UpdateSubnet(nodeID, subnetID, data string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["subnets"], subnetID)

	log.info("Updating node subnet...")
	return u.do("PATCH", url, data, nil)
}

// ShipCard ships a physical debit card out to the user
func (u *User) ShipCard(nodeID, subnetID, data string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["subnets"], subnetID, "ship")

	log.info("Shipping card...")
	return u.do("PATCH", url, data, nil)
}

/********** TRANSACTION **********/

// GetTransactions returns transactions associated with a user
func (u *User) GetTransactions(queryParams ...string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["transactions"])

	log.info("Getting list of user transactions...")
	return u.do("GET", url, "", queryParams)
}

// GetNodeTransactions returns transactions associated with a node
func (u *User) GetNodeTransactions(nodeID string, queryParams ...string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["transactions"])

	log.info("Getting list of node transactions...")
	return u.do("GET", url, "", queryParams)
}

// GetTransaction returns a specific transaction associated with a node
func (u *User) GetTransaction(nodeID, transactionID string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["transactions"], transactionID)

	log.info("Getting user transaction...")
	return u.do("GET", url, "", nil)
}

// CreateTransaction creates a transaction for the specified node
func (u *User) CreateTransaction(nodeID, data string, idempotencyKey ...string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["transactions"])

	log.info("Creating user transaction...")
	return u.do("POST", url, data, idempotencyKey)
}

// CancelTransaction deletes/cancels a transaction
func (u *User) CancelTransaction(nodeID, transactionID string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["transactions"], transactionID)

	log.info("Cancelling user transaction...")
	return u.do("DELETE", url, "", nil)
}

// CommentOnTransactionStatus adds comment to the transaction status
func (u *User) CommentOnTransactionStatus(nodeID, transactionID, data string, idempotencyKey ...string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["transactions"], transactionID)

	log.info("Adding comment to transaction status...")
	return u.do("POST", url, data, idempotencyKey)
}

// DisputeTransaction disputes a transaction for a user
func (u *User) DisputeTransaction(nodeID, transactionID, data string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["transactions"], transactionID, "dispute")

	log.info("Disputing transaction...")
	return u.do("PATCH", url, data, nil)
}

// CreateDummyTransaction triggers external dummy transactions on internal accounts
func (u *User) CreateDummyTransaction(nodeID string, queryParams ...string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID) + "/dummy-tran"

	log.info("Triggering dummy transactions...")
	return u.do("GET", url, "", queryParams)
}

/********** USER **********/

// Update updates a single user and returns the updated user information
func (u *User) Update(data string) (*User, error) {
	url := buildURL(path["users"], u.UserID)

	res, err := u.do("PATCH", url, data, nil)

	mapstructure.Decode(res, u)

	u.Response = res

	log.info("Updating user...")
	return u, err
}

// CreateUBO creates and uploads an Ultimate Beneficial Ownership (UBO) and REG GG form as a physical document under the Business’s base document
func (u *User) CreateUBO(data string) (map[string]interface{}, error) {
	url := buildURL(path["users"], u.UserID, "ubo")

	log.info("Creating UBO...")
	return u.do("PATCH", url, data, nil)
}
