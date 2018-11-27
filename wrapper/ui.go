package wrapper

import (
	"strings"
)

/********** GLOBAL VARIABLES **********/
const publicKeyURL = _url + "/client?issue_public_key=YES&amp;scope="

/********** METHODS **********/

// GetPublicKey returns a public key as a token representing client credentials
func (c *Client) GetPublicKey(scope ...string) map[string]interface{} {
	var urlParams = publicKeyURL

	for i := 0; i < len(scope); i++ {
		urlParams += scope[i] + ","
	}

	urlParams = strings.TrimSuffix(urlParams, ",")

	h := c.getHeaderInfo("")
	r := request(GET, urlParams, h, nil)

	return response(r, "public_key_obj")
}
