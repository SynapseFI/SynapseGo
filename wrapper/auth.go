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
