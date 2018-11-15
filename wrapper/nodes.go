package wrapper

/********** METHODS **********/

// GetUserNodes returns all the nodes associated with a user
func GetUserNodes(cred ClientCredentials, userID string) User {
	url := _usersURL + "/" + userID + "/nodes"

	return handleRequest(cred, "GET", url, nil)
}
