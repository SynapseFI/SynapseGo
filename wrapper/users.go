package wrapper

/*********** GLOBAL VARIABLES ***********/
const usersURL = _url + "/users"

/********** METHODS **********/

// GetUsers returns a list of users
func (c *ClientCredentials) GetUsers() map[string]interface{} {
	header(c, "")

	res, body, errs := request.
		Get(usersURL).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return format(body)
}

// GetUser returns a single user
func (c *ClientCredentials) GetUser(userID string) map[string]interface{} {
	url := usersURL + "/" + userID

	header(c, "")

	res, body, errs := request.
		Get(url).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return format(body)
}

// CreateUser creates a single user and returns the new user data
func (c *ClientCredentials) CreateUser(data string) map[string]interface{} {
	header(c, "")

	res, body, errs := request.
		Send(data).
		Post(usersURL).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return format(body)
}
