package wrapper

import (
	"bytes"
	"net/http"
)

/********** METHODS **********/

// OAuth triggers external dummy transactions on deposit or card accounts
func (c *ClientCredentials) OAuth(userID, refreshToken string) *http.Response {
	url := _url + "/oauth/" + userID

	var h Headers
	h["x-sp-gateway"] = c.gateway
	h["x-sp-user-ip"] = c.ipAddress
	h["x-sp-user"] = c.userID

	req := createRequest(h, "POST", url, bytes.NewBufferString(refreshToken))

	return execRequest(req)
}
