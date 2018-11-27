package wrapper

/********** GLOBAL VARIABLES **********/
const nodesURL = _url + "/nodes"

/********** METHODS **********/

// GetAllNodes returns all of the nodes
func (c *Client) GetAllNodes(queryParams ...string) map[string]interface{} {
	h := c.getHeaderInfo("gateway")
	r := request(GET, usersURL, h, queryParams)

	return responseMulti(r, "nodes")
}
