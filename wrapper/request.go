package wrapper

import (
	"github.com/parnurzeal/gorequest"
)

/********** GLOBAL VARIABLES **********/
const version = "v3.2"

// const _url = "https://api.synapsefi.com/" + version
const _url = "https://uat-api.synapsefi.com/" + version

var goreq = gorequest.New()
var request *Request

/********** TYPES **********/

type (
	// Request represents the http request client
	Request struct {
		fingerprint, gateway, ipAddress, authKey string
	}
)

/********** METHODS **********/

func newRequest(clientID, clientSecret, fingerprint, ipAddress string) *Request {
	return &Request{
		fingerprint: "|" + fingerprint,
		gateway:     clientID + "|" + clientSecret,
		ipAddress:   ipAddress,
	}
}

/********** REQUEST **********/

// Get performs a GET request
func (req *Request) Get(url, params string, result interface{}) ([]byte, error) {
	res, body, errs := goreq.
		Get(url).
		Set("x-sp-gateway", req.gateway).
		Set("x-sp-user-ip", req.ipAddress).
		Set("x-sp-user", req.fingerprint).
		Query(params).
		EndStruct(result)

	if res.StatusCode != 200 || len(errs) > 0 {
		return nil, handleHTTPError(body)
	}

	return body, nil
}

// Post performs a POST request
func (req *Request) Post(url, data, params string, result interface{}) ([]byte, error) {

	res, body, errs := goreq.
		Post(url).
		Set("x-sp-gateway", req.gateway).
		Set("x-sp-user-ip", req.ipAddress).
		Set("x-sp-user", req.fingerprint).
		Query(params).
		Send(data).
		EndStruct(result)

	if res.StatusCode != 200 || len(errs) > 0 {
		return nil, handleHTTPError(body)
	}

	return body, nil
}

// Patch performs a PATCH request
func (req *Request) Patch(url, data, params string, result interface{}) ([]byte, error) {
	res, body, errs := goreq.
		Patch(url).
		Set("x-sp-gateway", req.gateway).
		Set("x-sp-user-ip", req.ipAddress).
		Set("x-sp-user", req.fingerprint).
		Query(params).
		Send(data).
		EndStruct(result)

	if res.StatusCode != 200 || len(errs) > 0 {
		return nil, handleHTTPError(body)
	}

	return body, nil
}

// Delete performs a DELETE request
func (req *Request) Delete(url string, result interface{}) ([]byte, error) {
	res, body, errs := goreq.
		Delete(url).
		Set("x-sp-gateway", req.gateway).
		Set("x-sp-user-ip", req.ipAddress).
		Set("x-sp-user", req.fingerprint).
		EndStruct(result)

	if res.StatusCode != 200 || len(errs) > 0 {
		return nil, handleHTTPError(body)
	}

	return body, nil
}
