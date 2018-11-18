package wrapper

/********** GLOBAL VARIABLES **********/

/********** METHODS **********/

// GetUserNodes returns all of the nodes associated with a user
func (c *ClientCredentials) GetUserNodes(userID string) map[string]interface{} {
	url := usersURL + "/" + userID + "/nodes"

	res, body, errs := request.
		Get(url).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return format(body)
}
