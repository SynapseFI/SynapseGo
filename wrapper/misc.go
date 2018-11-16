package wrapper

/********** METHODS **********/

// DummyTransactions trigger external dummy transactions on deposit or card accounts
func (c *ClientCredentials) DummyTransactions(cred ClientCredentials, userID, nodeID string) User {
	url := _usersURL + "/" + userID + "/nodes" + nodeID + "dummy-tran"
	return handleRequest(cred, "GET", url, nil)
}
