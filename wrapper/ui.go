package wrapper

import (
	"strings"
)

/********** GLOBAL VARIABLES **********/
const publicKeyURL = _url + "/client?issue_public_key=YES&amp;scope="

/********** METHODS **********/

// GetPublicKey returns a public key as a token representing client credentials
func (c *ClientCredentials) GetPublicKey(scope ...string) map[string]interface{} {
	var urlParams = publicKeyURL

	for i := 0; i < len(scope); i++ {
		urlParams += scope[i] + ","
	}

	urlParams = strings.TrimSuffix(urlParams, ",")

	res, body, errs := request.
		Get(urlParams).
		Set("x-sp-gateway", c.gateway).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return data(body)
}
