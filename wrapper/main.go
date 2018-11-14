package wrapper

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

/********** GLOBAL VARIABLES **********/
const _url = "http://uat-api.synapsefi.com/v3.1"

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
// *CHECK* Confirm the correct type to return from function
func GetUser(cred ClientCredentials) map[string]interface{} {
	url := _url + "/users"

	c := &http.Client{}

	req := updateRequest(cred, "GET", url, nil)

	resp := execRequest(c, req)

	body := readResponse(resp)

	return jsonifyData(cred, body)
}

/********** HELPER FUNCTIONS **********/

// executes request
func execRequest(client *http.Client, request *http.Request) *http.Response {
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
