package wrapper

/********** GLOBAL VARIABLES **********/
const subsURL = _url + "/subscriptions"

/********** METHODS **********/

// GetSubscriptions returns all of the nodes associated with a user
func (c *ClientCredentials) GetSubscriptions() map[string]interface{} {
	res, body, errs := request.
		Get(subsURL).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return format(body)
}

// GetSubscription returns a single subscription
func (c *ClientCredentials) GetSubscription(subID string) map[string]interface{} {
	url := subsURL + "/" + subID

	res, body, errs := request.
		Get(url).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return format(body)
}
