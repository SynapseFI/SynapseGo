package wrapper

import (
	"io"
	"net/http"
)

func createRequest(httpMethod, url string, body io.Reader) *http.Request {
	request, err := http.NewRequest(httpMethod, url, body)

	if err != nil {
		errorLog(err)
	}

	return request
}
