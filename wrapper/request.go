package wrapper

import (
	"github.com/parnurzeal/gorequest"
)

/********** GLOBAL VARIABLES **********/
var request = gorequest.New()

const get = "get"
const post = "post"
const patch = "patch"

/********** METHODS **********/

func apiRequest(method, url string, data ...string) *gorequest.SuperAgent {
	var req = gorequest.New()
	req = setMethod(method, url)

	switch len(data) {
	case 1:
		req = req.Send(data[0])

	case 2:
		req = req.Send(data[0]).Query(data[1])
	}

	return req
}

func setHeader() {

}

func setMethod(m, u string) *gorequest.SuperAgent {
	switch m {
	case "post":
		return request.Post(u)

	case "patch":
		return request.Patch(u)

	default:
		return request.Get(u)
	}
}
