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
func (c *Client) GetSubscriptions(queryParams ...string) (*Subscriptions, *Error) {
	var subscriptions Subscriptions

	h := c.getHeaderInfo("")
	req := newRequest(c, h)

	_, err := req.Get(subscriptionsURL, "", &subscriptions)

	if err != nil {
		return nil, err
	}

	return &subscriptions, nil
}

// GetSubscription returns a single subscription
func (c *Client) GetSubscription(subID string, queryParams ...string) (*Subscription, *Error) {
	var subscription Subscription

	url := subscriptionsURL + "/" + subID

	h := c.getHeaderInfo("")
	req := newRequest(c, h)

	body, err := req.Get(url, "", &subscription)

	if err != nil {
		return nil, err
	}

	subscription.Response = read(body)

	return &subscription, nil
}

// CreateSubscription creates a subscription and returns the subscription data
func (c *Client) CreateSubscription(data string, queryParams ...string) (*Subscription, *Error) {
	var subscription Subscription

	h := c.getHeaderInfo("")
	req := newRequest(c, h)

	body, err := req.Post(subscriptionsURL, data, "", &subscription)

	if err != nil {
		return nil, err
	}

	subscription.Response = read(body)

	return &subscription, nil
}

// UpdateSubscription updates an existing subscription
func (c *Client) UpdateSubscription(subID string, data string, queryParams ...string) (*Subscription, *Error) {
	var subscription Subscription

	h := c.getHeaderInfo("")
	req := newRequest(c, h)

	body, err := req.Patch(subscriptionsURL, data, "", &subscription)

	if err != nil {
		return nil, err
	}

	subscription.Response = read(body)

	return &subscription, nil
}
