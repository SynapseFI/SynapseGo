package wrapper

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func loadFile(t *testing.T, name string) []map[string]interface{} {
	path := filepath.Join("testdata", name+".json") // relative path
	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		t.Fatal(err)
	}

	var data []map[string]interface{}
	e := json.Unmarshal(bytes, &data)

	if e != nil {
		t.Fatal(e)
	}

	return data

}

func Test_HandleHTTPError(t *testing.T) {
	data := loadFile(t, "error_responses")

	for i := range data {

		testErrRes, _ := json.Marshal(data[i])
		testErr := handleHTTPError(testErrRes)

		httpCode := data[i]["http_code"].(string)
		errCode := data[i]["error_code"].(string)
		msg := data[i]["error"].(map[string]interface{})["en"].(string)
		responseMsg := "HTTP_CODE " + httpCode + " ERROR_CODE " + errCode + "\n" + msg

		// error message should be an error and print error code plus original API message
		assert.EqualError(t, testErr, responseMsg)
	}
}
