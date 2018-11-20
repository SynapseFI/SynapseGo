package wrapper

/********** METHODS **********/

// GenerateClient creates a client object
// func GenerateClient(gateway, ipAddress, userID string, devMode ...bool) *Client {
func GenerateClient(params interface{}) *Client {
	p := params.(map[string]interface{})

	if p["devMode"] == true {
		developerMode = true
	}

	client := &Client{
		gateway:   p["clientID"].(string) + "|" + p["clientSecret"].(string),
		ipAddress: p["ipAddress"].(string),
		userID:    "|" + p["userID"].(string),
	}

	return client
}
