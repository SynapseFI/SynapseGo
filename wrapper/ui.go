package wrapper

// GetPublicKey returns public key
func (c *ClientCredentials) GetPublicKey(scope []string) User {
	url := _url + "/client?issue_public_key=YES"

	return handleRequest(c, "GET", url, nil)
}
