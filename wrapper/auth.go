package wrapper

/********** GLOBAL VARIABLES **********/
const authURL = _url + "/oauth"

/********** TYPES **********/

type (
	// Auth represents an oauth key
	Auth struct {
		AuthKey string `json:"oauth_key"`
	}

	// Refresh represents a refresh token
	Refresh struct {
		Token string `json:"refresh_token"`
	}
)

// Auth returns an oauth key and sets it to the user object
func (u *User) Auth(data string) *Auth {
	var auth Auth

	url := authURL + "/" + u.UserID

	req := u.newRequest()

	_, err := req.Post(url, data, "", &auth)

	if err != nil {
		panic(err)
	}

	u.AuthKey = auth.AuthKey

	return &auth
}
