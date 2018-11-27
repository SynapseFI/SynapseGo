package wrapper

type header interface {
	getHeaderInfo() map[string]interface{}
}

func (c *Client) getHeaderInfo(setting string) map[string]interface{} {
	var info = make(map[string]interface{})

	switch setting {
	default:
		info["x-sp-gateway"] = c.gateway
		info["x-sp-user-ip"] = c.ipAddress
		info["x-sp-user"] = c.fingerprint
	}

	return info
}

func (u *User) getHeaderInfo(setting string) map[string]interface{} {
	var info = make(map[string]interface{})

	switch setting {
	case "no gateway":
		info["x-sp-user-ip"] = u.clientIP
		info["x-sp-user"] = u.AuthKey + "|" + u.clientFingerprint

	default:
		info["x-sp-gateway"] = u.clientGateway
		info["x-sp-user-ip"] = u.clientIP
		info["x-sp-user"] = u.AuthKey + "|" + u.clientFingerprint
	}

	return info
}
