package wrapper

import (
	"github.com/parnurzeal/gorequest"
)

/********** TYPES **********/

type (
	// Request represents the http request client
	Request struct {
		authKey, clientID, clientSecret, fingerprint, ipAddress string
	}

	// Response is the generic form of all responses from the API
	Response map[string]interface{}
)

/********** GLOBAL VARIABLES **********/
const version = "v3.1"

// const baseUrl = "https://api.synapsefi.com/" + version
const baseURL = "https://uat-api.synapsefi.com/" + version

var goreq = gorequest.New()
var request *Request

var path = map[string]string{
	"auth":          "oauth",
	"client":        "client",
	"institutions":  "institutions",
	"nodes":         "nodes",
	"statements":    "statements",
	"subnets":       "subnets",
	"subscriptions": "subscriptions",
	"transactions":  "trans",
	"users":         "users",
}

var authURL = buildURL(baseURL, path["auth"])
var clientURL = buildURL(baseURL, path["client"])
var institutionsURL = buildURL(baseURL, path["institutions"])
var nodesURL = buildURL(baseURL, path["nodes"])
var subscriptionsURL = buildURL(baseURL, path["subscriptions"])
var transactionsURL = buildURL(baseURL, path["transactions"])
var usersURL = buildURL(baseURL, path["users"])

/********** METHODS **********/

func buildURL(basePath string, uri ...string) string {
	url := basePath

	for i := range uri {
		url += "/" + uri[i]
	}

	return url
}

func (req *Request) updateRequest(clientID, clientSecret, fingerprint, ipAddress string, authKey ...string) *Request {
	var aKey string

	if len(authKey) > 0 {
		aKey = authKey[0]
	}

	return &Request{
		authKey:      aKey,
		clientID:     clientID,
		clientSecret: clientSecret,
		fingerprint:  fingerprint,
		ipAddress:    ipAddress,
	}
}

/********** REQUEST **********/

// Get performs a GET request
func (req *Request) Get(url string, queryParams []string, result interface{}) ([]byte, error) {
	var params string
	if len(queryParams) > 0 {
		params = queryParams[0]
	}

	res, body, errs := goreq.
		Get(url).
		Set("x-sp-gateway", req.clientID+"|"+req.clientSecret).
		Set("x-sp-user-ip", req.ipAddress).
		Set("x-sp-user", req.authKey+"|"+req.fingerprint).
		Query(params).
		EndStruct(&result)

	if len(errs) > 0 {
		panic(errs)
	}

	if res.StatusCode != 200 {
		return nil, handleHTTPError(body)
	}

	return body, nil
}

// Post performs a POST request
func (req *Request) Post(url, data string, queryParams []string, result interface{}) ([]byte, error) {
	var params string
	if len(queryParams) > 0 {
		params = queryParams[0]
	}

	res, body, errs := goreq.
		Post(url).
		Set("x-sp-gateway", req.clientID+"|"+req.clientSecret).
		Set("x-sp-user-ip", req.ipAddress).
		Set("x-sp-user", req.authKey+"|"+req.fingerprint).
		Query(params).
		Send(data).
		EndStruct(&result)

	if len(errs) > 0 {
		panic(errs)
	}

	if res.StatusCode != 200 {
		return nil, handleHTTPError(body)
	}

	return body, nil
}

// Patch performs a PATCH request
func (req *Request) Patch(url, data string, queryParams []string, result interface{}) ([]byte, error) {
	var params string
	if len(queryParams) > 0 {
		params = queryParams[0]
	}

	res, body, errs := goreq.
		Patch(url).
		Set("x-sp-gateway", req.clientID+"|"+req.clientSecret).
		Set("x-sp-user-ip", req.ipAddress).
		Set("x-sp-user", req.authKey+"|"+req.fingerprint).
		Query(params).
		Send(data).
		EndStruct(&result)

	if len(errs) > 0 {
		panic(errs)
	}

	if res.StatusCode != 200 {
		return nil, handleHTTPError(body)
	}

	return body, nil
}

// Delete performs a DELETE request
func (req *Request) Delete(url string, result interface{}) ([]byte, error) {
	res, body, errs := goreq.
		Delete(url).
		Set("x-sp-gateway", req.clientID+"|"+req.clientSecret).
		Set("x-sp-user-ip", req.ipAddress).
		Set("x-sp-user", req.authKey+"|"+req.fingerprint).
		EndStruct(&result)

	if len(errs) > 0 {
		panic(errs)
	}

	if res.StatusCode != 200 {
		return nil, handleHTTPError(body)
	}

	return body, nil
}
