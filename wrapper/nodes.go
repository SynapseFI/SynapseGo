package wrapper

/********** GLOBAL VARIABLES **********/

/********** METHODS **********/

// GetUserNodes returns all of the nodes associated with a user
func (c *ClientCredentials) GetUserNodes(userID string) map[string]interface{} {
	url := usersURL + "/" + userID + "/nodes"

	header(c, authUserSetting)

	res, body, errs := request.
		Get(url).
		Set("x-sp-user-ip", c.ipAddress).
		Set("x-sp-user", c.userID).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return multiData(body, "nodes")
}
