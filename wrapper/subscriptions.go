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
