package wrapper

import (
	"bytes"
	"reflect"

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
func (c *ClientCredentials) CreateUser(body []byte) map[string]interface{} {
	header := genHeader(c)
	req := createRequest(header, "POST", _usersURL, bytes.NewBuffer(body))
	res := execRequest(req)
	data := readResponse(res)

	return formatResponse(data, "user")
	// return handleRequest(c, "POST", _usersURL, bytes.NewBuffer(data))
}

// GetUsers GET method to GET information about users associated with client
func (c *ClientCredentials) GetUsers() map[string]interface{} {
	header := genHeader(c)
	req := createRequest(header, "GET", _usersURL, nil)
	res := execRequest(req)
	data := readResponse(res)

	return formatResponse(data, "users")
}

// GetUser GET method for information about single user associated with client
func (c *ClientCredentials) GetUser(userID string) map[string]interface{} {
	url := _usersURL + "/" + userID

	header := genHeader(c)
	req := createRequest(header, "GET", url, nil)
	res := execRequest(req)
	data := readResponse(res)

	return formatResponse(data, "user")
}

// HELPERS

func formatResponse(payload Payload, name string) map[string]interface{} {
	var response map[string]interface{}

	switch name {
	case "users":
		response = structs.Map(formatUsers(payload, name))

	default:
		response = structs.Map(formatUser(payload))
	}

	return response
}

func formatUser(p Payload) User {
	var user User
	user.UserID = p["_id"].(string)
	user.FullDehydrate = "yes"
	user.Payload = p

	return user
}

func formatUsers(p Payload, n string) Users {
	var users Users
	users.Limit = p["limit"].(float64)
	users.Page = p["page"].(float64)
	users.PageCount = p["page_count"].(float64)
	users.Payload = p

	list := reflect.ValueOf(p[n])

	for i := 0; i < list.Len(); i++ {
		user := list.Index(i).Interface().(map[string]interface{})
		users.UserList = append(users.UserList, formatUser(user))
	}

	return users
}

func genHeader(cred *ClientCredentials) Header {
	header := make(Header)
	header["x-sp-gateway"] = cred.gateway
	header["x-sp-user-ip"] = cred.ipAddress
	header["x-sp-user"] = cred.userID

	return header
}
