package wrapper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var clientData map[string]interface{}

func init() {
	data, err := loadFile("client_credentials")

	if err != nil {
		panic(err)
	}

	clientData = data["clientData"].(map[string]interface{})
}

func createTestClient() *Client {
	return New(
		clientData["clientID"].(string),
		clientData["clientSecret"].(string),
		clientData["ipAddress"].(string),
		clientData["fingerprint"].(string),
	)
}

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
	data, err := testClient.GetNodes()

	assert.NoError(err)
	assert.NotNil(data.Limit)
	assert.NotNil(data.Page)
	assert.NotNil(data.PageCount)
	assert.NotNil(data.NodeCount)
	assert.NotNil(data.Nodes)
}

/********** OTHER **********/

func Test_GetInstitutions(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	data, err := testClient.GetInstitutions()

	assert.NoError(err)
	assert.NotNil(data.Banks)
}

func Test_GetPublicKey(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	data, err := testClient.GetPublicKey()

	assert.NoError(err)
	assert.NotNil(data.Response)
}

/********** SUBSCRIPTION **********/

func Test_GetSubscriptions(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	data, err := testClient.GetSubscriptions()

	assert.NoError(err)
	assert.NotNil(data.Limit)
	assert.NotNil(data.Page)
	assert.NotNil(data.PageCount)
	assert.NotNil(data.Subscriptions)
	assert.NotNil(data.SubscriptionsCount)
}

func Test_GetSubscription(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	data, err := testClient.GetSubscription("")

	assert.NoError(err)
	assert.NotNil(data.Response)
	assert.NotNil(data.SubscriptionID)
	assert.NotNil(data.URL)
}

func Test_CreateSubscription(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	data, err := testClient.CreateSubscription("")

	assert.NoError(err)
	assert.NotNil(data.Response)
	assert.NotNil(data.SubscriptionID)
	assert.NotNil(data.URL)
}

func Test_UpdateSubscription(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	data, err := testClient.CreateSubscription("")

	assert.NoError(err)
	assert.NotNil(data.Response)
	assert.NotNil(data.SubscriptionID)
	assert.NotNil(data.URL)
}

/********** USER **********/

func Test_GetUsers(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	data, err := testClient.GetUsers()

	assert.NoError(err)
	assert.NotNil(data.Limit)
	assert.NotNil(data.Page)
	assert.NotNil(data.PageCount)
	assert.NotNil(data.Users)
	assert.NotNil(data.UsersCount)
}

func Test_GetUser(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	data, err := testClient.GetUser("", false)

	assert.NoError(err)
	assert.NotNil(data.UserID)
	assert.NotNil(data.AuthKey)
	assert.NotNil(data.FullDehydrate)
	assert.NotNil(data.UserID)
	assert.NotNil(data.RefreshToken)
	assert.NotNil(data.Response)
	assert.NotNil(data.request)
}
