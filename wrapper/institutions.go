package wrapper

/********** GLOBAL VARIABLES **********/
const _institutionsURL = _url + "/institutions"

/********** METHODS **********/

// GetInstitutions returns all the institutions
func (c *ClientCredentials) GetInstitutions() User {
	return handleRequest(c, "GET", _institutionsURL, nil)
}
