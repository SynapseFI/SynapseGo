package wrapper

/*********** GLOBAL VARIABLES ***********/
const authURL = _url + "/oauth"
const usersURL = _url + "/users"

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

	// User represents a single user object
	User struct {
		AuthKey       string `json:"oauth_key"`
		client        *Client
		FullDehydrate bool
		UserID        string `json:"_id"`
		RefreshToken  string `json:"refresh_token"`
		Response      interface{}
	}

	// Users represents a collection of user objects
	Users struct {
		Limit      int64  `json:"limit"`
		Page       int64  `json:"page"`
		PageCount  int64  `json:"page_count"`
		UsersCount int64  `json:"users_count"`
		Users      []User `json:"users"`
	}
)

/********** METHODS **********/

/********** AUTHENTICATION **********/

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

/********** NODE **********/

// GetNodes returns all of the nodes associated with a user
func (u *User) GetNodes(queryParams ...string) *Nodes {
	var nodes Nodes

	url := usersURL + "/" + u.UserID + "/nodes"

	req := u.newRequest()

	_, err := req.Get(url, "", &nodes)

	if err != nil {
		panic(err)
	}

	return &nodes
}

// CreateDepositAccount creates an deposit account
func (u *User) CreateDepositAccount(data string) *Nodes {
	var nodes Nodes

	url := usersURL + "/" + u.UserID + "/nodes"

	req := u.newRequest()

	_, err := req.Post(url, data, "", &nodes)

	if err != nil {
		panic(err)
	}

	return &nodes
}

/********** USER **********/

func (u *User) newRequest() *Request {
	return &Request{
		fingerprint: u.AuthKey + u.client.Fingerprint,
		gateway:     u.client.Gateway,
		ipAddress:   u.client.IP,
	}
}

// Update updates a single user and returns the updated user information
func (u *User) Update(data string, queryParams ...string) *User {
	var user User

	url := usersURL + "/" + u.UserID

	req := u.newRequest()

	body, err := req.Patch(url, data, "", &user)

	if err != nil {
		panic(err)
	}

	user.Response = read(body)

	return &user
}

// AddNewDocuments adds new documents to a user
func (u *User) AddNewDocuments(data string) *User {
	var user User

	url := usersURL + "/" + u.UserID

	req := u.newRequest()

	body, err := req.Patch(url, data, "", &user)

	if err != nil {
		panic(err)
	}

	user.Response = read(body)

	return &user
}

// UpdateExistingDocument updates existing user documents
func (u *User) UpdateExistingDocument(data string) *User {
	var user User

	url := usersURL + "/" + u.UserID

	req := u.newRequest()

	body, err := req.Patch(url, data, "", &user)

	if err != nil {
		panic(err)
	}

	user.Response = read(body)

	return &user
}

// DeleteExistingDocument updates existing user documents
func (u *User) DeleteExistingDocument(data string) *User {
	var user User

	url := usersURL + "/" + u.UserID

	req := u.newRequest()

	body, err := req.Patch(url, data, "", &user)

	if err != nil {
		panic(err)
	}

	user.Response = read(body)

	return &user
}
