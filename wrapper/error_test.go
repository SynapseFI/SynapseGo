package wrapper

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var errorData []map[string]interface{}

func init() {
	data, err := loadFile("error_responses")

	if err != nil {
		panic(err)
	}

	errorData = data
}

func loadFile(name string) ([]map[string]interface{}, error) {
	path := filepath.Join("testdata", name+".json") // relative path
	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var data []map[string]interface{}
	e := json.Unmarshal(bytes, &data)

	return data, e

}

func Test_HandleHTTPError(t *testing.T) {
	assert := assert.New(t)

	for i := range errorData {

		testErrRes, _ := json.Marshal(errorData[i])
		testErr := handleHTTPError(testErrRes)

		httpCode := errorData[i]["http_code"].(string)
		errCode := errorData[i]["error_code"].(string)
		msg := errorData[i]["error"].(map[string]interface{})["en"].(string)
		responseMsg := "HTTP_CODE " + httpCode + " ERROR_CODE " + errCode + "\n" + msg

		// error message should be an error and print error code plus original API message
		assert.EqualError(testErr, responseMsg)
	}
}
