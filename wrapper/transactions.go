package wrapper

/********** GLOBAL VARIABLES **********/
const _clientTransactionsURL = _url + "/trans"

/********** METHODS  **********/

// GetClientTransactions returns transactions made by all clients
func (c *ClientCredentials) GetClientTransactions() Users {
	return handleRequestMulti(c, "GET", _clientTransactionsURL, "trans", nil)
}

// GetUserTransactions returns transactions made by client users
// *CHECK* need OAuth key to make request
func (c *ClientCredentials) GetUserTransactions(userID string) Users {
	url := _usersURL + "/" + userID + "/trans"

	return handleRequestMulti(c, "GET", url, "trans", nil)
}
