package wrapper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var userClientData []map[string]interface{}
var testUser *User

func init() {
	data, err := loadFile("client_credentials")

	if err != nil {
		panic(err)
	}

	userClientData = data
}

func createUserTestClient() *Client {
	cred := userClientData[0]

	return New(
		cred["clientID"].(string),
		cred["clientSecret"].(string),
		cred["ipAddress"].(string),
		cred["fingerprint"].(string),
	)
}

/********** AUTHENTICATION **********/
func Test_Authenticate(t *testing.T) {
	assert := assert.New(t)

	// No parameters
	data, err := testUser.Authenticate("")

	assert.NoError(err)
	assert.NotNil(data)
}

func Test_GetRefreshToken(t *testing.T) {
	assert := assert.New(t)

	// No parameters
	data, err := testUser.GetRefreshToken()

	assert.NoError(err)
	assert.NotNil(data.Token)
}

/********** NODE **********/

func Test_GetUserNodes(t *testing.T) {
	assert := assert.New(t)

	// No parameters
	data, err := testUser.GetNodes()

	assert.NoError(err)
	assert.NotNil(data.Limit)
	assert.NotNil(data.NodeCount)
	assert.NotNil(data.Nodes)
	assert.NotNil(data.Page)
	assert.NotNil(data.PageCount)
}

func Test_GetNode(t *testing.T) {
	assert := assert.New(t)

	// No parameters
	data, err := testUser.GetNode("")

	assert.NoError(err)
	assert.NotNil(data.FullDehydrate)
	assert.NotNil(data.NodeID)
	assert.NotNil(data.Response)
	assert.NotNil(data.UserID)
}

func Test_CreateNode(t *testing.T) {
	assert := assert.New(t)

	// No parameters
	data, err := testUser.CreateNode("")

	assert.NoError(err)
	assert.NotNil(data.Limit)
	assert.NotNil(data.NodeCount)
	assert.NotNil(data.Nodes)
	assert.NotNil(data.Page)
	assert.NotNil(data.PageCount)
}

func Test_UpdateNode(t *testing.T) {
	assert := assert.New(t)

	// No parameters
	data, err := testUser.UpdateNode("", "")

	assert.NoError(err)
	assert.NotNil(data.FullDehydrate)
	assert.NotNil(data.NodeID)
	assert.NotNil(data.Response)
	assert.NotNil(data.UserID)
}

// should delete node created in previous test?
func Test_DeleteNode(t *testing.T) {
	assert := assert.New(t)

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

	// No parameters
	data, err := testUser.GetApplePayToken("", "")

	assert.NoError(err)
	assert.NotNil(data)
}

func Test_ReinitiateMicroDeposit(t *testing.T) {
	assert := assert.New(t)

	// No parameters
	data, err := testUser.ReinitiateMicroDeposit("")

	assert.NoError(err)
	assert.NotNil(data.FullDehydrate)
	assert.NotNil(data.NodeID)
	assert.NotNil(data.Response)
	assert.NotNil(data.UserID)
}

func Test_ResetDebitCard(t *testing.T) {
	assert := assert.New(t)

	// No parameters
	data, err := testUser.ResetDebitCard("")

	assert.NoError(err)
	assert.NotNil(data)
}

func Test_ShipDebitCard(t *testing.T) {
	assert := assert.New(t)

	// No parameters
	data, err := testUser.ShipDebitCard("", "")

	assert.NoError(err)
	assert.NotNil(data)
}

func Test_TriggerDummyTransactions(t *testing.T) {
	assert := assert.New(t)

	// No parameters
	data, err := testUser.TriggerDummyTransactions("", false)

	assert.NoError(err)
	assert.NotNil(data)
}

func Test_VerifyMicroDeposit(t *testing.T) {
	assert := assert.New(t)

	// No parameters
	data, err := testUser.VerifyMicroDeposit("", "")

	assert.NoError(err)
	assert.NotNil(data.FullDehydrate)
	assert.NotNil(data.NodeID)
	assert.NotNil(data.Response)
	assert.NotNil(data.UserID)
}

