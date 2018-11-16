package wrapper

/********** GLOBAL VARIABLES **********/
const version = "v3.1"

var developerMode = false

// const _url = "https://api.synapsefi.com/" + version
const _url = "https://uat-api.synapsefi.com/" + version

/********** TYPES **********/

type (
	// ClientCredentials structure of client object
	ClientCredentials struct {
		gateway, ipAddress, userID string
	}

	// Payload type declaration
	Payload map[string]interface{}

	// User structure of user object
	User struct {
		ID, FullDehydrate string
		Payload           Payload
	}

	// Users structure of users object
	Users struct {
		Limit, Page, PageCount float64
		UsersList              []User
		Payload                Payload
	}
)
