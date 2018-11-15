package wrapper

/********** GLOBAL VARIABLES **********/
const _clientTransactionsURL = _url + "/trans"

/********** METHODS  **********/

// GetClientTransactions returns transactions made by all clients
func GetClientTransactions(cred ClientCredentials) ([]byte, error) {
	return handleRequest(cred, "GET", _clientTransactionsURL, nil)
}

// GetUserTransactions returns transactions made by client users
// *CHECK* need OAuth key to make request
func GetUserTransactions(cred ClientCredentials, userID string) ([]byte, error) {
	url := _usersURL + "/" + userID + "/trans"

	return handleRequest(cred, "GET", url, nil)
}
