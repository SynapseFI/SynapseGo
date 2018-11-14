package wrapper

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

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

// GetUser GET method to GET user information
// Confirm the correct type to return from function
func GetUser(credentials ClientCredentials) map[string]interface{} {
	_url := "http://uat-api.synapsefi.com/v3.1/users"

	client := &http.Client{}

	req := updateRequest(credentials, "GET", _url, nil)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	var data interface{}
	json.Unmarshal(body, &data)

	jsonData, ok := data.(map[string]interface{})

	if ok != false {
		jsonData["id"] = credentials.userID
	}

	return jsonData
}

/********** HELPER FUNCTIONS **********/

// sets client headers using client credentials
func setHeaders(client ClientCredentials, request *http.Request) {
	request.Header.Set("x-sp-gateway", client.gateway)
	request.Header.Set("x-sp-user-ip", client.ipAddress)
	request.Header.Set("x-sp-user", client.userID)
	request.Header.Set("content-type", "application/json;charset=UTF-8")
}

// updates request headers with method, url, and body (if applicable)
func updateRequest(client ClientCredentials, method, url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)

	setHeaders(client, req)

	if err != nil {
		fmt.Println(err)
	}

	return req
}
