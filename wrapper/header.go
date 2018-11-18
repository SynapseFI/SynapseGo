package wrapper

const authUserSetting = "auth_user"
const userSetting = "user"
const gatewaySetting = "gateway"

func header(c *ClientCredentials, setting string) {
	switch setting {
	case authUserSetting:
		request.
			Set("x-sp-user-ip", c.ipAddress).
			Set("x-sp-user", "*CHECK* needs OAUTH")
	case userSetting:
		request.
			Set("x-sp-user-ip", c.ipAddress).
			Set("x-sp-user", c.userID)
	case gatewaySetting:
		request.
			Set("x-sp-gateway", c.gateway)
	default:
		request.
			Set("x-sp-gateway", c.gateway).
			Set("x-sp-user-ip", c.ipAddress).
			Set("x-sp-user", c.userID)
	}
}
