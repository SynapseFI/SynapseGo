package wrapper

/********** GLOBAL VARIABLES **********/
const nodesURL = _url + "/nodes"

/********** METHODS **********/

// GetAllNodes returns all of the nodes
func (c *Client) GetAllNodes(queryParams ...map[string]interface{}) map[string]interface{} {
	res, body, errs := request.
		Get(nodesURL).
		Query(queryString(queryParams)).
		Set("x-sp-gateway", c.gateway).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return responseMulti(body, "nodes")
}
