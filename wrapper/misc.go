package wrapper

/********** GLOBAL VARIABLES **********/
const instiURL = _url + "/institutions"

/********** METHODS **********/

// GetInstitutions returns all of the nodes associated with a user
func (c *Client) GetInstitutions() map[string]interface{} {
	res, body, errs := request.
		Get(instiURL).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return response(body)
}
