package wrapper

/********** GLOBAL VARIABLES **********/
const authURL = _url + "/oauth"

func (c *Client) Auth(userID string, refreshToken string, bodyParams ...map[string]interface{}) map[string]interface{} {
	url := authURL + "/" + userID
	rt := map[string]interface{}{
		"refresh_token": refreshToken,
		// "phone_number":  bodyParams[0]["phone_number"],
		// "validation_pin": bodyParams[0]["validation_pin"],
	}

	res, body, errs := request.
		Post(url).
		Set("x-sp-gateway", c.gateway).
		Set("x-sp-user-ip", c.ipAddress).
		Set("x-sp-user", c.userID+"|e88f41462eca394f6691da155d0cb73d").
		Send(rt).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return data(body)
}
