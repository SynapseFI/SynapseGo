package wrapper

import (
	"github.com/parnurzeal/gorequest"
)

type (
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

/********** TYPES **********/

type (
	// Request represents the http request client
	Request struct {
		authKey, clientID, clientSecret, fingerprint, gateway, ipAddress string
	}
)

/********** METHODS **********/

func (req *Request) authenticate(userID, refreshToken string, bodyData ...string) *Auth {
	var auth Auth
	var data string

	url := buildURL(authURL, userID)

	if len(bodyData) > 0 {
		data = bodyData[0]
	}

	rt := `{"refresh_token":"` + refreshToken + `"}`
	// try and retrieve an oauth token
	_, err := req.Post(url, rt, data, &auth)

	// if retrieval fails, check if it was because of an invalid refresh token
	if _, ok := err.(*IncorrectValues); ok {
		var refresh Refresh
		url := buildURL(usersURL, userID)
		// try and get a new refresh token
		_, err := req.Get(url, "", &refresh)

		if err != nil {
			panic(err)
		}

		return req.authenticate(userID, refresh.Token)
	}

	return &auth
}

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

	request = &Request{
		authKey:      aKey,
		clientID:     clientID,
		clientSecret: clientSecret,
		fingerprint:  aKey + "|" + fingerprint,
		gateway:      clientID + "|" + clientSecret,
		ipAddress:    ipAddress,
	}

	return request
}

/********** REQUEST **********/

// Get performs a GET request
func (req *Request) Get(url, params string, result interface{}, userData ...*User) ([]byte, error) {
	req = req.updateRequest(req.clientID, req.clientSecret, req.fingerprint, req.ipAddress, req.authKey)

	res, body, errs := goreq.
		Get(url).
		Set("x-sp-gateway", req.gateway).
		Set("x-sp-user-ip", req.ipAddress).
		Set("x-sp-user", req.fingerprint).
		Query(params).
		EndStruct(&result)

	if len(errs) > 0 {
		panic(errs)
	}

	if res.StatusCode != 200 && res.StatusCode != 202 {
		err := handleHTTPError(body)

		// check if err is of type IncorrectUserCredentials
		if _, ok := err.(*IncorrectUserCredentials); ok {
			user := userData[0]
			auth := req.authenticate(user.UserID, user.RefreshToken)
			req.updateRequest(req.clientID, req.clientSecret, req.fingerprint, req.ipAddress, auth.Key)

			return req.Get(url, params, &result)
		}

		return nil, err
	}

	return body, nil
}

// Post performs a POST request
func (req *Request) Post(url, data, params string, result interface{}, userData ...*User) ([]byte, error) {
	req = req.updateRequest(req.clientID, req.clientSecret, req.fingerprint, req.ipAddress, req.authKey)

	res, body, errs := goreq.
		Post(url).
		Set("x-sp-gateway", req.gateway).
		Set("x-sp-user-ip", req.ipAddress).
		Set("x-sp-user", req.fingerprint).
		Query(params).
		Send(data).
		EndStruct(&result)

	if len(errs) > 0 {
		panic(errs)
	}

	if res.StatusCode != 200 && res.StatusCode != 202 {
		err := handleHTTPError(body)

		// check if err is of type IncorrectUserCredentials
		if _, ok := err.(*IncorrectUserCredentials); ok {
			user := userData[0]
			auth := req.authenticate(user.UserID, user.RefreshToken)
			req.updateRequest(req.clientID, req.clientSecret, req.fingerprint, req.ipAddress, auth.Key)

			return req.Post(url, data, params, &result)
		}

		return nil, err
	}

	return body, nil
}

// Patch performs a PATCH request
func (req *Request) Patch(url, data, params string, result interface{}, userData ...*User) ([]byte, error) {
	req = req.updateRequest(req.clientID, req.clientSecret, req.fingerprint, req.ipAddress, req.authKey)

	res, body, errs := goreq.
		Patch(url).
		Set("x-sp-gateway", req.gateway).
		Set("x-sp-user-ip", req.ipAddress).
		Set("x-sp-user", req.fingerprint).
		Query(params).
		Send(data).
		EndStruct(&result)

	if len(errs) > 0 {
		panic(errs)
	}

	if res.StatusCode != 200 && res.StatusCode != 202 {
		err := handleHTTPError(body)

		// check if err is of type IncorrectUserCredentials
		if _, ok := err.(*IncorrectUserCredentials); ok {
			user := userData[0]
			auth := req.authenticate(user.UserID, user.RefreshToken)
			req.updateRequest(req.clientID, req.clientSecret, req.fingerprint, req.ipAddress, auth.Key)

			return req.Patch(url, data, params, &result)
		}

		return nil, err
	}

	return body, nil
}

// Delete performs a DELETE request
func (req *Request) Delete(url string, result interface{}, userData ...*User) ([]byte, error) {
	req = req.updateRequest(req.clientID, req.clientSecret, req.fingerprint, req.ipAddress, req.authKey)

	res, body, errs := goreq.
		Delete(url).
		Set("x-sp-gateway", req.gateway).
		Set("x-sp-user-ip", req.ipAddress).
		Set("x-sp-user", req.fingerprint).
		EndStruct(&result)

	if len(errs) > 0 {
		panic(errs)
	}

	if res.StatusCode != 200 && res.StatusCode != 202 {
		err := handleHTTPError(body)

		// check if err is of type IncorrectUserCredentials
		if _, ok := err.(*IncorrectUserCredentials); ok {
			user := userData[0]
			auth := req.authenticate(user.UserID, user.RefreshToken)
			req.updateRequest(req.clientID, req.clientSecret, req.fingerprint, req.ipAddress, auth.Key)

			return req.Delete(url, &result)
		}
	}

	return body, nil
}
