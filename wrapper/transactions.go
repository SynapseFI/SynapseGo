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

/********** TRANSACTION **********/

// CommentOnStatus adds comment to the transaction status
func (t *Transaction) CommentOnStatus(data string) *Transaction {
	var transaction Transaction

	url := usersURL + "/" + t.node.UserID + "/nodes/" + t.node.NodeID + "/trans/" + t.TransactionID

	_, err := request.Post(url, data, "", &transaction)

	if err != nil {
		panic(err)
	}

	return &transaction
}

// CancelTransaction cancels a transaction
func (t *Transaction) CancelTransaction(data string) *Transaction {
	var transaction Transaction

	url := usersURL + "/" + t.node.UserID + "/nodes/" + t.node.NodeID + "/trans/" + t.TransactionID

	_, err := request.Delete(url, &transaction)

	if err != nil {
		panic(err)
	}

	return &transaction
}
