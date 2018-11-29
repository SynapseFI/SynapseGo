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
func (c *Client) GetPublicKey(scope ...string) (*PublicKey, *Error) {
	var urlParams = publicKeyURL

	for i := 0; i < len(scope); i++ {
		urlParams += scope[i] + ","
	}

	urlParams = strings.TrimSuffix(urlParams, ",")

	h := c.getHeaderInfo("")
	req := newRequest(c, h)

	var pk PublicKey
	_, err := req.Get(urlParams, "", &pk)

	if err != nil {
		return nil, err
	}

	return &pk, nil
}
