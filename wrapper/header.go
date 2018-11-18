package wrapper

func header(c *ClientCredentials, setting string) {
	switch setting {
	default:
		request.
			Set("x-sp-gateway", c.gateway).
			Set("x-sp-user-ip", c.ipAddress).
			Set("x-sp-user", c.userID)
	}
}
