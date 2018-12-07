package wrapper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var clientData []map[string]interface{}

func init() {
	data, err := loadFile("client_credentials")

	if err != nil {
		panic(err)
	}

	clientData = data
}

func createTestClient() *Client {
	cred := clientData[0]

	return New(
		cred["clientID"].(string),
		cred["clientSecret"].(string),
		cred["ipAddress"].(string),
		cred["fingerprint"].(string),
	)
}

func Test_New(t *testing.T) {
	assert := assert.New(t)
	cred := clientData[0]
	testClient := createTestClient()

	// Client credentials should match input credentials
	assert.Equal(testClient.ClientID, cred["clientID"].(string))
	assert.Equal(testClient.ClientSecret, cred["clientSecret"].(string))
	assert.Equal(testClient.IP, cred["ipAddress"].(string))
	assert.Equal(testClient.Fingerprint, cred["fingerprint"].(string))

	// Client request headers should match client credentials
	assert.Equal(testClient.request.clientID, testClient.ClientID)
	assert.Equal(testClient.request.clientSecret, testClient.ClientSecret)
	assert.Equal(testClient.request.ipAddress, testClient.IP)
	assert.Equal(testClient.request.fingerprint, testClient.Fingerprint)
}

func Test_GetNodes(t *testing.T) {
	assert := assert.New(t)
	testClient := createTestClient()

	// No parameters
	data, err := testClient.GetNodes()

	assert.NoError(err)
	assert.NotNil(data.Limit)
}
