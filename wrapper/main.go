package wrapper

/********** GLOBAL VARIABLES **********/

const version = "v3.1"

var authKey string
var developerMode = false

// const _url = "https://api.synapsefi.com/" + version
const _url = "https://uat-api.synapsefi.com/" + version

/********** TYPES **********/

type (
	// ClientCredentials represents the credentials used by the developer to instantiate a client
	ClientCredentials struct {
		gateway, ipAddress, userID string
	}

	// Header represents headers used in http requests
	Header map[string]string

	// A Node represents a single node object
	Node struct {
		NodeID, UserID string
		FullDehydrate  bool
		Payload        Payload
	}

	// Nodes represents multiple node objects
	Nodes struct {
		Limit, NodeCount, Page, PageCount float64
		NodeList                          []Node
	}

	// A Payload is a payload object used to handle a response from the SynapseFI API
	Payload map[string]interface{}

	// A Subscription represents a single subscription object
	Subscription struct {
		SubID, URL string
		Payload    Payload
	}

	// Subscriptions represents multiple transaction objects
	Subscriptions struct {
		Limit, Page, PageCount, SubCount float64
		SubList                          []Subscription
	}

	// A Transaction represents a single transaction object
	Transaction struct {
		TransID string
		Payload Payload
	}

	// Transactions represents multiple transaction objects
	Transactions struct {
		Limit, Page, PageCount, TransCount float64
		TransList                          []Transaction
		Payload                            Payload
	}

	// A User represents a single user object
	User struct {
		UserID, FullDehydrate string
		Payload               Payload
	}

	// Users represents multiple user objects
	Users struct {
		Limit, Page, PageCount float64
		UsersList              []User
		Payload                Payload
	}

	UserTest map[string]interface{}
)
