package wrapper

/********** METHODS **********/

// GenerateClient creates a client object
// func GenerateClient(gateway, ipAddress, userID string, devMode ...bool) *Client {
func GenerateClient(params interface{}) *Client {
	p := params.(map[string]interface{})

	if p["devMode"] == true {
		developerMode = true
	}

	return &Client{
		gateway:     p["clientID"].(string) + "|" + p["clientSecret"].(string),
		ipAddress:   p["ipAddress"].(string),
		fingerprint: "|" + p["fingerprint"].(string),
	}
}

// Info returns client credentials
func (c *Client) Info() map[string]interface{} {
	return map[string]interface{}{
		"gateway":     c.gateway,
		"ipAddress":   c.ipAddress,
		"fingerprint": c.fingerprint,
	}
}
