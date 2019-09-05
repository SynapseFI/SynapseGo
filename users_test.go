// +build mock

package synapse

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var userClientData map[string]interface{}
var mockUsersResponse = make(map[string]interface{})

/********** METHODS **********/

func init() {
	testRes, err := readFile("client_credentials")

	if err != nil {
		panic(err)
	}

	userClientData = testRes
	marshallErr := json.Unmarshal(mockResponse, &mockUsersResponse)

	// if data is an empty stream this will cause an unmarshal error
	if marshallErr != nil {
		panic(marshallErr)
	}
}

func createTestUser() *User {
	cred := userClientData["clientData"].(map[string]interface{})

	testClient := New(
		cred["clientID"].(string),
		cred["clientSecret"].(string),
		cred["fingerprint"].(string),
		cred["ipAddress"].(string),
	)

	testUser, err := testClient.GetUser("", "", "")

	if err != nil {
		panic(err)
	}

	return testUser
}

/********** TESTS **********/

/********** AUTHENTICATION **********/
func Test_Authenticate(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.Authenticate("", "", "")

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

func Test_GetRefreshToken(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.GetRefreshToken()

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

func Test_RegisterFingerprint(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.RegisterFingerprint("TEST_FINGERPRINT")

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

func Test_Select2FA(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.Select2FA("")

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

func Test_SubmitMFA(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.SubmitMFA("")

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

func Test_VerifyPIN(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.VerifyPIN("")

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

/********** NODE **********/

func Test_GetNodes(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.GetNodes()

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

func Test_GetNode(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.GetNode("")

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

func Test_CreateNode(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.CreateNode("")

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

func Test_UpdateNode(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.UpdateNode("", "")

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

// should delete node created in previous test?
func Test_DeleteNode(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.DeleteNode("")

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

/********** NODE (OTHER) **********/

func Test_VerifyMicroDeposit(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.VerifyMicroDeposit("", "")

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

func Test_ReinitiateMicroDeposits(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.ReinitiateMicroDeposits("")

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

func Test_ResetCardNode(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.ResetCardNode("")

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

func Test_ShipCardNode(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.ShipCardNode("", "")

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

func Test_GetApplePayToken(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.GetApplePayToken("", "")

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

/********** STATEMENT **********/

func Test_GetStatements(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.GetStatements("")

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

func Test_GetNodeStatements(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.GetNodeStatements("")

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

func Test_CreateNodeStatements(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.CreateNodeStatements("", "")

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

/********** SUBNET **********/

func Test_GetSubnets(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.GetSubnets()

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

func Test_GetNodeSubnets(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.GetNodeSubnets("")

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

func Test_GetSubnet(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.GetSubnet("", "")

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

func Test_CreateSubnet(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.CreateSubnet("", "")

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

func Test_UpdateSubnet(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.UpdateSubnet("", "", "")

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

func Test_ShipCard(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	testRes, err := testUser.ShipCard("", "", "")

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

/********** TRANSACTION **********/

func Test_GetTransactions(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.GetTransactions("", "")

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

func Test_GetNodeTransactions(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	testRes, err := testUser.GetNodeTransactions("")

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

func Test_GetTransaction(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.GetTransaction("", "")

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

func Test_CreateTransaction(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.CreateTransaction("", "")

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

func Test_CancelTransaction(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.CancelTransaction("", "")

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)

}

func Test_CommentOnTransactionStatus(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.CommentOnTransactionStatus("", "", "")

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

func Test_DisputeTransaction(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.DisputeTransaction("", "", "")

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

func Test_CreateDummyTransaction(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.CreateDummyTransaction("", "")

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}

/********** USER **********/

func Test_Update(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.Update("")

	assert.NoError(err)
	assert.NotNil(testRes.AuthKey)
	assert.NotNil(testRes.UserID)
	assert.NotNil(testRes.RefreshToken)
	assert.NotNil(testRes.Response)
	assert.NotNil(testRes.request)
}

func Test_CreateUBO(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	testRes, err := testUser.CreateUBO("")

	assert.NoError(err)
	assert.Equal(testRes, mockUsersResponse)
}
