package wrapper

/********** GLOBAL VARIABLES **********/
const _subscriptionsURL = _url + "/subscriptions"

// GetSubscriptions returns all the nodes associated with a user
func GetSubscriptions(cred ClientCredentials) ([]byte, error) {
	return handleRequest(cred, "GET", _subscriptionsURL, nil)
}
