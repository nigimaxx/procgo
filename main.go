//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./proto/procgo.proto

package main

import (
	"github.com/nigimaxx/procgo/cmd"
)

func main() {
	cmd.Execute()
}
