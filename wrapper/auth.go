package wrapper

/********** GLOBAL VARIABLES **********/
const authURL = _url + "/oauth"

/********** TYPES **********/

type (
	// Auth represents an oauth key
	Auth struct {
		Key string `json:"oauth_key"`
	}
)

// Auth returns an oauth key and sets it to the user object
func (u *User) Auth(data string) (*Auth, *Error) {
	var auth Auth

	url := authURL + "/" + u.UserID

	h := u.getHeaderInfo("")
	req := u.newRequest(h)

	_, err := req.Post(url, data, "", &auth)

	if err != nil {
		return nil, err
	}

	u.AuthKey = auth.Key

	return &auth, nil
}
