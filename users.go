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
	log.info("========== AUTHENTICATE ==========")
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

	return res, err
}

// GetRefreshToken performs a GET request and returns a new refresh token
func (u *User) GetRefreshToken() (map[string]interface{}, error) {
	log.info("========== GET REFRESH TOKEN ==========")
	url := buildURL(path["users"], u.UserID)

	res, err := u.do("GET", url, "", nil)

	if res["refresh_token"] != nil {
		u.RefreshToken = res["refresh_token"].(string)
	}

	return res, err
}

// RegisterFingerprint submits a new fingerprint and triggers the MFA flow
func (u *User) RegisterFingerprint(fp string) (map[string]interface{}, error) {
	log.info("========== REGISTER FINGERPRINT ==========")
	url := buildURL(path["auth"], u.UserID)

	u.request.fingerprint = fp

	data := `{ "refresh_token": "` + u.RefreshToken + `" }`

	res, err := u.do("POST", url, data, nil)

	return res, err
}

// Select2FA sends the 2FA device selection to the API
func (u *User) Select2FA(device string) (map[string]interface{}, error) {
	log.info("========== SELECT 2FA DEVICE ==========")
	url := buildURL(path["auth"], u.UserID)

	data := `{ "refresh_token": "` + u.RefreshToken + `", "phone_number": "` + device + `" }`

	res, err := u.do("POST", url, data, nil)

	return res, err
}

// SubmitMFA submits the access token and mfa answer
func (u *User) SubmitMFA(data string) (map[string]interface{}, error) {
	log.info("========== SUBMIT MFA RESPONSE ==========")
	url := buildURL(path["users"], u.UserID, path["nodes"])

	return u.do("POST", url, data, nil)
}

// VerifyPIN sends the requested pin to the API to complete the 2FA process
func (u *User) VerifyPIN(pin string) (map[string]interface{}, error) {
	log.info("========== VERIFY PIN ==========")
	url := buildURL(path["auth"], u.UserID)

	data := `{ "refresh_token": "` + u.RefreshToken + `", "validation_pin": "` + pin + `" }`

	res, err := u.do("POST", url, data, nil)

	return res, err
}

/********** NODE **********/

// GetNodes returns all of the nodes associated with a user
func (u *User) GetNodes(queryParams ...string) (map[string]interface{}, error) {
	log.info("========== GET USER NODES ==========")
	url := buildURL(path["users"], u.UserID, path["nodes"])

	return u.do("GET", url, "", queryParams)
}

// GetNode returns a single node object
func (u *User) GetNode(nodeID string) (map[string]interface{}, error) {
	log.info("========== GET NODE ==========")
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID)

	res, err := u.do("GET", url, "", nil)

	return res, err
}

// CreateNode creates a node depending on the type of node specified
func (u *User) CreateNode(data string, idempotencyKey ...string) (map[string]interface{}, error) {
	log.info("========== CREATE NODE ==========")
	url := buildURL(path["users"], u.UserID, path["nodes"])

	return u.do("POST", url, data, idempotencyKey)
}

// UpdateNode updates a node
func (u *User) UpdateNode(nodeID, data string) (map[string]interface{}, error) {
	log.info("========== UPDATE NODE ==========")
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID)

	return u.do("PATCH", url, data, nil)
}

// DeleteNode deletes a node
func (u *User) DeleteNode(nodeID string) (map[string]interface{}, error) {
	log.info("========== DELETE NODE ==========")
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID)

	return u.do("DELETE", url, "", nil)
}

/********** NODE (OTHER) **********/

// VerifyMicroDeposit verifies micro-deposit amounts for a node
func (u *User) VerifyMicroDeposit(nodeID, data string) (map[string]interface{}, error) {
	log.info("========== VERIFY MICRO-DEPOSITS ==========")
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID)

	return u.do("PATCH", url, data, nil)
}

// ReinitiateMicroDeposits reinitiates micro-deposits for an ACH-US node with AC/RT
func (u *User) ReinitiateMicroDeposits(nodeID string) (map[string]interface{}, error) {
	log.info("========== RE-INITIATE MICRO-DEPOSITS ==========")
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID) + "?resend_micro=YES"

	return u.do("PATCH", url, "", nil)
}

// ResetCardNode resets the debit card number, card cvv, and expiration date
func (u *User) ResetCardNode(nodeID string) (map[string]interface{}, error) {
	log.info("========== RESET CARD ==========)")
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID) + "?reset=YES"

	return u.do("PATCH", url, "", nil)
}

// ShipCardNode ships a physical debit card out to the user
func (u *User) ShipCardNode(nodeID, data string) (map[string]interface{}, error) {
	log.info("========== SHIP CARD ==========")
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID) + "?ship=YES"

	return u.do("PATCH", url, data, nil)
}

// GetApplePayToken generates tokenized info for Apple Wallet
func (u *User) GetApplePayToken(nodeID, data string) (map[string]interface{}, error) {
	log.info("========== GET APPLE PAY TOKEN ==========")
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, "applepay")

	return u.do("PATCH", url, data, nil)
}

/********** STATEMENT **********/

// GetStatements gets all of the user statements
func (u *User) GetStatements(queryParams ...string) (map[string]interface{}, error) {
	log.info("========== GET USER STATEMENTS ==========")
	url := buildURL(path["users"], u.UserID, path["statements"])

	return u.do("GET", url, "", queryParams)
}

// GetNodeStatements gets all of the node statements
func (u *User) GetNodeStatements(nodeID string, queryParams ...string) (map[string]interface{}, error) {
	log.info("========== GET NODE STATEMENTS ==========")
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["statements"])

	return u.do("GET", url, "", queryParams)
}

