package wrapper

import (
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

/********** STRUCTS **********/

// ClientCredentials structure of client object
type ClientCredentials struct {
	gateway, ipAddress, userID string
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

// GetUser GET method for information about single user associated with client
func GetUser(cred ClientCredentials, userID string) map[string]interface{} {
	url := _url + "/users/" + userID

	req := updateRequest(cred, "GET", url, nil)

	resp := execRequest(req)

	body := readResponse(resp)

	return jsonifyData(cred, body)
}

// GetUsers GET method to GET information about users associated with client
// *CHECK* Confirm the correct type to return from function
func GetUsers(cred ClientCredentials) map[string]interface{} {
	url := _url + "/users"

	req := updateRequest(cred, "GET", url, nil)

	resp := execRequest(req)

	body := readResponse(resp)

	return jsonifyData(cred, body)
}

/********** HELPER FUNCTIONS **********/

// executes request
func execRequest(request *http.Request) *http.Response {
	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		fmt.Println(err)
	}

	return response
}

func jsonifyData(credentials ClientCredentials, responseBody []byte) map[string]interface{} {
	var data interface{}
	json.Unmarshal(responseBody, &data)

	jsonData, isOK := data.(map[string]interface{})

	// add userID as "id" to jsonData
	if isOK != false {
		jsonData["id"] = credentials.userID
	}

	return jsonData

}

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
	request.Header.Set("content-type", "application/json;charset=UTF-8")
}

// updates request headers with method, url, and body (if applicable)
func updateRequest(credentials ClientCredentials, method, url string, body io.Reader) *http.Request {
	request, err := http.NewRequest(method, url, body)

	setHeaders(credentials, request)

	if err != nil {
		fmt.Println(err)
	}

	return request
}
