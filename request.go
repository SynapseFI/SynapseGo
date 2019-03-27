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
func (req *Request) Get(url string, params []string) ([]byte, error) {
	var p string
	if len(params) > 0 {
		p = params[0]
	}

	res, body, errs := goreq.
		Get(url).
		Set("x-sp-gateway", req.clientID+"|"+req.clientSecret).
		Set("x-sp-user-ip", req.ipAddress).
		Set("x-sp-user", req.authKey+"|"+req.fingerprint).
		Query(p).
		EndBytes()

	if len(errs) > 0 {
		panic(errs)
	}

	if res.StatusCode != 200 {
		return body, handleHTTPError(body)
	}

	log.info("GET", res.Status, url)
	return body, nil
}

// Post performs a POST request
func (req *Request) Post(url, data string, params []string) ([]byte, error) {
	newRequest := goreq.
		Post(url).
		Set("x-sp-gateway", req.clientID+"|"+req.clientSecret).
		Set("x-sp-user-ip", req.ipAddress).
		Set("x-sp-user", req.authKey+"|"+req.fingerprint)

	if len(params) > 0 {
		newRequest.Set("x-sp-idempotency-key", params[0])
	}

	res, body, errs := newRequest.
		Send(data).
		EndBytes()

	if len(errs) > 0 {
		panic(errs)
	}

	if res.StatusCode != 200 && res.StatusCode != 202 {
		return body, handleHTTPError(body)
	}

	log.info("POST", res.Status, url)
	return body, nil
}

// Patch performs a PATCH request
func (req *Request) Patch(url, data string, params []string) ([]byte, error) {
	var p string
	if len(params) > 0 {
		p = params[0]
	}

	res, body, errs := goreq.
		Patch(url).
		Set("x-sp-gateway", req.clientID+"|"+req.clientSecret).
		Set("x-sp-user-ip", req.ipAddress).
		Set("x-sp-user", req.authKey+"|"+req.fingerprint).
		Query(p).
		Send(data).
		EndBytes()

	if len(errs) > 0 {
		panic(errs)
	}

	if res.StatusCode != 200 {
		return body, handleHTTPError(body)
	}

	log.info("PATCH", res.Status, url)
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

	log.info("DELETE", res.Status, url)
	return body, nil
}
