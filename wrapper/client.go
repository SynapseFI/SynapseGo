package wrapper

/********** METHODS **********/

// NewClient creation of client object
func NewClient(gateway, ipAddress, userID string, devMode ...bool) ClientCredentials {
	if len(devMode) == 1 && devMode[0] == true {
		developerMode = true
	}

	return ClientCredentials{
		gateway:   gateway,
		ipAddress: ipAddress,
		userID:    userID,
	}
}

// GetPublicKey returns public key
func (c *ClientCredentials) GetPublicKey(scope []string) User {
	url := _url + "/client?issue_public_key=YES"

	return handleRequest(c, "GET", url, nil)
}
