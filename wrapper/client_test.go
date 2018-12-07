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

func createTestClient(t *testing.T) *Client {
	cred := clientData[0]

	return New(
		cred["clientID"].(string),
		cred["clientSecret"].(string),
		cred["ipAddress"].(string),
		cred["fingerprint"].(string),
	)
}

func Test_New(t *testing.T) {
	cred := clientData[0]
	testClient := createTestClient(t)

	// Client credentials should match input credentials
	assert.Equal(t, testClient.ClientID, cred["clientID"].(string))
	assert.Equal(t, testClient.ClientSecret, cred["clientSecret"].(string))
	assert.Equal(t, testClient.IP, cred["ipAddress"].(string))
	assert.Equal(t, testClient.Fingerprint, cred["fingerprint"].(string))

	// Client request headers should match client credentials
	assert.Equal(t, testClient.request.clientID, testClient.ClientID)
	assert.Equal(t, testClient.request.clientSecret, testClient.ClientSecret)
	assert.Equal(t, testClient.request.ipAddress, testClient.IP)
	assert.Equal(t, testClient.request.fingerprint, testClient.Fingerprint)
}
