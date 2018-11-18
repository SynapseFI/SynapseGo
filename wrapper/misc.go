package wrapper

/********** GLOBAL VARIABLES **********/

/********** METHODS **********/

// GetPublicKey returns all of the nodes associated with a user
func (c *ClientCredentials) GetPublicKey(scope ...string) map[string]interface{} {
	url := _url + "/client?issue_public_key=YES"

	res, body, errs := request.
		Get(url).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return format(body)
}
