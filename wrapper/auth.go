package wrapper

/********** GLOBAL VARIABLES **********/
const authURL = _url + "/oauth"

// Auth returns an oauth key and sets it to the user object
func (u *User) Auth(data string) map[string]interface{} {
	url := authURL + "/" + u.UserID

	res, body, errs := request.
		Post(url).
		Set("x-sp-gateway", u.clientGateway).
		Set("x-sp-user-ip", u.clientIP).
		Set("x-sp-user", u.clientFingerprint).
		Send(data).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	u.AuthKey = read(body)["oauth_key"].(string)

	return response(body, "oauth_key")
}
