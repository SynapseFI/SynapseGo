package wrapper

/********** GLOBAL VARIABLES **********/
const transURL = _url + "/trans"

/********** METHODS **********/

// GetClientTransactions returns all client transactions
func (c *Client) GetClientTransactions(queryParams ...string) map[string]interface{} {

	h := c.getHeaderInfo("")
	r := apiRequest(GET, transURL, h, queryParams)

	return responseMulti(r, "transactions")
}
