package wrapper

/********** GLOBAL VARIABLES **********/
const transURL = _url + "/trans"

/********** METHODS **********/

// GetClientTransactions returns all client transactions
func (c *ClientCredentials) GetClientTransactions() map[string]interface{} {
	res, body, errs := request.
		Get(transURL).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return format(body)
}
