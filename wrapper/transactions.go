package wrapper

/********** GLOBAL VARIABLES **********/
const transURL = _url + "/trans"

/********** TYPES **********/

type (
	// Transaction represents a single transaction object
	Transaction struct {
		transactionID string
		response      interface{}
	}

	// Transactions represents a list of transaction objects
	Transactions struct {
		limit, transactionCount, page, pageCount int
		transactions                             []Transaction
	}
)

/********** METHODS **********/

// GetClientTransactions returns all client transactions
func (c *Client) GetClientTransactions(queryParams ...string) map[string]interface{} {

	h := c.getHeaderInfo("")
	r := request(GET, transURL, h, queryParams)

	return responseMulti(r, "transactions")
}
