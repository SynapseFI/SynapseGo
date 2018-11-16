package wrapper

import "bytes"

/********** GLOBAL VARIABLES **********/
const _usersURL = _url + "/users"

/********** STRUCTS **********/

// NewUserData structure of new user data
type NewUserData struct {
	logins, phoneNumbers, legalNames []string
}

/********** METHODS **********/

// CreateUser POST method for creating a single user
func (c *ClientCredentials) CreateUser(data []byte) User {
	return handleRequest(c, "POST", _usersURL, bytes.NewBuffer(data))
}

// GetUsers GET method to GET information about users associated with client
// *CHECK* Confirm the correct type to return from function
func (c *ClientCredentials) GetUsers() Users {
	return handleRequestMulti(c, "GET", _usersURL, "users", nil)
}

// GetUser GET method for information about single user associated with client
func (c *ClientCredentials) GetUser(userID string) User {
	url := _usersURL + "/" + userID

	return handleRequest(c, "GET", url, nil)
}
