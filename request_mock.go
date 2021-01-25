// +build mock

package synapse

import (
	"github.com/parnurzeal/gorequest"
)

/********** TYPES **********/

type (
	// Request represents the http request client
	Request struct {
		authKey, clientID, clientSecret, fingerprint, ipAddress string
	}
)

/********** GLOBAL VARIABLES **********/
var goreq = gorequest.New()
var mockResponse = []byte(`{
	"message": "This is a mock response"
}`)

/********** METHODS **********/

func (req *Request) updateRequest(clientID, clientSecret, fingerprint, ipAddress string, authKey ...string) {
	if len(authKey) > 0 {
		req.authKey = authKey[0]
	}

	req.clientID = clientID
	req.clientSecret = clientSecret
	req.fingerprint = fingerprint
	req.ipAddress = ipAddress
}

/********** REQUEST **********/

// Get performs a GET request
func (req *Request) Get(url string, queryParams []string) ([]byte, error) {
	return mockResponse, nil
}

// Post performs a POST request
func (req *Request) Post(url, data string, queryParams []string) ([]byte, error) {
	return mockResponse, nil
}

// Patch performs a PATCH request
func (req *Request) Patch(url, data string, queryParams []string) ([]byte, error) {
	return mockResponse, nil
}

// Delete performs a DELETE request
func (req *Request) Delete(url string) ([]byte, error) {
	return mockResponse, nil
}
