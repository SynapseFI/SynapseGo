package wrapper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var userClientData map[string]interface{}

/********** METHODS **********/

func init() {
	data, err := loadFile("client_credentials")

	if err != nil {
		panic(err)
	}

	userClientData = data
}

func createTestUser() *User {
	cred := userClientData["clientData"].(map[string]interface{})

	testClient := New(
		cred["clientID"].(string),
		cred["clientSecret"].(string),
		cred["ipAddress"].(string),
		cred["fingerprint"].(string),
	)

	testUser, err := testClient.GetUser("", false)

	if err != nil {
		panic(err)
	}

	return testUser
}

/********** AUTHENTICATION **********/
func Test_Authenticate(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	data, err := testUser.Authenticate("")

	assert.NoError(err)
	assert.NotNil(data)
}

func Test_GetRefreshToken(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	data, err := testUser.GetRefreshToken()

	assert.NoError(err)
	assert.NotNil(data)
}

/********** NODE **********/

func Test_GetUserNodes(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	data, err := testUser.GetNodes()

	assert.NoError(err)
	assert.NotNil(data)
}

func Test_GetNode(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	data, err := testUser.GetNode("")

	assert.NoError(err)
	assert.NotNil(data)
}

func Test_CreateNode(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	data, err := testUser.CreateNode("")

	assert.NoError(err)
	assert.NotNil(data)
}

func Test_UpdateNode(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	data, err := testUser.UpdateNode("", "")

	assert.NoError(err)
	assert.NotNil(data)
}

// should delete node created in previous test?
func Test_DeleteNode(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	data, err := testUser.DeleteNode("")

	assert.NoError(err)
	assert.NotNil(data)
}

/********** NODE (OTHER) **********/

func Test_AnswerMFA(t *testing.T) {
	// ...
}

func Test_GetApplePayToken(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	data, err := testUser.GetApplePayToken("", "")

	assert.NoError(err)
	assert.NotNil(data)
}

func Test_ReinitiateMicroDeposit(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	data, err := testUser.ReinitiateMicroDeposit("")

	assert.NoError(err)
	assert.NotNil(data)
}

func Test_ResetDebitCard(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	data, err := testUser.ResetDebitCard("")

	assert.NoError(err)
	assert.NotNil(data)
}

func Test_ShipDebitCard(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	data, err := testUser.ShipDebitCard("", "")

	assert.NoError(err)
	assert.NotNil(data)
}

func Test_TriggerDummyTransactions(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	data, err := testUser.TriggerDummyTransactions("", false)

	assert.NoError(err)
	assert.NotNil(data)
}

func Test_VerifyMicroDeposit(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	data, err := testUser.VerifyMicroDeposit("", "")

	assert.NoError(err)
	assert.NotNil(data)
}

/********** STATEMENT **********/

func Test_GetNodeStatements(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	data, err := testUser.GetNodeStatements("")

	assert.NoError(err)
	assert.NotNil(data)
}

func Test_GetStatements(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	data, err := testUser.GetStatements("")

	assert.NoError(err)
	assert.NotNil(data)
}

/********** SUBNET **********/

func Test_GetSubnets(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	data, err := testUser.GetSubnets("")

	assert.NoError(err)
	assert.NotNil(data)
}

func Test_GetSubnet(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	data, err := testUser.GetSubnet("", "")

	assert.NoError(err)
	assert.NotNil(data)
}

func Test_CreateSubnet(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	data, err := testUser.CreateSubnet("", "")

	assert.NoError(err)
	assert.NotNil(data)
}

/********** TRANSACTION **********/

func Test_GetTransactions(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	data, err := testUser.GetTransactions("", "")

	assert.NoError(err)
	assert.NotNil(data)
}

func Test_GetTransaction(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	data, err := testUser.GetTransaction("", "")

	assert.NoError(err)
	assert.NotNil(data)
}

func Test_CreateTransaction(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	data, err := testUser.CreateTransaction("", "", "")

	assert.NoError(err)
	assert.NotNil(data)
}

func Test_DeleteTransaction(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	data, err := testUser.DeleteTransaction("", "")

	assert.NoError(err)
	assert.NotNil(data)

}

func Test_CommentOnTransactionStatus(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	data, err := testUser.CommentOnTransactionStatus("", "", "")

	assert.NoError(err)
	assert.NotNil(data)
}

func Test_DisputeTransaction(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	data, err := testUser.DisputeTransaction("", "")

	assert.NoError(err)
	assert.NotNil(data)
}

/********** USER **********/

func Test_Update(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	data, err := testUser.Update("")

	assert.NoError(err)
	assert.NotNil(data.AuthKey)
	assert.NotNil(data.FullDehydrate)
	assert.NotNil(data.UserID)
	assert.NotNil(data.RefreshToken)
	assert.NotNil(data.Response)
	assert.NotNil(data.request)
}

func Test_CreateUBO(t *testing.T) {
	assert := assert.New(t)
	testUser := createTestUser()

	// No parameters
	data, err := testUser.CreateUBO("")

	assert.NoError(err)
	assert.NotNil(data)
}
