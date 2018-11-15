package wrapper

/********** GLOBAL VARIABLES **********/
const _clientTransactionsURL = _url + "/trans"

/********** METHODS  **********/

// GetClientTransactions returns transactions made by all clients
func GetClientTransactions(cred ClientCredentials) Response {
	return handleRequest(cred, "GET", _clientTransactionsURL, nil)
}

// GetUserTransactions returns transactions made by client users
// *CHECK* need OAuth key to make request
func GetUserTransactions(cred ClientCredentials, userID string) Response {
	url := _usersURL + "/" + userID + "/trans"

	return handleRequest(cred, "GET", url, nil)
}
