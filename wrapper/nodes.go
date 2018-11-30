package wrapper

/********** GLOBAL VARIABLES **********/
const nodesURL = _url + "/nodes"

/********** TYPES **********/

type (
	// Node represents a single node object
	Node struct {
		NodeID        string `json:"_id"`
		UserID        string `json:"user_id"`
		FullDehydrate bool
		user          *User
		Response      interface{}
	}

	// Nodes represents a list of node objects
	Nodes struct {
		Limit     int64  `json:"limit"`
		NodeCount int64  `json:"node_count"`
		Page      int64  `json:"page"`
		PageCount int64  `json:"page_count"`
		Nodes     []Node `json:"nodes"`
	}
)

/********** CLIENT METHODS **********/

// GetAllNodes returns all of the nodes
func (c *Client) GetAllNodes(queryParams ...string) *Nodes {
	var nodes Nodes

	req := c.newRequest()

	_, err := req.Get(nodesURL, queryParams[0], &nodes)

	if err != nil {
		panic(err)
	}

	return &nodes
}
