package wrapper

import (
	"github.com/parnurzeal/gorequest"
)

/********** GLOBAL VARIABLES **********/

const version = "v3.1"

var developerMode = false
var request = gorequest.New()

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
		authKey, clientGateway, clientIP, clientFingerprint, userID, refreshToken string
		fullDehydrate                                                             bool
		Payload                                                                   interface{}
	}
)
