package wrapper

/********** GLOBAL VARIABLES **********/
const instiURL = _url + "/institutions"

/********** METHODS **********/

// GetInstitutions returns all of the nodes associated with a user
func (c *ClientCredentials) GetInstitutions() map[string]interface{} {
	header(c, authUserSetting)

	res, body, errs := request.
		Get(instiURL).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return response(body, "institution")
}
