package wrapper

/*********** GLOBAL VARIABLES ***********/
const usersURL = _url + "/users"

/********** METHODS **********/

// GetUsers returns a list of users
func (c *Client) GetUsers(queryParams ...map[string]interface{}) map[string]interface{} {
	res, body, errs := request.
		Get(usersURL).
		Query(queryString(queryParams)).
		Set("x-sp-gateway", c.gateway).
		Set("x-sp-user-ip", c.ipAddress).
		Set("x-sp-user", c.userID).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

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
		Set("x-sp-user", c.userID).
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
		Set("x-sp-user", c.userID).
		Send(data).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return responseSingle(read(body), "user")
}
