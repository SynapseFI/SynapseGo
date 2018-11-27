package wrapper

/*********** GLOBAL VARIABLES ***********/
const usersURL = _url + "/users"

/********** CLIENT METHODS **********/

// GenerateUser creates a new user object
func generateUser(c *Client, data []byte, dehydrate bool) *User {
	d := read(data)

	// get refresh token
	rt := d["refresh_token"].(string)

	// get auth key
	// ak := auth(c, d["_id"].(string), rt)["payload"].(map[string]interface{})["oauth_key"].(string)

	return &User{
		// AuthKey:           ak,
		clientGateway:     c.gateway,
		clientFingerprint: c.fingerprint,
		clientIP:          c.ipAddress,
		fullDehydrate:     dehydrate,
		RefreshToken:      rt,
		UserID:            d["_id"].(string),
		Payload:           d,
	}
}

// GetUsers returns a list of users
func (c *Client) GetUsers(queryParams ...string) map[string]interface{} {
	headers := map[string]interface{}{
		"x-sp-gateway": c.gateway,
		"x-sp-user-ip": c.ipAddress,
		// "x-sp-user":    c.fingerprint,
	}

	r := apiRequest(GET, usersURL, headers, queryParams)

	return responseMulti(r, "users")
}

// GetUser returns a single user
func (c *Client) GetUser(UserID string, fullDehydrate bool, queryParams ...map[string]interface{}) *User {
	url := usersURL + "/" + UserID

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

	return generateUser(c, body, fullDehydrate)
}

// CreateUser creates a single user and returns the new user data
func (c *Client) CreateUser(data string, queryParams ...map[string]interface{}) *User {
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

	return generateUser(c, body, false)
}

/********** USER METHODS **********/

// Update updates a single user and returns the updated user information
func (u *User) Update(data string, queryParams ...map[string]interface{}) map[string]interface{} {
	url := usersURL + "/" + u.UserID

	res, body, errs := request.
		Patch(url).
		Set("x-sp-gateway", u.clientGateway).
		Set("x-sp-user-ip", u.clientIP).
		Set("x-sp-user", u.AuthKey+"|"+u.clientFingerprint).
		Send(data).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return responseSingle(body, "user")
}

// AddNewDocuments adds new documents to a user
func (u *User) AddNewDocuments(data string) map[string]interface{} {
	url := usersURL + "/" + u.UserID

	res, body, errs := request.
		Patch(url).
		Set("x-sp-gateway", u.clientGateway).
		Set("x-sp-user-ip", u.clientIP).
		Set("x-sp-user", u.AuthKey+"|"+u.clientFingerprint).
		Send(data).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return responseSingle(body, "user")
}

// UpdateExistingDocument updates existing user documents
func (u *User) UpdateExistingDocument(data string) map[string]interface{} {
	url := usersURL + "/" + u.UserID

	res, body, errs := request.
		Patch(url).
		Set("x-sp-gateway", u.clientGateway).
		Set("x-sp-user-ip", u.clientIP).
		Set("x-sp-user", u.AuthKey+"|"+u.clientFingerprint).
		Send(data).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return responseSingle(body, "user")
}

// DeleteExistingDocument updates existing user documents
func (u *User) DeleteExistingDocument(data string) map[string]interface{} {
	url := usersURL + "/" + u.UserID

	res, body, errs := request.
		Patch(url).
		Set("x-sp-gateway", u.clientGateway).
		Set("x-sp-user-ip", u.clientIP).
		Set("x-sp-user", u.AuthKey+"|"+u.clientFingerprint).
		Send(data).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return responseSingle(body, "user")
}

// GetNodes returns all of the nodes associated with a user
func (u *User) GetNodes(queryParams ...map[string]interface{}) map[string]interface{} {
	url := usersURL + "/" + u.UserID + "/nodes"

	res, body, errs := request.
		Get(url).
		Set("x-sp-user-ip", u.clientIP).
		Set("x-sp-user", u.AuthKey+"|"+u.clientFingerprint).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	u.AuthKey = "NEW KEY"

	return responseMulti(body, "nodes")
}

// CreateDepositNode creates an deposit account
func (u *User) CreateDepositNode(data string) map[string]interface{} {
	url := usersURL + "/" + u.UserID + "/nodes"

	res, body, errs := request.
		Post(url).
		Set("x-sp-user-ip", u.clientIP).
		Set("x-sp-user", u.AuthKey+"|"+u.clientFingerprint).
		Send(data).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return responseSingle(body, "node")
}
