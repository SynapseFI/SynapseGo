package wrapper

/********** GLOBAL VARIABLES **********/
const transactionsURL = _url + "/trans"

/********** TYPES **********/

type (
	// Transaction represents a single transaction object
	Transaction struct {
		TransactionID string `json:"_id"`
		node          *Node
		Response      interface{}
	}

	// Transactions represents a list of transaction objects
	Transactions struct {
		Limit            int64         `json:"limit"`
		Page             int64         `json:"page"`
		PageCount        int64         `json:"page_count"`
		TransactionCount int64         `json:"trans_count"`
		Transactions     []Transaction `json:"trans"`
	}
)

/********** METHODS **********/
func (t *Transaction) newRequest() *Request {
	return &Request{
		fingerprint: t.node.user.AuthKey + t.node.user.client.Fingerprint,
		gateway:     t.node.user.client.Gateway,
		ipAddress:   t.node.user.client.IP,
	}
}

/********** CLIENT METHODS **********/

// GetClientTransactions returns all client transactions
func (c *Client) GetClientTransactions(queryParams ...string) *Transactions {
	var transactions Transactions

	req := c.newRequest()

	_, err := req.Get(transactionsURL, queryParams[0], &transactions)

	if err != nil {
		panic(err)
	}

	return &transactions
}

// CommentOnStatus adds comment to the transaction status
func (t *Transaction) CommentOnStatus(data string) *Transaction {
	var transaction Transaction

	url := usersURL + "/" + t.node.UserID + "/nodes/" + t.node.NodeID + "/trans/" + t.TransactionID

	req := t.newRequest()

	_, err := req.Post(url, data, "", &transaction)

	if err != nil {
		panic(err)
	}

	return &transaction
}

// CancelTransaction cancels a transaction
func (t *Transaction) CancelTransaction(data string) *Transaction {
	var transaction Transaction

	url := usersURL + "/" + t.node.UserID + "/nodes/" + t.node.NodeID + "/trans/" + t.TransactionID

	req := t.newRequest()

	_, err := req.Delete(url, &transaction)

	if err != nil {
		panic(err)
	}

	return &transaction
}

/********** NODE METHODS **********/

// GetTransaction returns a specific transaction associated with a node
func (n *Node) GetTransaction(transactionID string) *Transaction {
	var transaction Transaction

	url := usersURL + "/" + n.UserID + "/nodes/" + n.NodeID + "/trans/" + transactionID

	req := n.newRequest()

	_, err := req.Get(url, "", &transaction)

	if err != nil {
		panic(err)
	}

	return &transaction
}

// CreateTransaction creates a transaction for the specified node
func (n *Node) CreateTransaction(transactionID, data string) *Transaction {
	var transaction Transaction

	url := usersURL + "/" + n.UserID + "/nodes/" + n.NodeID + "/trans/" + transactionID

	req := n.newRequest()

	_, err := req.Post(url, data, "", &transaction)

	if err != nil {
		panic(err)
	}

	return &transaction
}
