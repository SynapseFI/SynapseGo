package wrapper

/********** GLOBAL VARIABLES **********/
const transURL = _url + "/trans"

/********** METHODS **********/

// GetClientTransactions returns all client transactions
func (c *ClientCredentials) GetClientTransactions() map[string]interface{} {
	header(c, "")

	res, body, errs := request.
		Get(transURL).
		Set("x-sp-gateway", c.gateway).
		Set("x-sp-user-ip", c.ipAddress).
		Set("x-sp-user", c.userID).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return multiData(body, "transactions")
}
