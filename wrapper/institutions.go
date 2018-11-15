package wrapper

/********** GLOBAL VARIABLES **********/
const _institutionsURL = _url + "/institutions"

/********** METHODS **********/

// GetInstitutions returns all the institutions
func GetInstitutions(cred ClientCredentials) User {
	return handleRequest(cred, "GET", _institutionsURL, nil)
}