/********** STATEMENT **********/

func Test_GetNodeStatements(t *testing.T) {
	assert := assert.New(t)

	// No parameters
	data, err := testUser.GetNodeStatements("")

	assert.NoError(err)
	assert.NotNil(data.Limit)
	assert.NotNil(data.Page)
	assert.NotNil(data.PageCount)
	assert.NotNil(data.StatementCount)
	assert.NotNil(data.Statements)
}

func Test_GetStatements(t *testing.T) {
	assert := assert.New(t)

	// No parameters
	data, err := testUser.GetStatements("")

	assert.NoError(err)
	assert.NotNil(data.Limit)
	assert.NotNil(data.Page)
	assert.NotNil(data.PageCount)
	assert.NotNil(data.StatementCount)
	assert.NotNil(data.Statements)
}

/********** SUBNET **********/

func Test_GetSubnets(t *testing.T) {
	assert := assert.New(t)

	// No parameters
	data, err := testUser.GetSubnets("")

	assert.NoError(err)
	assert.NotNil(data.Limit)
	assert.NotNil(data.Page)
	assert.NotNil(data.PageCount)
	assert.NotNil(data.SubnetCount)
	assert.NotNil(data.Subnets)
}

func Test_GetSubnet(t *testing.T) {
	assert := assert.New(t)

	// No parameters
	data, err := testUser.GetSubnet("", "")

	assert.NoError(err)
	assert.NotNil(data.NodeID)
	assert.NotNil(data.Response)
	assert.NotNil(data.SubnetID)
	assert.NotNil(data.UserID)
}

func Test_CreateSubnet(t *testing.T) {
	assert := assert.New(t)

	// No parameters
	data, err := testUser.CreateSubnet("", "")

	assert.NoError(err)
	assert.NotNil(data.NodeID)
	assert.NotNil(data.Response)
	assert.NotNil(data.SubnetID)
	assert.NotNil(data.UserID)
}

/********** TRANSACTION **********/

func Test_GetTransactions(t *testing.T) {
	assert := assert.New(t)

	// No parameters
	data, err := testUser.GetTransactions("", "")

	assert.NoError(err)
	assert.NotNil(data.Limit)
	assert.NotNil(data.Page)
	assert.NotNil(data.PageCount)
	assert.NotNil(data.TransactionCount)
	assert.NotNil(data.Transactions)
}

func Test_GetTransaction(t *testing.T) {
	assert := assert.New(t)

	// No parameters
	data, err := testUser.GetTransaction("", "")

	assert.NoError(err)
	assert.NotNil(data.Response)
	assert.NotNil(data.TransactionID)
}

func Test_CreateTransaction(t *testing.T) {
	assert := assert.New(t)

	// No parameters
	data, err := testUser.CreateTransaction("", "", "")

	assert.NoError(err)
	assert.NotNil(data.Response)
	assert.NotNil(data.TransactionID)
}

func Test_DeleteTransaction(t *testing.T) {
	assert := assert.New(t)

	// No parameters
	data, err := testUser.DeleteTransaction("", "")

	assert.NoError(err)
	assert.NotNil(data.Response)
	assert.NotNil(data.TransactionID)
}

func Test_CommentOnTransactionStatus(t *testing.T) {
	assert := assert.New(t)

	// No parameters
	data, err := testUser.CommentOnTransactionStatus("", "", "")

	assert.NoError(err)
	assert.NotNil(data.Response)
	assert.NotNil(data.TransactionID)
}

func Test_DisputeTransaction(t *testing.T) {
	assert := assert.New(t)

	// No parameters
	data, err := testUser.DisputeTransaction("", "")

	assert.NoError(err)
	assert.NotNil(data)
}

/********** USER **********/

func Test_Update(t *testing.T) {
	assert := assert.New(t)

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

	// No parameters
	data, err := testUser.CreateUBO("")

	assert.NoError(err)
	assert.NotNil(data)
}
