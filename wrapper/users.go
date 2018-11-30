package wrapper

/*********** GLOBAL VARIABLES ***********/

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

	url := buildURL(authURL, u.UserID)

	_, err := request.Post(url, data, "", &auth)

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

	url := buildURL(usersURL, u.UserID, path["nodes"])

	_, err := request.Get(url, "", &nodes)

	if err != nil {
		panic(err)
	}

	return &nodes
}

// CreateDepositAccount creates an deposit account
func (u *User) CreateDepositAccount(data string) *Nodes {
	var nodes Nodes

	url := buildURL(usersURL, u.UserID, path["nodes"])

	_, err := request.Post(url, data, "", &nodes)

	if err != nil {
		panic(err)
	}

	return &nodes
}

/********** USER **********/

// Update updates a single user and returns the updated user information
func (u *User) Update(data string, queryParams ...string) *User {
	var user User

	url := buildURL(usersURL, u.UserID)

	body, err := request.Patch(url, data, "", &user)

	if err != nil {
		panic(err)
	}

	user.Response = read(body)

	return &user
}

// AddNewDocuments adds new documents to a user
func (u *User) AddNewDocuments(data string) *User {
	var user User

	url := buildURL(usersURL, u.UserID)

	body, err := request.Patch(url, data, "", &user)

	if err != nil {
		panic(err)
	}

	user.Response = read(body)

	return &user
}

// UpdateExistingDocument updates existing user documents
func (u *User) UpdateExistingDocument(data string) *User {
	var user User

	url := buildURL(usersURL, u.UserID)

	body, err := request.Patch(url, data, "", &user)

	if err != nil {
		panic(err)
	}

	user.Response = read(body)

	return &user
}

// DeleteExistingDocument updates existing user documents
func (u *User) DeleteExistingDocument(data string) *User {
	var user User

	url := buildURL(usersURL, u.UserID)

	body, err := request.Patch(url, data, "", &user)

	if err != nil {
		panic(err)
	}

	user.Response = read(body)

	return &user
}
