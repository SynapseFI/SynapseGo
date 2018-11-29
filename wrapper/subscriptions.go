package wrapper

/********** GLOBAL VARIABLES **********/
const subscriptionsURL = _url + "/subscriptions"

/********** TYPES **********/

type (
	// Subscription represents a single subscription object
	Subscription struct {
		SubscriptionID string `json:"_id"`
		URL            string `json:"url"`
		Response       interface{}
	}

	// Subscriptions represents a list of subscription objects
	Subscriptions struct {
		Limit              int64          `json:"limit"`
		Page               int64          `json:"page"`
		PageCount          int64          `json:"page_count"`
		SubscriptionsCount int64          `json:"subscriptions_count"`
		Subscriptions      []Subscription `json:"subscriptions"`
	}
)

/********** METHODS **********/

// GetSubscriptions returns all of the nodes associated with a user
func (c *Client) GetSubscriptions(queryParams ...string) *Subscriptions {
	var subscriptions Subscriptions

	h := c.getHeaderInfo("")
	req := c.newRequest(h)

	_, err := req.Get(subscriptionsURL, "", &subscriptions)

	if err != nil {
		panic(err)
	}

	return &subscriptions
}

// GetSubscription returns a single subscription
func (c *Client) GetSubscription(subID string, queryParams ...string) *Subscription {
	var subscription Subscription

	url := subscriptionsURL + "/" + subID

	h := c.getHeaderInfo("")
	req := c.newRequest(h)

	body, err := req.Get(url, "", &subscription)

	if err != nil {
		panic(err)
	}

	subscription.Response = read(body)

	return &subscription
}

// CreateSubscription creates a subscription and returns the subscription data
func (c *Client) CreateSubscription(data string, queryParams ...string) *Subscription {
	var subscription Subscription

	h := c.getHeaderInfo("")
	req := c.newRequest(h)

	body, err := req.Post(subscriptionsURL, data, "", &subscription)

	if err != nil {
		panic(err)
	}

	subscription.Response = read(body)

	return &subscription
}

// UpdateSubscription updates an existing subscription
func (c *Client) UpdateSubscription(subID string, data string, queryParams ...string) *Subscription {
	var subscription Subscription

	h := c.getHeaderInfo("")
	req := c.newRequest(h)

	body, err := req.Patch(subscriptionsURL, data, "", &subscription)

	if err != nil {
		panic(err)
	}

	subscription.Response = read(body)

	return &subscription
}
