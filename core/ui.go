package core

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func (c *Core) Menu() {
	items := []Profile{}

	for _, p := range c.Profiles {
		items = append(items, p)
	}

	slices.SortFunc(items, func(a, b Profile) int {
		if a.Name < b.Name {
			return -1
		}
		return 1
	})

	for i := 0; i < len(items); i++ {
		fmt.Fprintf(os.Stderr, "%4d. %s\n", i+1, items[i].Name)
	}

	selected := readSelect()
	c.SelectedRole = items[selected-1].Name
}

func (c *Core) Output() {
	output := fmt.Sprintf("export AWS_ACCESS_KEY_ID=\"%s\" ", c.SelectedCreds.AccessKeyID)
	output += fmt.Sprintf("AWS_SECRET_ACCESS_KEY=\"%s\" ", c.SelectedCreds.SecretAccessKey)
	output += fmt.Sprintf("AWS_SESSION_TOKEN=\"%s\" ", c.SelectedCreds.SessionToken)
	output += fmt.Sprintf("AWS_SECURITY_TOKEN=\"%s\" ", c.SelectedCreds.SessionToken)
	output += fmt.Sprintf("AWS_REGION=\"%s\" ", c.Profiles[c.SelectedRole].Region)
	output += fmt.Sprintf("AWS_DEFAULT_REGION=\"%s\"", c.Profiles[c.SelectedRole].Region)
	fmt.Println(output)
	fmt.Fprint(os.Stderr, "AWS Creds Loaded!\n")
}

func readSelect() int {
	var selected int
	fmt.Fprint(os.Stderr, "\nSelect: ")
	fmt.Scanln(&selected)
	return selected
}

func readMFAinput() string {
	var code string
	fmt.Fprint(os.Stderr, "MFA Code: ")
	fmt.Scanln(&code)
	code = strings.TrimSpace(code)
	return code
}

func readDurationInput() time.Duration {
	var hrs int
	fmt.Fprint(os.Stderr, "Duration in hours [1]: ")
	fmt.Scanln(&hrs)
	if hrs == 0 {
		hrs = 1
	}

	hrs *= 60 * 60
	return time.Duration(hrs) * time.Second
}
