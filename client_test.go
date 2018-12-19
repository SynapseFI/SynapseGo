// +build mock

package synapse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var clientData map[string]interface{}
var clientMethodsData map[string]interface{}

/********** METHODS **********/

func init() {
	data, err := readFile("client_credentials")

	if err != nil {
		panic(err)
	}

	mData, mErr := readFile("client_methods")

	if mErr != nil {
		panic(err)
	}

	clientData = data["clientData"].(map[string]interface{})
	clientMethodsData = mData
}

func createTestClient() *Client {
	return New(
		clientData["clientID"].(string),
		clientData["clientSecret"].(string),
		clientData["ipAddress"].(string),
		clientData["fingerprint"].(string),
	)
}

/********** TESTS **********/

/********** CLIENT **********/

func Test_New(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// Client credentials should match input credentials
	assert.Equal(clientData["clientID"].(string), testClient.ClientID)
	assert.Equal(clientData["clientSecret"].(string), testClient.ClientSecret)
	assert.Equal(clientData["ipAddress"].(string), testClient.IP)
	assert.Equal(clientData["fingerprint"].(string), testClient.Fingerprint)
}

/********** NODE **********/

func Test_GetNodes(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	testRes, err := testClient.GetNodes()

	assert.NoError(err)
	assert.NotNil(testRes)
}

/********** OTHER **********/

func Test_GetCryptoMarketData(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	testRes, err := testClient.GetCryptoMarketData()

	assert.NoError(err)
	assert.NotNil(testRes)
}

func Test_GetCryptoQuotes(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	testRes, err := testClient.GetCryptoQuotes()

	assert.NoError(err)
	assert.NotNil(testRes)
}

func Test_GetInstitutions(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	testRes, err := testClient.GetInstitutions()

	assert.NoError(err)
	assert.NotNil(testRes)
}

func Test_LocateATMs(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	testRes, err := testClient.LocateATMs()

	assert.NoError(err)
	assert.NotNil(testRes)
}

func Test_GetPublicKey(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	testRes, err := testClient.GetPublicKey()

	assert.NoError(err)
	assert.NotNil(testRes)
}

/********** SUBSCRIPTION **********/

func Test_GetSubscriptions(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	testRes, err := testClient.GetSubscriptions()

	assert.NoError(err)
	assert.NotNil(testRes)
}

func Test_GetSubscription(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	testRes, err := testClient.GetSubscription("")

	assert.NoError(err)
	assert.NotNil(testRes)
}

func Test_CreateSubscription(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	testRes, err := testClient.CreateSubscription("")

	assert.NoError(err)
	assert.NotNil(testRes)
}

func Test_UpdateSubscription(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	testRes, err := testClient.CreateSubscription("")

	assert.NoError(err)
	assert.NotNil(testRes)
}

/********** USER **********/

func Test_GetUsers(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	testRes, err := testClient.GetUsers()

	assert.NoError(err)
	assert.NotNil(testRes)
}

func Test_GetUser(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	testRes, err := testClient.GetUser("", false)

	assert.NoError(err)
	assert.NotNil(testRes)
}

func Test_CreateUser(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	testUser, err := testClient.CreateUser("")

	assert.NoError(err)
	assert.NotNil(testUser)
}
