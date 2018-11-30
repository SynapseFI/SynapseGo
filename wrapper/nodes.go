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
		NodeCount int64  `json:"node_count"`
		Page      int64  `json:"page"`
		PageCount int64  `json:"page_count"`
		Nodes     []Node `json:"nodes"`
	}
)

/********** METHODS **********/

/********** NODE **********/

// ShipDebitCard ships a physical debit card out to the user
func (n *Node) ShipDebitCard(data string) *Node {
	var node Node

	url := usersURL + "/" + n.UserID + "/nodes/" + n.NodeID + "?ship=YES"

	_, err := request.Patch(url, data, "", &node)

	if err != nil {
		panic(err)
	}

	return &node
}

// ResetDebitCard resets the debit card number, card cvv, and expiration date
func (n *Node) ResetDebitCard() *Node {
	var node Node

	url := usersURL + "/" + n.UserID + "/nodes/" + n.NodeID + "?reset=YES"

	_, err := request.Patch(url, "", "", &node)

	if err != nil {
		panic(err)
	}

	return &node
}

/********** OTHER **********/

// DummyTransactions triggers external dummy transactions on deposit or card accounts
func (n *Node) DummyTransactions(credit bool) map[string]interface{} {
	var response map[string]interface{}
	url := usersURL + "/" + n.UserID + "/nodes/" + n.NodeID + "/dummy-tran"

	if credit == true {
		url += "?is_credit=YES"
	}

	_, err := request.Get(url, "", response)

	if err != nil {
		panic(err)
	}

	return response
}

/********** TRANSACTION **********/

// GetTransaction returns a specific transaction associated with a node
func (n *Node) GetTransaction(transactionID string) *Transaction {
	var transaction Transaction

	url := usersURL + "/" + n.UserID + "/nodes/" + n.NodeID + "/trans/" + transactionID

	_, err := request.Get(url, "", &transaction)

	if err != nil {
		panic(err)
	}

	return &transaction
}

// CreateTransaction creates a transaction for the specified node
func (n *Node) CreateTransaction(transactionID, data string) *Transaction {
	var transaction Transaction

	url := usersURL + "/" + n.UserID + "/nodes/" + n.NodeID + "/trans/" + transactionID

	_, err := request.Post(url, data, "", &transaction)

	if err != nil {
		panic(err)
	}

	return &transaction
}
