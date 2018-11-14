package wrapper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

/********** GLOBAL VARIABLES **********/
const version = "v3.1"

// const _url = "https://api.synapsefi.com/" + version
const _url = "https://uat-api.synapsefi.com/" + version

const _usersURL = _url + "/users"

/********** STRUCTS **********/

// ClientCredentials structure of client object
type ClientCredentials struct {
	gateway, ipAddress, userID string
}

// NewUserData structure of new user data
type NewUserData struct {
	logins, phoneNumbers, legalNames []string
}

/********** EXPORTED FUNCTIONS ***********/

// NewClient creation of client object
func NewClient(gateway, ipAddress, userID string) ClientCredentials {
	return ClientCredentials{
		gateway:   gateway,
		ipAddress: ipAddress,
		userID:    userID,
	}
}

// HTTP METHODS //

// CreateUser POST method for creating a single user
func CreateUser(cred ClientCredentials, data []byte) ([]byte, error) {
	req := setRequest(cred, "POST", _usersURL, bytes.NewBuffer(data))

	resp := execRequest(req)

	body := readResponse(resp)

	return formatData(cred, body)
}

// GetUsers GET method to GET information about users associated with client
// *CHECK* Confirm the correct type to return from function
func GetUsers(cred ClientCredentials) ([]byte, error) {
	req := setRequest(cred, "GET", _usersURL, nil)

	resp := execRequest(req)

	body := readResponse(resp)

	return formatData(cred, body)
}

// GetUser GET method for information about single user associated with client
func GetUser(cred ClientCredentials, userID string) ([]byte, error) {
	url := _usersURL + "/" + userID

	req := setRequest(cred, "GET", url, nil)

	resp := execRequest(req)

	body := readResponse(resp)

	return formatData(cred, body)
}

/********** HELPER FUNCTIONS **********/

// executes request
func execRequest(request *http.Request) *http.Response {
	client := &http.Client{}

	fmt.Println(request)

	response, err := client.Do(request)

	if err != nil {
		fmt.Println(err)
	}

	return response
}

func formatPayload(data NewUserData) map[string]interface{} {
	return map[string]interface{}{
		"logins":        data.logins,
		"phone_numbers": data.phoneNumbers,
		"legal_names":   data.legalNames,
	}
}

func formatData(credentials ClientCredentials, responseBody []byte) ([]byte, error) {
	var data interface{}
	json.Unmarshal(responseBody, &data)

	jsonData, isOK := data.(map[string]interface{})

	// add userID as "id" to jsonData
	if isOK != false {
		jsonData["id"] = credentials.userID
	}

	return json.MarshalIndent(jsonData, "", "  ")
}

// reads response from api and returns it in readable format
func readResponse(response *http.Response) []byte {
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
	}

	return body
}

// sets client headers using client credentials
func setHeaders(credentials ClientCredentials, request *http.Request) {
	request.Header.Set("x-sp-gateway", credentials.gateway)
	request.Header.Set("x-sp-user-ip", credentials.ipAddress)
	request.Header.Set("x-sp-user", credentials.userID)
	request.Header.Set("content-type", "application/json")
}

// updates request headers with method, url, and body (if applicable)
func setRequest(credentials ClientCredentials, method, url string, body io.Reader) *http.Request {
	request, err := http.NewRequest(method, url, body)

	setHeaders(credentials, request)

	if err != nil {
		fmt.Println(err)
	}

	return request
}
