package wrapper

/********** GLOBAL VARIABLES **********/
const instiURL = _url + "/institutions"

/********** METHODS **********/

// GetInstitutions returns all of the nodes associated with a user
func (c *Client) GetInstitutions() map[string]interface{} {
	r := request(GET, instiURL, nil, nil)

	return response(r)
}
