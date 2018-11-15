package wrapper

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

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

func formatResponse(credentials ClientCredentials, response []byte) ([]byte, error) {
	var payload interface{}
	json.Unmarshal(response, &payload)

	data, isOK := payload.(map[string]interface{})

	// add userID as "id" to jsonData
	if isOK != false {
		data["id"] = credentials.userID
	}

	return json.MarshalIndent(payload, "", "  ")
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

func handleRequest(credentials ClientCredentials, httpMethod, url string, body io.Reader) ([]byte, error) {
	request := setRequest(credentials, httpMethod, url, body)

	response := execRequest(request)

	responseData := readResponse(response)

	return formatResponse(credentials, responseData)
}
