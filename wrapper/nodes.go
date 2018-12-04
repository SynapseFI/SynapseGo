package wrapper

/********** GLOBAL VARIABLES **********/

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
		MFA       MFA    `json:"mfa"`
		NodeCount int64  `json:"node_count"`
		Nodes     []Node `json:"nodes"`
		Page      int64  `json:"page"`
		PageCount int64  `json:"page_count"`
	}
)

/********** METHODS **********/

/********** NODE **********/

/********** OTHER **********/

// DummyTransactions triggers external dummy transactions on deposit or card accounts
func (n *Node) DummyTransactions(credit bool) *Response {
	var response Response

	url := buildURL(usersURL, n.UserID, path["nodes"], n.NodeID) + "/dummy-tran"

	if credit == true {
		url += "?is_credit=YES"
	}

	_, err := request.Get(url, "", &response)

	if err != nil {
		panic(err)
	}

	return &response
}

// ResetDebitCard resets the debit card number, card cvv, and expiration date
func (n *Node) ResetDebitCard() *Response {
	var response Response

	url := buildURL(usersURL, n.UserID, path["nodes"], n.NodeID) + "?reset=YES"

	_, err := request.Patch(url, "", "", &response)

	if err != nil {
		panic(err)
	}

	return &response
}

// ShipDebitCard ships a physical debit card out to the user
func (n *Node) ShipDebitCard(data string) *Response {
	var response Response

	url := buildURL(usersURL, n.UserID, path["nodes"], n.NodeID) + "?ship=YES"

	_, err := request.Patch(url, data, "", &response)

	if err != nil {
		panic(err)
	}

	return &response
}

/********** TRANSACTION **********/
