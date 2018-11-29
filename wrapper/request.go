package wrapper

import (
	"github.com/parnurzeal/gorequest"
)

/********** GLOBAL VARIABLES **********/
const version = "v3.1"

// const _url = "https://api.synapsefi.com/" + version
const _url = "https://uat-api.synapsefi.com/" + version

var goreq = gorequest.New()

// http methods used
const (
	GET   = "GET"
	POST  = "POST"
	PATCH = "PATCH"
)

/********** TYPES **********/

type (
	// Request represents the http request client
	Request struct {
		fingerprint, gateway, ipAddress, authKey string
		headers                                  interface{}
	}
)

/********** METHODS **********/

func (c *Client) newRequest(headers interface{}) *Request {
	return &Request{
		fingerprint: c.Fingerprint,
		gateway:     c.Gateway,
		ipAddress:   c.IP,
		headers:     headers,
	}
}

func (u *User) newRequest(headers interface{}) *Request {
	return &Request{
		fingerprint: u.clientFingerprint,
		gateway:     u.clientGateway,
		ipAddress:   u.clientIP,
		headers:     headers,
	}
}

// Get performs a GET request
func (req *Request) Get(url, params string, result interface{}) ([]byte, *Error) {
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
func (req *Request) Post(url, data, params string, result interface{}) ([]byte, *Error) {
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
func (req *Request) Patch(url, data, params string, result interface{}) ([]byte, *Error) {
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

func request(method, url string, headers map[string]interface{}, params []string, data ...string) []byte {
	var req = gorequest.New()
	req = setMethod(method, url)
	req = setParams(req, params, data)
	req = setHeader(req, headers)

	res, body, errs := req.EndBytes()

	if len(errs) > 0 {
		errorLog(errs)
	}

	if res.StatusCode != 200 {
		handleHTTPError(body)
	}

	return body
}

func setHeader(r *gorequest.SuperAgent, h map[string]interface{}) *gorequest.SuperAgent {

	for k := range h {
		r.Set(k, h[k].(string))
	}

	return r
}

func setMethod(m, u string) *gorequest.SuperAgent {
	switch m {
	case POST:
		return goreq.Post(u)

	case PATCH:
		return goreq.Patch(u)

	default:
		return goreq.Get(u)
	}
}

func setParams(req *gorequest.SuperAgent, params, data []string) *gorequest.SuperAgent {
	var p, d string

	if len(params) > 0 {
		p = params[0]
	}

	if len(data) > 0 {
		d = data[0]
	}

	return req.Query(p).Send(d)
}
