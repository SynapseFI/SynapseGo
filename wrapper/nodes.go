package wrapper

/********** METHODS **********/

// GetUserNodes returns all the nodes associated with a user
func (c *ClientCredentials) GetUserNodes(userID string) Users {
	url := _usersURL + "/" + userID + "/nodes"

	return handleRequestMulti(c, "GET", url, "nodes", nil)
}
