package wrapper

// GetUserNodes returns all the nodes associated with a user
func GetUserNodes(cred ClientCredentials, userID string) ([]byte, error) {
	url := _usersURL + "/" + userID + "/nodes"

	return handleRequest(cred, "GET", url, nil)
}
