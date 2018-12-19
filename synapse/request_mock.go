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

/********** METHODS **********/

func (req *Request) updateRequest(clientID, clientSecret, fingerprint, ipAddress string, authKey ...string) Request {
	var aKey string

	if len(authKey) > 0 {
		aKey = authKey[0]
	}

	return Request{
		authKey:      aKey,
		clientID:     clientID,
		clientSecret: clientSecret,
		fingerprint:  fingerprint,
		ipAddress:    ipAddress,
	}
}

/********** REQUEST **********/

// Get performs a GET request
func (req *Request) Get(url string, queryParams []string) ([]byte, error) {
	var body []byte

	return body, nil
}

// Post performs a POST request
func (req *Request) Post(url, data string, queryParams []string) ([]byte, error) {
	var body []byte

	return body, nil
}

// Patch performs a PATCH request
func (req *Request) Patch(url, data string, queryParams []string) ([]byte, error) {
	var body []byte

	return body, nil
}

// Delete performs a DELETE request
func (req *Request) Delete(url string) ([]byte, error) {
	var body []byte

	return body, nil
}
