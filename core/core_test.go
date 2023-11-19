package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	RAW_CRED_DATA = `
[foo]
aws_access_key_id = ABCDEFGHIJKLMNOPQ123
aws_secret_access_key = abcdedfghijklmnopqrstuvwxyz1234567890000

[bar]
aws_access_key_id = ABCDEFGHIJKLMNOPQ888
aws_secret_access_key = abcdedfghijklmnopqrstuvwxyz8888888888888
`

	RAW_CONFIG_DATA = `
[profile foo]
output = json
region = us-east-1

[profile foo-prod]
role_arn = arn:aws:iam::123456789012:role/foo-admin-access
mfa_serial = arn:aws:iam::123456789012:mfa/myuser
source_profile = foo

[profile fizz-prod]
role_arn = arn:aws:iam::12345678900:role/OrganizationAccountAccessRole
mfa_serial = arn:aws:iam::123456789012:mfa/myuser
source_profile = foo

[profile bar]
mfa_serial = arn:aws:iam::120987654321:mfa/myuser
role_arn = arn:aws:iam::120987654321:role/OutoMeAdminAccess
region = us-west-2
output = json
`
)

func TestLoadAWSConfig(t *testing.T) {
	config := New()

	// simulate LoadRaw()
	config.AwsConfigFileData = []byte(RAW_CONFIG_DATA)

	err := config.LoadAWSConfig()
	assert.NoError(t, err)
	assert.Equal(t, "us-west-2", config.Profiles["bar"].Region)
	assert.Equal(t, "us-east-1", config.Profiles["foo-prod"].Region)
}
