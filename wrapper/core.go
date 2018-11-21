package wrapper

/********** GLOBAL VARIABLES **********/

const version = "v3.1"

var developerMode = false

// const _url = "https://api.synapsefi.com/" + version
const _url = "https://uat-api.synapsefi.com/" + version

/********** TYPES **********/

type (
	// Client represents the credentials used by the developer to instantiate a client
	Client struct {
		gateway, ipAddress, fingerprint string
	}

	// User represents a user object
	User struct {
		AuthKey, clientGateway, clientIP, clientFingerprint, UserID, RefreshToken string
		fullDehydrate                                                             bool
		Payload                                                                   interface{}
	}
)
