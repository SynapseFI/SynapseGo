package wrapper

/********** GLOBAL VARIABLES **********/
const _clientTransactionsURL = _url + "/trans"

/********** METHODS  **********/

// GetClientTransactions returns transactions made by all clients
func GetClientTransactions(cred ClientCredentials) ([]byte, error) {
	req := setRequest(cred, "GET", _clientTransactionsURL, nil)

	resp := execRequest(req)

	body := readResponse(resp)

	return formatResponse(cred, body)
}

// GetUserTransactions returns transactions made by client users
// *CHECK* need OAuth key to make request
func GetUserTransactions(cred ClientCredentials, userID string) ([]byte, error) {
	url := _usersURL + "/" + userID + "/trans"

	req := setRequest(cred, "GET", url, nil)

	resp := execRequest(req)

	body := readResponse(resp)

	return formatResponse(cred, body)
}
