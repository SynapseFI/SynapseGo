package wrapper

/********** GLOBAL VARIABLES **********/
const authURL = _url + "/oauth"

func auth(c *Client, userID string, refreshToken string, bodyParams ...map[string]interface{}) map[string]interface{} {
	url := authURL + "/" + userID
	rt := map[string]interface{}{
		"refresh_token": refreshToken,
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

	return response(body)
}
