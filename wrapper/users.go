package wrapper

import (
	"bytes"

	"github.com/fatih/structs"
)

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
	header := genHeader(c)
	req := createRequest(header, "GET", _usersURL, nil)
	res := execRequest(req)
	data := readResponse(res)

	return formatMultiUserObject(data, "users")

	// return handleRequestMulti(c, "GET", _usersURL, "users", nil)
}

// GetUser GET method for information about single user associated with client
func (c *ClientCredentials) GetUser(userID string) User {
	url := _usersURL + "/" + userID

	header := genHeader(c)
	req := createRequest(header, "GET", url, nil)
	res := execRequest(req)
	data := readResponse(res)

	return formatResponse(data, "user")
	// return testResponse(data, "user")
	// return handleRequest(c, "GET", url, nil)
}

// HELPERS

func testResponse(payload Payload, name string) map[string]interface{} {
	// var response UserTest
	var user User

	switch name {
	case "user":
		// response["userID"] = payload["_id"].(string)
		// response["fullDehydrate"] = "yes"
		// response["payload"] = payload
		// }

		// user := UserTest{
		// 	"userID":        payload["_id"].(string),
		// 	"fullDehydrate": "yes",
		// 	"payload":       payload,
		// }
		// if payload["_id"] != nil {
		user.UserID = payload["_id"].(string)
		user.FullDehydrate = "yes"
		user.Payload = payload
	}

	// return response
	return structs.Map(user)
}

func formatResponse(payload Payload, name string) User {
	var user User
	if payload["_id"] != nil {
		user.UserID = payload["_id"].(string)
		user.FullDehydrate = "yes"
		user.Payload = payload
	}

	return user
}

func genHeader(cred *ClientCredentials) Header {
	header := make(Header)
	header["x-sp-gateway"] = cred.gateway
	header["x-sp-user-ip"] = cred.ipAddress
	header["x-sp-user"] = cred.userID

	return header
}
