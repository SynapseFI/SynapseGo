package wrapper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTestClient(t *testing.T) *Client {
	data := loadFile(t, "client_credentials")
	cred := data[0]

	return New(
		cred["clientID"].(string),
		cred["clientSecret"].(string),
		cred["ipAddress"].(string),
		cred["fingerprint"].(string),
	)
}

func Test_New(t *testing.T) {
	data := loadFile(t, "client_credentials")
	cred := data[0]

	testClient := createTestClient(t)

	assert.Equal(t, testClient.ClientID, cred["clientID"].(string))
	assert.Equal(t, testClient.ClientSecret, cred["clientSecret"].(string))
	assert.Equal(t, testClient.IP, cred["ipAddress"].(string))
	assert.Equal(t, testClient.Fingerprint, cred["fingerprint"].(string))
}
