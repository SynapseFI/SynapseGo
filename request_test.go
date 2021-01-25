// +build !mock

package synapse

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var credentials map[string]interface{}
var requestData map[string]interface{}
var userID string = "5e6917b85b5a1e0081e0e309" // Change this to test a user on your platform
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
func Test_updateRequest(t *testing.T) {
	assert := assert.New(t)
	testReq := createRequestClient()

	outdatedReq := *testReq
	expectedReq := Request{
		authKey:      "expected_" + testReq.authKey,
		clientID:     "expected_" + testReq.clientID,
		clientSecret: "expected_" + testReq.clientSecret,
		fingerprint:  "expected_" + testReq.fingerprint,
		ipAddress:    "expected_" + testReq.ipAddress,
	}

	testReq.updateRequest(
		expectedReq.clientID,
		expectedReq.clientSecret,
		expectedReq.fingerprint,
		expectedReq.ipAddress,
		expectedReq.authKey,
	)

	assert.NotEqual(outdatedReq, *testReq)
	assert.Equal(expectedReq, *testReq)
}

func Test_Get(t *testing.T) {
	assert := assert.New(t)
	testReq := createRequestClient()

	res, err := testReq.Get(buildURL(path["users"], userID), nil)

	assert.NoError(err)
	assert.NotNil(res)
}

func Test_Post(t *testing.T) {
	assert := assert.New(t)
	testReq := createRequestClient()

	testData := requestData["POST"].(map[string]interface{})["data"]

	jsonData, jsonErr := json.Marshal(testData)

	if jsonErr != nil {
		t.Error(jsonErr)
	}

	userRes, userErr := testReq.Get(buildURL(path["users"], userID), nil)

	if userErr != nil {
		t.Error(userErr)
	}

	rt := readStream(userRes)["refresh_token"].(string)
	authRes, authErr := testReq.Post(buildURL(path["auth"], userID), `{ "refresh_token": "`+rt+`" }`, nil)

	if authErr != nil {
		t.Error(authErr)
	}

	authKey = readStream(authRes)["oauth_key"].(string)
	testReq.authKey = authKey

	res, err := testReq.Post(buildURL(path["users"], userID, "nodes"), string(jsonData), nil)

	if err != nil {
		t.Error(err)
	}

	testID = readStream(res)["nodes"].([]interface{})[0].(map[string]interface{})["_id"].(string)

	assert.NotNil(string(res))
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

	res, err := testReq.Patch(buildURL(path["users"], userID, "nodes", testID), string(jsonData), nil)

	assert.NoError(err)
	assert.NotNil(res)
}

func Test_Delete(t *testing.T) {
	assert := assert.New(t)
	testReq := createRequestClient()
	testReq.authKey = authKey

	res, err := testReq.Delete(buildURL(path["users"], userID, "nodes", testID))

	assert.NoError(err)
	assert.NotNil(res)
}
