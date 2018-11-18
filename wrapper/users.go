package wrapper

/*********** GLOBAL VARIABLES ***********/
const usersURL = _url + "/users"

// GetUsers returns a list of users
func (c *ClientCredentials) GetUsers() map[string]interface{} {
	response, body, errs := request.Get(usersURL).
		Set("x-sp-gateway", c.gateway).
		Set("x-sp-user-ip", c.ipAddress).
		Set("x-sp-user", c.userID).
		EndBytes()

	if response != nil {
		// fmt.Println(response)
	}

	if errs != nil {
		errorLog(errs)
	}

	return format(body)
}

// GetUser returns a single user
func (c *ClientCredentials) GetUser(userID string) map[string]interface{} {
	url := usersURL + "/" + userID

	response, bytesBody, errs := request.Get(url).
		Set("x-sp-gateway", c.gateway).
		Set("x-sp-user-ip", c.ipAddress).
		Set("x-sp-user", c.userID).
		EndBytes()

	if response != nil {
	}

	if errs != nil {
		errorLog(errs)
	}

	return format(bytesBody)
}

// CreateUser creates a single user and returns the new user data
func (c *ClientCredentials) CreateUser(data string) map[string]interface{} {
	response, bytesBody, errs := request.Post(usersURL).
		Set("x-sp-gateway", c.gateway).
		Set("x-sp-user-ip", c.ipAddress).
		Set("x-sp-user", c.userID).
		Send(data).
		EndBytes()

	if response != nil {
	}

	if errs != nil {
		errorLog(errs)
	}

	return format(bytesBody)
}
