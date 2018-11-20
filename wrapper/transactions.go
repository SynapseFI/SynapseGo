package wrapper

/********** GLOBAL VARIABLES **********/
const transURL = _url + "/trans"

/********** METHODS **********/

// GetClientTransactions returns all client transactions
func (c *Client) GetClientTransactions(queryParams ...map[string]interface{}) map[string]interface{} {

	res, body, errs := request.
		Get(transURL).
		Query(queryString(queryParams)).
		Set("x-sp-gateway", c.gateway).
		Set("x-sp-user-ip", c.ipAddress).
		Set("x-sp-user", c.fingerprint).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return responseMulti(body, "transactions")
}
