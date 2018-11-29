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

func (c *Client) newRequest(headers interface{}) *Request {
	return &Request{
		fingerprint: c.Fingerprint,
		gateway:     c.Gateway,
		ipAddress:   c.IP,
		headers:     headers,
	}
}

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
