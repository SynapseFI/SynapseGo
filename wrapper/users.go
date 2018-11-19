package wrapper

/*********** GLOBAL VARIABLES ***********/
const usersURL = _url + "/users"

/********** METHODS **********/

// GetUsers returns a list of users
func (c *ClientCredentials) GetUsers() map[string]interface{} {
	// header(c, "")

	res, body, errs := request.
		Get(usersURL).
		Set("x-sp-gateway", c.gateway).
		Set("x-sp-user-ip", c.ipAddress).
		Set("x-sp-user", c.userID).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return multiData(body, "users")
}

// GetUser returns a single user
func (c *ClientCredentials) GetUser(userID string) map[string]interface{} {
	url := usersURL + "/" + userID

	res, body, errs := request.
		Get(url).
		Set("x-sp-gateway", c.gateway).
		Set("x-sp-user-ip", c.ipAddress).
		Set("x-sp-user", c.userID).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return singleData(read(body), "user")
}

// CreateUser creates a single user and returns the new user data
func (c *ClientCredentials) CreateUser(data string) map[string]interface{} {
	// header(c, "")

	res, body, errs := request.
		Post(usersURL).
		Set("x-sp-gateway", c.gateway).
		Set("x-sp-user-ip", c.ipAddress).
		Set("x-sp-user", c.userID).
		Send(data).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return multiData(body, "users")
}
