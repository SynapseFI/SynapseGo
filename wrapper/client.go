package wrapper

/********** GLOBAL VARIABLES **********/
var developerMode = false

/********** TYPES **********/

type (
	// Client represents the credentials used by the developer to instantiate a client
	Client struct {
		Fingerprint string `json:"fingerprint"`
		Gateway     string `json:"gateway"`
		IP          string `json:"ip"`
	}
)

/********** METHODS **********/

// NewClient creates a client object
func NewClient(clientID, clientSecret, fingerprint, ipAddress string, devMode ...bool) *Client {
	if len(devMode) > 0 && devMode[0] == true {
		developerMode = true
	}

	return &Client{
		Fingerprint: "|" + fingerprint,
		Gateway:     clientID + "|" + clientSecret,
		IP:          ipAddress,
	}
}
