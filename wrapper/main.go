package wrapper

/********** GLOBAL VARIABLES **********/
const version = "v3.1"
const debugMode = false

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
	// Response structure of response to developer
	Response struct {
		ID      string
		Payload Payload
	}
)
