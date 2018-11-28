package wrapper

/********** GLOBAL VARIABLES **********/
var developerMode = false

/********** TYPES **********/

type (
	// Client represents the credentials used by the developer to instantiate a client
	Client struct {
		Gateway, IP, Fingerprint string
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
		Gateway:     p["clientID"].(string) + "|" + p["clientSecret"].(string),
		IP:          p["ipAddress"].(string),
		Fingerprint: "|" + p["fingerprint"].(string),
	}
}
