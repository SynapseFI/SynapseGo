package wrapper

/********** GLOBAL VARIABLES **********/

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
