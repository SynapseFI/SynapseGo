package wrapper

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
)

/********** HELPER FUNCTIONS **********/

// executes request
func execRequest(request *http.Request) *http.Response {
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		errorLog(err)
	}

	return response
}

func formatUserObject(payload Payload) User {
	var user User

	if payload["_id"] != nil {
		user.UserID = payload["_id"].(string)
		user.FullDehydrate = "yes"
		user.Payload = payload
	}

	return user
}

func formatMultiUserObject(payload Payload, arrName string) Users {
	var users Users

	if payload[arrName] != nil {
		users.Limit = payload["limit"].(float64)
		users.Page = payload["page"].(float64)
		users.PageCount = payload["page_count"].(float64)
		users.Payload = payload
	}

	if payload[arrName] != nil {
		list := reflect.ValueOf(payload[arrName])

		for i := 0; i < list.Len(); i++ {
			var user User

			userValue := list.Index(i).Interface().(map[string]interface{})
			user.UserID = userValue["_id"].(string)
			user.FullDehydrate = "yes"
			user.Payload = userValue

			users.UserList = append(users.UserList, user)
		}
	}

	return users
}

// main handler called by wrapper methods to execute API calls
func handleRequest(credentials *ClientCredentials, httpMethod, url string, body io.Reader) User {
	request := setRequest(credentials, httpMethod, url, body)

	response := execRequest(request)

	responseData := readResponse(response)

	return formatUserObject(responseData)
}

// main handler called by wrapper methods that return multiple users in payload
func handleRequestMulti(credentials *ClientCredentials, httpMethod, url, arrName string, body io.Reader) Users {
	request := setRequest(credentials, httpMethod, url, body)

	response := execRequest(request)

	responseData := readResponse(response)

	return formatMultiUserObject(responseData, arrName)
}

// reads response from api and returns it in readable format
func readResponse(response *http.Response) map[string]interface{} {
	var payload interface{}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		errorLog(err)
	}

	json.Unmarshal(body, &payload)

	return payload.(map[string]interface{})
}

func makeHeader(h Header, r *http.Request) {
	for k := range h {
		r.Header.Set(k, h[k])
	}

	r.Header.Set("content-type", "application/json")
}

func createRequest(headers Header, httpMethod, url string, body io.Reader) *http.Request {
	request, err := http.NewRequest(httpMethod, url, body)

	if err != nil {
		errorLog(err)
	}

	makeHeader(headers, request)

	return request
}

// sets client headers using client credentials
func setHeaders(credentials *ClientCredentials, request *http.Request) {
	request.Header.Set("x-sp-gateway", credentials.gateway)
	request.Header.Set("x-sp-user-ip", credentials.ipAddress)
	request.Header.Set("x-sp-user", credentials.userID)
	request.Header.Set("content-type", "application/json")
}

// updates request headers with method, url, and body (if applicable)
func setRequest(credentials *ClientCredentials, httpMethod, url string, body io.Reader) *http.Request {
	request, err := http.NewRequest(httpMethod, url, body)
	setHeaders(credentials, request)

	if err != nil {
		errorLog(err)
	}

	return request
}
