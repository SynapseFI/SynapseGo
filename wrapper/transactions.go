package wrapper

/********** GLOBAL VARIABLES **********/
const transactionsURL = _url + "/trans"

/********** TYPES **********/

type (
	// Transaction represents a single transaction object
	Transaction struct {
		TransactionID string
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

/********** METHODS **********/

// GetClientTransactions returns all client transactions
func (c *Client) GetClientTransactions(queryParams ...string) *Transactions {
	var transactions Transactions

	h := c.getHeaderInfo("")
	req := c.newRequest(h)

	_, err := req.Get(transactionsURL, queryParams[0], &transactions)

	if err != nil {
		panic(err)
	}

	return &transactions
}
