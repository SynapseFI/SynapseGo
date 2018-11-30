package wrapper

import (
	"strings"
)

/********** GLOBAL VARIABLES **********/
const publicKeyURL = _url + "/client?issue_public_key=YES&amp;scope="

/********** TYPES **********/

type (
	// PublicKey represents the structure of a public key object
	PublicKey struct {
		Response interface{} `json:"public_key_obj"`
	}
)

/********** METHODS **********/

// GetPublicKey returns a public key as a token representing client credentials
func (c *Client) GetPublicKey(scope ...string) *PublicKey {
	var publicKey PublicKey
	var urlParams = publicKeyURL

	for i := 0; i < len(scope); i++ {
		urlParams += scope[i] + ","
	}

	urlParams = strings.TrimSuffix(urlParams, ",")

	req := c.newRequest()

	_, err := req.Get(urlParams, "", &publicKey)

	if err != nil {
		panic(err)
	}

	return &publicKey
}
