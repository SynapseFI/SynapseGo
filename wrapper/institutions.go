package wrapper

/********** GLOBAL VARIABLES **********/
const _institutionsURL = _url + "/institutions"

/********** METHODS **********/

// GetInstitutions returns all the institutions
func (c *ClientCredentials) GetInstitutions(cred ClientCredentials) User {
	return handleRequest(cred, "GET", _institutionsURL, nil)
}
