package wrapper

/********** GLOBAL VARIABLES **********/
const subsURL = _url + "/subscriptions"

/********** METHODS **********/

// GetSubscriptions returns all of the nodes associated with a user
func (c *Client) GetSubscriptions(queryParams ...map[string]interface{}) map[string]interface{} {
	res, body, errs := request.
		Get(subsURL).
		Query(queryString(queryParams)).
		Set("x-sp-gateway", c.gateway).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return responseMulti(body, "subscriptions")
}

// GetSubscription returns a single subscription
func (c *Client) GetSubscription(subID string, queryParams ...map[string]interface{}) map[string]interface{} {
	url := subsURL + "/" + subID

	res, body, errs := request.
		Get(url).
		Query(queryString(queryParams)).
		Set("x-sp-gateway", c.gateway).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return responseSingle(read(body), "subscription")
}

// CreateSubscription creates a subscription and returns the subscription data
func (c *Client) CreateSubscription(data string, queryParams ...map[string]interface{}) map[string]interface{} {
	res, body, errs := request.
		Post(subsURL).
		Query(queryString(queryParams)).
		Set("x-sp-gateway", c.gateway).
		Send(data).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return responseSingle(read(body), "subscription")
}

// UpdateSubscription updates an existing subscription
func (c *Client) UpdateSubscription(subID string, data string, queryParams ...map[string]interface{}) map[string]interface{} {

	url := subsURL + "/" + subID

	res, body, errs := request.
		Patch(url).
		Query(queryString(queryParams)).
		Set("x-sp-gateway", c.gateway).
		Send(data).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return responseSingle(read(body), "subscription")
}
