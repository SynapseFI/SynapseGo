// +build !mock

package synapse

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var credentials map[string]interface{}
var requestData map[string]interface{}
var testID string
var authKey string

/********** METHODS **********/

func init() {
	data, err := readFile("client_credentials")

	if err != nil {
		panic(err)
	}

	reqData, requestErr := readFile("request_methods")

	if requestErr != nil {
		panic(err)
	}

	credentials = data["clientData"].(map[string]interface{})
	requestData = reqData
}

func createRequestClient() *Request {
	return &Request{
		clientID:     credentials["clientID"].(string),
		clientSecret: credentials["clientSecret"].(string),
		fingerprint:  credentials["fingerprint"].(string),
		ipAddress:    credentials["ipAddress"].(string),
	}
}

/********** TESTS **********/
func Test_Get(t *testing.T) {
	assert := assert.New(t)
	testReq := createRequestClient()

	res, err := testReq.Get(buildURL(path["users"], "/5bec6ebebaabfc00ab168fa0"), nil)

	assert.NoError(err)
	assert.NotNil(res)
}

func Test_Post(t *testing.T) {
	assert := assert.New(t)
	testReq := createRequestClient()

	testData := requestData["POST"].(map[string]interface{})["data"]

	jsonData, err := json.Marshal(testData)

	if err != nil {
		t.Error(err)
	}

	userRes, err := testReq.Get(buildURL(path["users"], "/5bec6ebebaabfc00ab168fa0"), nil)

	if err != nil {
		t.Error(err)
	}

	res, err := readStream(userRes)

	if err != nil {
		t.Error(err)
	}

	authRes, err := testReq.Post(buildURL(path["auth"], "/5bec6ebebaabfc00ab168fa0"), `{ "refresh_token": "`+res["refresh_token"].(string)+`" }`, nil)

	if err != nil {
		t.Error(err)
	}

	authResponse, err := readStream(authRes)

	if err != nil {
		t.Error(err)
	}

	authKey = authResponse["oauth_key"].(string)
	testReq.authKey = authKey

	postRes, err := testReq.Post(buildURL(path["users"], "/5bec6ebebaabfc00ab168fa0/nodes"), string(jsonData), nil)

	if err != nil {
		t.Error(err)
	}

	postResponse, err := readStream(postRes)

	if err != nil {
		t.Error(err)
	}

	testID = postResponse["nodes"].([]interface{})[0].(map[string]interface{})["_id"].(string)

	assert.NotNil(string(postRes))
}

func Test_Patch(t *testing.T) {
	assert := assert.New(t)
	testReq := createRequestClient()
	testReq.authKey = authKey

	testData := requestData["PATCH"].(map[string]interface{})["data"]

	jsonData, err := json.Marshal(testData)

	if err != nil {
		t.Error(err)
	}

	res, err := testReq.Patch(buildURL(path["users"], "/5bec6ebebaabfc00ab168fa0/nodes/", testID), string(jsonData), nil)

	assert.NoError(err)
	assert.NotNil(res)
}

func Test_Delete(t *testing.T) {
	assert := assert.New(t)
	testReq := createRequestClient()
	testReq.authKey = authKey

	res, err := testReq.Delete(buildURL(path["users"], "/5bec6ebebaabfc00ab168fa0/nodes/", testID))

	assert.NoError(err)
	assert.NotNil(res)
}