// CreateNodeStatements creates ad-hoc statements for the specified node
func (u *User) CreateNodeStatements(nodeID, data string, idempotencyKey ...string) (map[string]interface{}, error) {
	log.info("========== CREATE AD-HOC STATEMENT ==========")
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["statements"])

	return u.do("POST", url, data, idempotencyKey)
}

/********** SUBNET **********/

// GetSubnets gets all subnets associated with a user
func (u *User) GetSubnets(queryParams ...string) (map[string]interface{}, error) {
	log.info("========== GET USER SUBNETS ==========")
	url := buildURL(path["users"], u.UserID, path["subnets"])

	return u.do("GET", url, "", queryParams)
}

// GetNodeSubnets gets all subnets associated with a node
func (u *User) GetNodeSubnets(nodeID string, queryParams ...string) (map[string]interface{}, error) {
	log.info("========== GET NODE SUBNETS ==========")
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["subnets"])

	return u.do("GET", url, "", queryParams)
}

// GetSubnet gets a single subnet object
func (u *User) GetSubnet(nodeID, subnetID string) (map[string]interface{}, error) {
	log.info("========== GET SUBNET ==========")
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["subnets"], subnetID)

	return u.do("GET", url, "", nil)
}

// CreateSubnet creates a subnet object
func (u *User) CreateSubnet(nodeID, data string, idempotencyKey ...string) (map[string]interface{}, error) {
	log.info("========== CREATE SUBNET ==========")
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["subnets"])

	return u.do("POST", url, data, idempotencyKey)
}

// UpdateSubnet updates a subnet object
func (u *User) UpdateSubnet(nodeID, subnetID, data string) (map[string]interface{}, error) {
	log.info("========== UPDATE SUBNET ==========")
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["subnets"], subnetID)

	return u.do("PATCH", url, data, nil)
}

// ShipCard ships a physical debit card out to the user
func (u *User) ShipCard(nodeID, subnetID, data string) (map[string]interface{}, error) {
	log.info("========== SHIP CARD SUBNET ==========")
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["subnets"], subnetID, "ship")

	return u.do("PATCH", url, data, nil)
}

// GetCardShipment gets card shipment details
func (u *User) GetCardShipment(nodeID, subnetID string) (map[string]interface{}, error) {
	log.info("==========  GET CARD SHIPMENT DETAILS ==========")
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["subnets"], subnetID, "ship")

	return u.do("GET", url, "", nil)
}

/********** TRANSACTION **********/

// GetTransactions returns transactions associated with a user
func (u *User) GetTransactions(queryParams ...string) (map[string]interface{}, error) {
	log.info("========== GET USER TRANSACTIONS ==========")
	url := buildURL(path["users"], u.UserID, path["transactions"])

	return u.do("GET", url, "", queryParams)
}

// GetNodeTransactions returns transactions associated with a node
func (u *User) GetNodeTransactions(nodeID string, queryParams ...string) (map[string]interface{}, error) {
	log.info("========== GET NODE TRANSACTIONS ==========")
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["transactions"])

	return u.do("GET", url, "", queryParams)
}

// GetTransaction returns a specific transaction associated with a node
func (u *User) GetTransaction(nodeID, transactionID string) (map[string]interface{}, error) {
	log.info("========== GET TRANSACTION ==========")
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["transactions"], transactionID)

	return u.do("GET", url, "", nil)
}

// CreateTransaction creates a transaction for the specified node
func (u *User) CreateTransaction(nodeID, data string, idempotencyKey ...string) (map[string]interface{}, error) {
	log.info("========== CREATE TRANSACTION ==========")
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["transactions"])

	return u.do("POST", url, data, idempotencyKey)
}

// CancelTransaction deletes/cancels a transaction
func (u *User) CancelTransaction(nodeID, transactionID string) (map[string]interface{}, error) {
	log.info("========== CANCEL TRANSACTION ==========")
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["transactions"], transactionID)

	return u.do("DELETE", url, "", nil)
}

// CommentOnTransactionStatus adds comment to the transaction status
func (u *User) CommentOnTransactionStatus(nodeID, transactionID, data string, idempotencyKey ...string) (map[string]interface{}, error) {
	log.info("========== COMMENT ON TRANSACTION STATUS ==========")
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["transactions"], transactionID)

	return u.do("POST", url, data, idempotencyKey)
}

// DisputeTransaction disputes a transaction for a user
func (u *User) DisputeTransaction(nodeID, transactionID, data string) (map[string]interface{}, error) {
	log.info("========== DISPUTE TRANSACTION ==========")
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID, path["transactions"], transactionID, "dispute")

	return u.do("PATCH", url, data, nil)
}

// CreateDummyTransaction triggers external dummy transactions on internal accounts
func (u *User) CreateDummyTransaction(nodeID string, queryParams ...string) (map[string]interface{}, error) {
	log.info("========== CREATE DUMMY TRANSACTION ==========")
	url := buildURL(path["users"], u.UserID, path["nodes"], nodeID) + "/dummy-tran"

	return u.do("GET", url, "", queryParams)
}

/********** USER **********/

// Update updates a single user and returns the updated user information
func (u *User) Update(data string) (*User, error) {
	log.info("========== UPDATE USER ==========")
	url := buildURL(path["users"], u.UserID)

	res, err := u.do("PATCH", url, data, nil)
	mapstructure.Decode(res, u)
	u.Response = res

	return u, err
}

// CreateUBO creates and uploads an Ultimate Beneficial Ownership (UBO) and REG GG form as a physical document under the Business’s base document
func (u *User) CreateUBO(data string) (map[string]interface{}, error) {
	log.info("========== CREATE UBO ==========")
	url := buildURL(path["users"], u.UserID, "ubo")

	return u.do("PATCH", url, data, nil)
}
