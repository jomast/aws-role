package core

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"gopkg.in/ini.v1"
)

const (
	AWS_ACCESS_KEY_ID     = "aws_access_key_id"
	AWS_SECRET_ACCESS_KEY = "aws_secret_access_key"
)

type Core struct {
	Context           context.Context
	AwsConfigFileData []byte
	Profiles          map[string]Profile
	SelectedRole      string
	SelectedCreds     *aws.Credentials
}

type Profile struct {
	Name          string `json:"name"`
	MfaSerial     string `json:"mfa_serial"`
	Region        string `json:"region"`
	SourceProfile string `json:"source_profile"`
}

func New() *Core {
	core := Core{
		Context:  context.Background(),
		Profiles: make(map[string]Profile),
	}

	return &core
}

func (c *Core) Load() {
	c.LoadRaw()
	c.LoadAWSConfig()
}

func (c *Core) LoadRaw() (err error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	c.AwsConfigFileData, err = os.ReadFile(fmt.Sprintf("%s/.aws/config", homeDir))
	if err != nil {
		return err
	}

	return
}

func (c *Core) LoadAWSConfig() error {
	raw, err := ini.Load(c.AwsConfigFileData)
	if err != nil {
		return err
	}

	sections := raw.Sections()

	// Pass 1 - load non-dependant profiles
	for _, section := range sections {
		if section.Name() == "DEFAULT" || section.HasKey("source_profile") {
			continue
		}
		c.loadProfileFromSection(section)
	}

	// Pass 2 - load dependant profiles
	for _, section := range sections {
		if section.Name() == "DEFAULT" || !section.HasKey("source_profile") {
			continue
		}
		c.loadProfileFromSection(section)
	}

	return nil
}

func (c *Core) loadProfileFromSection(section *ini.Section) {
	var mfaSerial, region string
	if section.HasKey("source_profile") {
		source := section.Key("source_profile").Value()
		mfaSerial = c.Profiles[source].MfaSerial
		region = c.Profiles[source].Region
	}

	profileName := strings.Replace(section.Name(), "profile ", "", 1)

	tempProfile := Profile{Name: profileName}
	tempProfile.LoadKey(section, "mfa_serial", mfaSerial)
	tempProfile.LoadKey(section, "region", region)
	tempProfile.LoadKey(section, "source_profile", "")

	c.Profiles[profileName] = tempProfile
}

func (p *Profile) LoadKey(section *ini.Section, key, defaultValue string) {
	if section.HasKey(key) {
		p.SetField(key, section.Key(key).Value())
	} else {
		p.SetField(key, defaultValue)
	}
}

func (p *Profile) SetField(field, value string) {
	switch field {
	case "mfa_serial":
		p.MfaSerial = value
	case "region":
		p.Region = value
	case "source_profile":
		p.SourceProfile = value
	}
}
