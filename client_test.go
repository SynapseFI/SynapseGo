// +build mock

package synapse

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var clientData map[string]interface{}
var mockClientResponse = make(map[string]interface{})

/********** METHODS **********/

func init() {
	data, err := readFile("client_credentials")

	if err != nil {
		panic(err)
	}

	clientData = data["clientData"].(map[string]interface{})
	marshallErr := json.Unmarshal(mockResponse, &mockClientResponse)

	// if data is an empty stream this will cause an unmarshal error
	if marshallErr != nil {
		panic(marshallErr)
	}
}

func createTestClient() *Client {
	return New(
		clientData["clientID"].(string),
		clientData["clientSecret"].(string),
		clientData["fingerprint"].(string),
		clientData["ipAddress"].(string),
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

/********** AUTHENTICATION **********/

func Test_GetPublicKey(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	testRes, err := testClient.GetPublicKey()

	assert.NoError(err)
	assert.Equal(testRes, mockClientResponse)
}

/********** NODE **********/

func Test_GetClientNodes(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	testRes, err := testClient.GetNodes()

	assert.NoError(err)
	assert.Equal(testRes, mockClientResponse)
}

func Test_GetTradeMarketData(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	testRes, err := testClient.GetTradeMarketData("")

	assert.NoError(err)
	assert.Equal(testRes, mockClientResponse)
}

/********** OTHER **********/

func Test_GetCryptoMarketData(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	testRes, err := testClient.GetCryptoMarketData()

	assert.NoError(err)
	assert.Equal(testRes, mockClientResponse)
}

func Test_GetCryptoQuotes(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	testRes, err := testClient.GetCryptoQuotes()

	assert.NoError(err)
	assert.Equal(testRes, mockClientResponse)
}

func Test_GetInstitutions(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	testRes, err := testClient.GetInstitutions()

	assert.NoError(err)
	assert.Equal(testRes, mockClientResponse)
}

func Test_LocateATMs(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	testRes, err := testClient.LocateATMs()

	assert.NoError(err)
	assert.Equal(testRes, mockClientResponse)
}

func Test_VerifyAddress(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	testRes, err := testClient.VerifyAddress("")

	assert.NoError(err)
	assert.Equal(testRes, mockClientResponse)
}

func Test_VerifyRoutingNumber(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	testRes, err := testClient.VerifyRoutingNumber("")

	assert.NoError(err)
	assert.Equal(testRes, mockClientResponse)
}

/********** SUBSCRIPTION **********/

func Test_GetSubscriptions(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	testRes, err := testClient.GetSubscriptions()

	assert.NoError(err)
	assert.Equal(testRes, mockClientResponse)
}

func Test_GetSubscription(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	testRes, err := testClient.GetSubscription("")

	assert.NoError(err)
	assert.Equal(testRes, mockClientResponse)
}

func Test_CreateSubscription(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	testRes, err := testClient.CreateSubscription("")

	assert.NoError(err)
	assert.Equal(testRes, mockClientResponse)
}

func Test_UpdateSubscription(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	testRes, err := testClient.CreateSubscription("")

	assert.NoError(err)
	assert.Equal(testRes, mockClientResponse)
}

func Test_GetWebhookLogs(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	testRes, err := testClient.GetWebhookLogs()

	assert.NoError(err)
	assert.Equal(testRes, mockClientResponse)
}

/********** TRANSACTION **********/

func Test_GetClientTransactions(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	testRes, err := testClient.GetTransactions()

	assert.NoError(err)
	assert.Equal(testRes, mockClientResponse)
}

/********** USER **********/

func Test_GetUsers(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	testRes, err := testClient.GetUsers()

	assert.NoError(err)
	assert.Equal(testRes, mockClientResponse)
}

func Test_GetUser(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	testUser, err := testClient.GetUser("", "", "")

	assert.NoError(err)
	assert.NotNil(testUser)
}

func Test_CreateUser(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	testUser, err := testClient.CreateUser("", "", "")

	assert.NoError(err)
	assert.NotNil(testUser)
}
