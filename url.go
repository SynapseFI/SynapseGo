package synapse

/********** GLOBAL VARIABLErS **********/
const version = "v3.1"

// const baseUrl = "https://api.synapsefi.com/" + version
const baseURL = "https://uat-api.synapsefi.com/" + version

var path = map[string]string{
	"auth":          "oauth",
	"client":        "client",
	"institutions":  "institutions",
	"nodes":         "nodes",
	"statements":    "statements",
	"subnets":       "subnets",
	"subscriptions": "subscriptions",
	"transactions":  "trans",
	"users":         "users",
}

/********** METHODS **********/

func buildURL(uri ...string) string {
	var baseURL = "https://uat-api.synapsefi.com/" + version

	if developerMode != true {
		baseURL = "https://api.synapsefi.com/" + version
	}

	for i := range uri {
		baseURL += "/" + uri[i]
	}

	return baseURL
}
