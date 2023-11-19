package core

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
)

func (c *Core) AssumeRole() error {
	token := readMFAinput()
	duration := readDurationInput()
	staticTokenProvider := func() (string, error) {
		return *token, nil
	}

	cfg, err := config.LoadDefaultConfig(
		c.Context,
		config.WithRegion(c.Profiles[c.SelectedRole].Region),
		config.WithSharedConfigProfile(c.Profiles[c.SelectedRole].Name),
		config.WithAssumeRoleCredentialOptions(func(aro *stscreds.AssumeRoleOptions) {
			aro.TokenProvider = staticTokenProvider
			aro.SerialNumber = aws.String(c.Profiles[c.SelectedRole].MfaSerial)
			aro.Duration = time.Duration(*duration) * time.Second
		}),
	)
	if err != nil {
		return err
	}

	creds, err := cfg.Credentials.Retrieve(c.Context)
	c.SelectedCreds = &creds
	return err
}
