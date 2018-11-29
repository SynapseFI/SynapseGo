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
		Limit     int64  `json:"limit"`
		NodeCount int64  `json:"node_count"`
		Page      int64  `json:"page"`
		PageCount int64  `json:"page_count"`
		Nodes     []Node `json:"nodes"`
	}
)

/********** METHODS **********/

// GetAllNodes returns all of the nodes
func (c *Client) GetAllNodes(queryParams ...string) (*Nodes, *Error) {
	var nodes Nodes

	h := c.getHeaderInfo("")
	req := newRequest(c, h)

	_, err := req.Get(nodesURL, queryParams[0], &nodes)

	if err != nil {
		return nil, err
	}

	return &nodes, nil
}
