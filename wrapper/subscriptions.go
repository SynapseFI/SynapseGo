package wrapper

/********** GLOBAL VARIABLES **********/
const subsURL = _url + "/subscriptions"

/********** METHODS **********/

// GetSubscriptions returns all of the nodes associated with a user
func (c *Client) GetSubscriptions(queryParams ...string) map[string]interface{} {
	h := c.getHeaderInfo("gateway")
	r := request(GET, subsURL, h, queryParams)

	return responseMulti(r, "subscriptions")
}

// GetSubscription returns a single subscription
func (c *Client) GetSubscription(subID string, queryParams ...string) map[string]interface{} {
	url := subsURL + "/" + subID

	h := c.getHeaderInfo("gateway")
	r := request(GET, url, h, queryParams)

	return responseSingle(r, "subscription")
}

// CreateSubscription creates a subscription and returns the subscription data
func (c *Client) CreateSubscription(data string, queryParams ...string) map[string]interface{} {
	h := c.getHeaderInfo("gateway")
	r := request(POST, usersURL, h, queryParams, data)

	return responseSingle(r, "subscription")
}

// UpdateSubscription updates an existing subscription
func (c *Client) UpdateSubscription(subID string, data string, queryParams ...string) map[string]interface{} {
	url := subsURL + "/" + subID

	h := c.getHeaderInfo("gateway")
	r := request(PATCH, url, h, queryParams, data)

	return responseSingle(r, "subscription")
}
