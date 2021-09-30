package main

import (
	"fennec/cmd"
	"fmt"
)

var (
	version string
	commit  string
)

func main() {
	version := fmt.Sprintf("%s.%s", version, commit)
	cmd.Execute(version)
}
