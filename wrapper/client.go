package wrapper

/********** METHODS **********/

// GenerateClient creates a client object
func GenerateClient(gateway, ipAddress, userID string, devMode ...bool) *Client {
	if len(devMode) == 1 && devMode[0] == true {
		developerMode = true
	}

	client := &Client{
		gateway:   gateway,
		ipAddress: ipAddress,
		userID:    userID,
	}

	return client
}
