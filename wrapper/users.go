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
	h := c.getHeaderInfo("")
	r := apiRequest(GET, usersURL, h, queryParams)

	return responseMulti(r, "users")
}

// GetUser returns a single user
func (c *Client) GetUser(UserID string, fullDehydrate bool, queryParams ...string) *User {
	url := usersURL + "/" + UserID

	if fullDehydrate != true {
		url += "?full_dehydrate=yes"
	}

	h := c.getHeaderInfo("")
	r := apiRequest(GET, url, h, queryParams)
	return generateUser(c, r, fullDehydrate)
}

// CreateUser creates a single user and returns the new user data
func (c *Client) CreateUser(data string, queryParams ...string) *User {
	h := c.getHeaderInfo("")
	r := apiRequest(POST, usersURL, h, queryParams, data)

	return generateUser(c, r, false)
}

/********** USER METHODS **********/

// Update updates a single user and returns the updated user information
func (u *User) Update(data string, queryParams ...string) map[string]interface{} {
	url := usersURL + "/" + u.UserID

	h := u.getHeaderInfo("")
	r := apiRequest(PATCH, url, h, queryParams, data)

	return responseSingle(r, "user")
}

// AddNewDocuments adds new documents to a user
func (u *User) AddNewDocuments(data string) map[string]interface{} {
	url := usersURL + "/" + u.UserID

	h := u.getHeaderInfo("")
	r := apiRequest(PATCH, url, h, nil, data)

	return responseSingle(r, "user")
}

// UpdateExistingDocument updates existing user documents
func (u *User) UpdateExistingDocument(data string) map[string]interface{} {
	url := usersURL + "/" + u.UserID

	h := u.getHeaderInfo("")
	r := apiRequest(PATCH, url, h, nil, data)

	return responseSingle(r, "user")
}

// DeleteExistingDocument updates existing user documents
func (u *User) DeleteExistingDocument(data string) map[string]interface{} {
	url := usersURL + "/" + u.UserID

	h := u.getHeaderInfo("")
	r := apiRequest(PATCH, url, h, nil, data)

	return responseSingle(r, "user")
}

// GetNodes returns all of the nodes associated with a user
func (u *User) GetNodes(queryParams ...string) map[string]interface{} {
	url := usersURL + "/" + u.UserID + "/nodes"

	h := u.getHeaderInfo("no gateway")
	r := apiRequest(PATCH, url, h, queryParams)

	u.AuthKey = "NEW KEY"

	return responseMulti(r, "nodes")
}

// CreateDepositNode creates an deposit account
func (u *User) CreateDepositNode(data string) map[string]interface{} {
	url := usersURL + "/" + u.UserID + "/nodes"

	h := u.getHeaderInfo("no gateway")
	r := apiRequest(PATCH, url, h, nil, data)

	return responseSingle(r, "node")
}
