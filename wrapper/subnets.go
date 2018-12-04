package wrapper

/********** GLOBAL VARIABLES **********/

/********** TYPES **********/

type (
	// Subnet represents a single Subnet object
	Subnet struct {
		SubnetID string `json:"_id"`
		NodeID   string `json:"node_id"`
		UserID   string `json:"user_id"`
		Response interface{}
	}

	// Subnets represents a list of transaction objects
	Subnets struct {
		Limit       int64    `json:"limit"`
		Page        int64    `json:"page"`
		PageCount   int64    `json:"page_count"`
		SubnetCount int64    `json:"subnets_count"`
		Subnets     []Subnet `json:"subnets"`
	}
)
