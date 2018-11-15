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
func CreateUser(cred ClientCredentials, data []byte) ([]byte, error) {
	req := setRequest(cred, "POST", _usersURL, bytes.NewBuffer(data))

	resp := execRequest(req)

	body := readResponse(resp)

	return formatResponse(cred, body)
}

// GetUsers GET method to GET information about users associated with client
// *CHECK* Confirm the correct type to return from function
func GetUsers(cred ClientCredentials) ([]byte, error) {
	req := setRequest(cred, "GET", _usersURL, nil)

	resp := execRequest(req)

	body := readResponse(resp)

	return formatResponse(cred, body)
}

// GetUser GET method for information about single user associated with client
func GetUser(cred ClientCredentials, userID string) ([]byte, error) {
	url := _usersURL + "/" + userID

	req := setRequest(cred, "GET", url, nil)

	resp := execRequest(req)

	body := readResponse(resp)

	return formatResponse(cred, body)
}
