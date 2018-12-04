package wrapper

/********** GLOBAL VARIABLES **********/

/********** TYPES **********/

type (
	// Node represents a single node object
	Node struct {
		NodeID        string `json:"_id"`
		UserID        string `json:"user_id"`
		FullDehydrate bool
		user          *User
		Response      interface{}
	}

	// Nodes represents a list of node objects
	Nodes struct {
		Limit     int64  `json:"limit"`
		MFA       MFA    `json:"mfa"`
		NodeCount int64  `json:"node_count"`
		Nodes     []Node `json:"nodes"`
		Page      int64  `json:"page"`
		PageCount int64  `json:"page_count"`
	}
)
