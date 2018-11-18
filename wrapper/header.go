package wrapper

import "net/http"

func createHeader(cred *ClientCredentials, r *http.Request, name string) Header {
	header := make(Header)
	header["content-type"] = "application/json"

	switch name {
	default:
		header["x-sp-gateway"] = cred.gateway
		header["x-sp-user-ip"] = cred.ipAddress
		header["x-sp-user"] = cred.userID
	}

	for k := range header {
		r.Header.Set(k, header[k])
	}

	return header
}

func makeHeader(h Header, r *http.Request) {
	for k := range h {
		r.Header.Set(k, h[k])
	}

	r.Header.Set("content-type", "application/json")
}
