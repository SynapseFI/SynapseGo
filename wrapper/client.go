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

// GenerateClient creates a client object
// func GenerateClient(gateway, ipAddress, userID string, devMode ...bool) *Client {
func GenerateClient(params interface{}) *Client {
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
