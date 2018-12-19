// +build !mock

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
		EndBytes()

	if len(errs) > 0 {
		panic(errs)
	}

	if res.StatusCode != 200 {
		return body, handleHTTPError(body)
	}

	return body, nil
}

// Post performs a POST request
func (req *Request) Post(url, data string, queryParams []string) ([]byte, error) {
	var params string

	newRequest := goreq.
		Post(url).
		Set("x-sp-gateway", req.clientID+"|"+req.clientSecret).
		Set("x-sp-user-ip", req.ipAddress).
		Set("x-sp-user", req.authKey+"|"+req.fingerprint)

	if len(queryParams) > 0 {
		params = queryParams[0]

		if len(queryParams) > 1 {
			newRequest.Set("x-sp-idempotency-key", queryParams[1])
		}
	}

	res, body, errs := newRequest.
		Query(params).
		Send(data).
		EndBytes()

	if len(errs) > 0 {
		panic(errs)
	}

	if res.StatusCode != 200 {
		return body, handleHTTPError(body)
	}

	return body, nil
}

// Patch performs a PATCH request
func (req *Request) Patch(url, data string, queryParams []string) ([]byte, error) {
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
		EndBytes()

	if len(errs) > 0 {
		panic(errs)
	}

	if res.StatusCode != 200 {
		return body, handleHTTPError(body)
	}

	return body, nil
}

// Delete performs a DELETE request
func (req *Request) Delete(url string) ([]byte, error) {
	res, body, errs := goreq.
		Delete(url).
		Set("x-sp-gateway", req.clientID+"|"+req.clientSecret).
		Set("x-sp-user-ip", req.ipAddress).
		Set("x-sp-user", req.authKey+"|"+req.fingerprint).
		EndBytes()

	if len(errs) > 0 {
		panic(errs)
	}

	if res.StatusCode != 200 {
		return body, handleHTTPError(body)
	}

	return body, nil
}
