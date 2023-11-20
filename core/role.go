package core

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
)

func (c *Core) AssumeRole() error {
	profile := c.Profiles[c.SelectedRole]
	duration := readDurationInput()

	cfg, err := config.LoadDefaultConfig(
		c.Context,
		config.WithRegion(profile.Region),
		config.WithSharedConfigProfile(profile.Name),
		config.WithAssumeRoleCredentialOptions(func(aro *stscreds.AssumeRoleOptions) {
			if profile.MfaSerial != "" {
				aro.SerialNumber = aws.String(profile.MfaSerial)
				aro.TokenProvider = func() (string, error) {
					return readMFAinput(), nil
				}
			}
			aro.Duration = duration
		}),
	)
	if err != nil {
		return err
	}

	creds, err := cfg.Credentials.Retrieve(c.Context)
	c.SelectedCreds = &creds
	return err
}
