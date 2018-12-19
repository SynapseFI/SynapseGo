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

	res, err := testReq.Get(usersURL+"/5bec6ebebaabfc00ab168fa0", nil)

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

	userRes, userErr := testReq.Get(usersURL+"/5bec6ebebaabfc00ab168fa0", nil)

	if userErr != nil {
		t.Error(userErr)
	}

	rt := read(userRes)["refresh_token"].(string)
	authRes, authErr := testReq.Post(authURL+"/5bec6ebebaabfc00ab168fa0", `{ "refresh_token": "`+rt+`" }`, nil)

	if authErr != nil {
		t.Error(authErr)
	}

	authKey = read(authRes)["oauth_key"].(string)
	testReq.authKey = authKey

	res, err := testReq.Post(usersURL+"/5bec6ebebaabfc00ab168fa0/nodes", string(jsonData), nil)

	if err != nil {
		t.Error(err)
	}

	testID = read(res)["nodes"].([]interface{})[0].(map[string]interface{})["_id"].(string)

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

	res, err := testReq.Patch(usersURL+"/5bec6ebebaabfc00ab168fa0/nodes/"+testID, string(jsonData), nil)

	assert.NoError(err)
	assert.NotNil(res)
}

func Test_Delete(t *testing.T) {
	assert := assert.New(t)
	testReq := createRequestClient()
	testReq.authKey = authKey

	res, err := testReq.Delete(usersURL + "/5bec6ebebaabfc00ab168fa0/nodes/" + testID)

	assert.NoError(err)
	assert.NotNil(res)
}
