package wrapper

/********** GLOBAL VARIABLES **********/
const nodesURL = _url + "/nodes"

/********** TYPES **********/

type (
	// Node represents a single node object
	Node struct {
		nodeID, userID string
		fullDehydrate  bool
		response       interface{}
	}

	// Nodes represents a list of node objects
	Nodes struct {
		limit, nodeCount, page, pageCount int
		nodes                             []Node
	}
)

/********** METHODS **********/

// GetAllNodes returns all of the nodes
func (c *Client) GetAllNodes(queryParams ...string) map[string]interface{} {
	h := c.getHeaderInfo("gateway")
	r := request(GET, usersURL, h, queryParams)

	return responseMulti(r, "nodes")
}
