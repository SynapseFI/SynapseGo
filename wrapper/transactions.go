package wrapper

/********** GLOBAL VARIABLES **********/

/********** TYPES **********/

type (
	// Transaction represents a single transaction object
	Transaction struct {
		TransactionID string `json:"_id"`
		node          *Node
		Response      interface{}
	}

	// Transactions represents a list of transaction objects
	Transactions struct {
		Limit            int64         `json:"limit"`
		Page             int64         `json:"page"`
		PageCount        int64         `json:"page_count"`
		TransactionCount int64         `json:"trans_count"`
		Transactions     []Transaction `json:"trans"`
	}
)
