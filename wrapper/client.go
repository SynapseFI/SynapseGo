package wrapper

/********** METHODS **********/

// NewClient creation of client object
func NewClient(gateway, ipAddress, userID string, devMode ...bool) *Client {
	if len(devMode) == 1 && devMode[0] == true {
		developerMode = true
	}

	credentials := &Client{
		gateway:   gateway,
		ipAddress: ipAddress,
		userID:    userID,
	}

	return credentials
}
