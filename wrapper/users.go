package wrapper

import (
	"fmt"
)

/*********** GLOBAL VARIABLES ***********/
const usersURL = _url + "/users"

/********** CLIENT METHODS **********/

// GenerateUser creates a new user object
func (c *Client) GenerateUser(userID string, devMode ...bool) *User {
	if len(devMode) == 1 && devMode[0] == true {
		developerMode = true
	}

	// get refresh token
	// rt := c.GetUser(userID, false)["payload"].(map[string]interface{})["refresh_token"].(string)
	payload := c.GetUser(userID, false)["payload"]
	rt := payload.(map[string]interface{})["refresh_token"].(string)

	// get auth key
	ak := auth(c, userID, rt)["payload"].(map[string]interface{})["oauth_key"].(string)

	user := &User{
		authKey:           ak,
		refreshToken:      rt,
		userID:            userID,
		clientGateway:     c.gateway,
		clientFingerprint: c.fingerprint,
		clientIP:          c.ipAddress,
		Payload:           payload,
	}

	return user
}

// GetUsers returns a list of users
func (c *Client) GetUsers(queryParams ...map[string]interface{}) map[string]interface{} {
	res, body, errs := request.
		Get(usersURL).
		Query(queryString(queryParams)).
		Set("x-sp-gateway", c.gateway).
		Set("x-sp-user-ip", c.ipAddress).
		Set("x-sp-user", c.fingerprint).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	fmt.Println(string(body))

	return responseMulti(body, "users")
}

// GetUser returns a single user
func (c *Client) GetUser(userID string, fullDehydrate bool, queryParams ...map[string]interface{}) map[string]interface{} {
	url := usersURL + "/" + userID

	if fullDehydrate != true {
		url += "?full_dehydrate=yes"
	}

	res, body, errs := request.
		Get(url).
		Query(queryString(queryParams)).
		Set("x-sp-gateway", c.gateway).
		Set("x-sp-user-ip", c.ipAddress).
		Set("x-sp-user", c.fingerprint).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	if fullDehydrate != true {
		return responseSingle(read(body), "user")
	}

	return responseSingle(read(body), "userDehydrate")
}

// CreateUser creates a single user and returns the new user data
func (c *Client) CreateUser(data string, queryParams ...map[string]interface{}) map[string]interface{} {
	res, body, errs := request.
		Post(usersURL).
		Query(queryString(queryParams)).
		Set("x-sp-gateway", c.gateway).
		Set("x-sp-user-ip", c.ipAddress).
		Set("x-sp-user", c.fingerprint).
		Send(data).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return responseSingle(read(body), "user")
}

/********** USER METHODS **********/

// UpdateUser updates a single user and returns the updated user information
func (u *User) UpdateUser(data string, queryParams ...map[string]interface{}) map[string]interface{} {
	url := usersURL + "/" + u.userID

	res, body, errs := request.
		Patch(url).
		Set("x-sp-gateway", u.clientGateway).
		Set("x-sp-user-ip", u.clientIP).
		Set("x-sp-user", u.authKey+"|"+u.clientFingerprint).
		Send(data).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return responseSingle(read(body), "user")
}

// AddNewDocuments adds new documents to a user
func (u *User) AddNewDocuments(data string) map[string]interface{} {
	url := usersURL + "/" + u.userID

	res, body, errs := request.
		Patch(url).
		Set("x-sp-gateway", u.clientGateway).
		Set("x-sp-user-ip", u.clientIP).
		Set("x-sp-user", u.authKey+"|"+u.clientFingerprint).
		Send(data).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return responseSingle(read(body), "user")
}

// UpdateExistingDocument updates existing user documents
func (u *User) UpdateExistingDocument(data string) map[string]interface{} {
	url := usersURL + "/" + u.userID

	res, body, errs := request.
		Patch(url).
		Set("x-sp-gateway", u.clientGateway).
		Set("x-sp-user-ip", u.clientIP).
		Set("x-sp-user", u.authKey+"|"+u.clientFingerprint).
		Send(data).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return responseSingle(read(body), "user")
}

// DeleteExistingDocument updates existing user documents
func (u *User) DeleteExistingDocument(data string) map[string]interface{} {
	url := usersURL + "/" + u.userID

	res, body, errs := request.
		Patch(url).
		Set("x-sp-gateway", u.clientGateway).
		Set("x-sp-user-ip", u.clientIP).
		Set("x-sp-user", u.authKey+"|"+u.clientFingerprint).
		Send(data).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return responseSingle(read(body), "user")
}

// GetNodes returns all of the nodes associated with a user
func (u *User) GetNodes(queryParams ...map[string]interface{}) map[string]interface{} {
	url := usersURL + "/" + u.userID + "/nodes"

	res, body, errs := request.
		Get(url).
		Set("x-sp-user-ip", u.clientIP).
		Set("x-sp-user", u.authKey+"|"+u.clientFingerprint).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	u.authKey = "NEW KEY"

	return responseMulti(body, "nodes")
}

// CreateDepositNode creates an deposit account
func (u *User) CreateDepositNode(data string) map[string]interface{} {
	url := usersURL + "/" + u.userID + "/nodes"

	res, body, errs := request.
		Post(url).
		Set("x-sp-user-ip", u.clientIP).
		Set("x-sp-user", u.authKey+"|"+u.clientFingerprint).
		Send(data).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return responseSingle(read(body), "node")
}

//
