package wrapper

/********** GLOBAL VARIABLES **********/
const _institutionsURL = _url + "/institutions"

// GetInstitutions returns all the nodes associated with a user
func GetInstitutions(cred ClientCredentials) ([]byte, error) {
	return handleRequest(cred, "GET", _institutionsURL, nil)
}
