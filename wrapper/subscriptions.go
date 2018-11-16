package wrapper

/********** GLOBAL VARIABLES **********/
const _subscriptionsURL = _url + "/subscriptions"

/********** METHODS **********/

// GetSubscriptions returns all the subscriptions
func (c *ClientCredentials) GetSubscriptions() Users {
	return handleRequestMulti(c, "GET", _subscriptionsURL, "subscriptions ", nil)
}
