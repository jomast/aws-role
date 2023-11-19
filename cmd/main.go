package main

import (
	"fmt"
	"os"

	"aws-role/core"
)

func main() {
	core := core.New()
	core.Load()
	core.Menu()
	err := core.AssumeRole()
	if err != nil {
		fmt.Printf("Error assuming role: %s\n", err)
		os.Exit(1)
	}

	core.Output()
}
