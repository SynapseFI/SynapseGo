package wrapper

/********** GLOBAL VARIABLES **********/
const version = "v3.1"
const debugMode = false

// const _url = "https://api.synapsefi.com/" + version
const _url = "https://uat-api.synapsefi.com/" + version

/********** STRUCTS **********/

// ClientCredentials structure of client object
type ClientCredentials struct {
	gateway, ipAddress, userID string
}
