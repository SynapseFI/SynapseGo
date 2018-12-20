package synapse

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

/********** METHODS **********/

func readFile(name string) (map[string]interface{}, error) {
	path := filepath.Join("testdata", name+".json") // relative path
	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	e := json.Unmarshal(bytes, &data)

	return data, e
}

/********** TESTS **********/

func Test_Read(t *testing.T) {
	var testData = map[string]interface{}{
		"TEST_KEY": "This is a test sample",
	}

	td, err := json.Marshal(testData)

	if err != nil {
		panic(err)
	}

	testRes := readStream(td)

	assert.Equal(t, testData, testRes)
}
