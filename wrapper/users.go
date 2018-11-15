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
func CreateUser(cred ClientCredentials, data []byte) Response {
	return handleRequest(cred, "POST", _usersURL, bytes.NewBuffer(data))
}

// GetUsers GET method to GET information about users associated with client
// *CHECK* Confirm the correct type to return from function
func GetUsers(cred ClientCredentials) Response {
	return handleRequest(cred, "GET", _usersURL, nil)
}

// GetUser GET method for information about single user associated with client
func GetUser(cred ClientCredentials, userID string) Response {
	url := _usersURL + "/" + userID

	return handleRequest(cred, "GET", url, nil)
}
