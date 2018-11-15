package wrapper

/********** GLOBAL VARIABLES **********/
const _subscriptionsURL = _url + "/subscriptions"

/********** METHODS **********/

// GetSubscriptions returns all the subscriptions
func GetSubscriptions(cred ClientCredentials) ([]byte, error) {
	return handleRequest(cred, "GET", _subscriptionsURL, nil)
}
