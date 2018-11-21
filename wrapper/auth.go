package wrapper

/********** GLOBAL VARIABLES **********/
const authURL = _url + "/oauth"

// Auth returns an oauth key and sets it to the user object
func (u *User) Auth(c *Client, bodyParams ...map[string]interface{}) map[string]interface{} {
	url := authURL + "/" + u.UserID
	rt := map[string]interface{}{
		"refresh_token": u.RefreshToken,
	}

	res, body, errs := request.
		Post(url).
		Set("x-sp-gateway", c.gateway).
		Set("x-sp-user-ip", c.ipAddress).
		Set("x-sp-user", c.fingerprint).
		Send(rt).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	u.AuthKey = read(body)["oauth_key"].(string)

	return response(body, "oauth_key")
}
