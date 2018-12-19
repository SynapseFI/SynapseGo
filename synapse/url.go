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

var authURL = buildURL(baseURL, path["auth"])
var clientURL = buildURL(baseURL, path["client"])
var institutionsURL = buildURL(baseURL, path["institutions"])
var nodesURL = buildURL(baseURL, path["nodes"])
var subscriptionsURL = buildURL(baseURL, path["subscriptions"])
var transactionsURL = buildURL(baseURL, path["transactions"])
var usersURL = buildURL(baseURL, path["users"])

/********** METHODS **********/

func buildURL(basePath string, uri ...string) string {
	url := basePath

	for i := range uri {
		url += "/" + uri[i]
	}

	return url
}
