package wrapper

/********** GLOBAL VARIABLES **********/
const authURL = _url + "/oauth"

// Auth returns an oauth key and sets it to the user object
func (u *User) Auth(data string) map[string]interface{} {
	url := authURL + "/" + u.UserID

	h := u.getHeaderInfo("")
	r := request(POST, url, h, nil, data)

	u.AuthKey = read(r)["oauth_key"].(string)

	return response(r, "oauth_key")
}
