package wrapper

/********** METHODS **********/

// GetUserNodes returns all the nodes associated with a user
func GetUserNodes(cred ClientCredentials, userID string) Users {
	url := _usersURL + "/" + userID + "/nodes"

	return handleRequestMulti(cred, "GET", url, "nodes", nil)
}
