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
func NewClient(params interface{}) *Client {
	p := params.(map[string]interface{})

	if p["devMode"] == true {
		developerMode = true
	}

	return &Client{
		Fingerprint: "|" + p["fingerprint"].(string),
		Gateway:     p["clientID"].(string) + "|" + p["clientSecret"].(string),
		IP:          p["ipAddress"].(string),
	}
}
