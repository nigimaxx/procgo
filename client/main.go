package main

import "github.com/nigimaxx/procgo/client/cmd"

var version = "dev"

func init() {
	cmd.SetVersion(version)
}

func main() {
	cmd.Execute()
}
