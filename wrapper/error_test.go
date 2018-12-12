package wrapper

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var errorData map[string]interface{}

/********** METHODS **********/

func init() {
	data, err := readFile("error_responses")

	if err != nil {
		panic(err)
	}

	errorData = data
}

/********** TESTS **********/

func Test_HandleHTTPError(t *testing.T) {
	assert := assert.New(t)

	for k := range errorData {

		testErrRes, _ := json.Marshal(errorData[k])
		testErr := handleHTTPError(testErrRes)

		errData := errorData[k].(map[string]interface{})
		httpCode := errData["http_code"].(string)
		errCode := errData["error_code"].(string)
		msg := errData["error"].(map[string]interface{})["en"].(string)
		responseMsg := "HTTP_CODE " + httpCode + " ERROR_CODE " + errCode + "\n" + msg

		// error message should be an error and print error code plus original API message
		assert.EqualError(testErr, responseMsg)
	}
}
