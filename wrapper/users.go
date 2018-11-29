package wrapper

/*********** GLOBAL VARIABLES ***********/
const usersURL = _url + "/users"

/********** TYPE **********/

type (
	// User represents a single user object
	User struct {
		AuthKey           string `json:"oauth_key"`
		clientGateway     string
		clientIP          string
		clientFingerprint string
		FullDehydrate     bool
		UserID            string `json:"_id"`
		RefreshToken      string `json:"refresh_token"`
		Response          interface{}
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

/********** CLIENT METHODS **********/

// GetUsers returns a list of users
func (c *Client) GetUsers(queryParams ...string) *Users {
	var users Users

	h := c.getHeaderInfo("")
	req := c.newRequest(h)

	_, err := req.Get(usersURL, "", &users)

	if err != nil {
		panic(err)
	}

	return &users
}

// GetUser returns a single user
func (c *Client) GetUser(UserID string, fullDehydrate bool, queryParams ...string) *User {
	var user User

	url := usersURL + "/" + UserID

	if fullDehydrate != true {
		url += "?full_dehydrate=yes"
	}

	h := c.getHeaderInfo("")
	req := c.newRequest(h)

	body, err := req.Get(url, "", &user)

	if err != nil {
		panic(err)
	}

	user.FullDehydrate = fullDehydrate
	user.Response = read(body)

	return &user
}

// CreateUser creates a single user and returns the new user data
func (c *Client) CreateUser(data string, queryParams ...string) *User {
	var user User

	h := c.getHeaderInfo("")
	req := c.newRequest(h)

	body, err := req.Post(usersURL, data, "", &user)

	if err != nil {
		panic(err)
	}

	user.Response = read(body)

	return &user
}

/********** USER METHODS **********/

// Update updates a single user and returns the updated user information
func (u *User) Update(data string, queryParams ...string) *User {
	var user User

	url := usersURL + "/" + u.UserID

	h := u.getHeaderInfo("")
	req := u.newRequest(h)

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

	h := u.getHeaderInfo("")
	req := u.newRequest(h)

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

	h := u.getHeaderInfo("")
	req := u.newRequest(h)

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

	h := u.getHeaderInfo("")
	req := u.newRequest(h)

	body, err := req.Patch(url, data, "", &user)

	if err != nil {
		panic(err)
	}

	user.Response = read(body)

	return &user
}

// GetNodes returns all of the nodes associated with a user
func (u *User) GetNodes(queryParams ...string) *Nodes {
	var nodes Nodes

	url := usersURL + "/" + u.UserID + "/nodes"

	h := u.getHeaderInfo("no gateway")
	req := u.newRequest(h)

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

	h := u.getHeaderInfo("no gateway")
	req := u.newRequest(h)

	req.Post(url, data, "", &nodes)

	return &nodes
}
