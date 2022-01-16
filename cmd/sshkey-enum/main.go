package main

import (
	"io"
	"log"
	"os"

	"github.com/ranguli/sshkey-enum/internal/sshkey-enum/cli"
)

var (
	username     string
	usernameFile string
	host         string
	hostFile     string
	key          string
	port         int
)

func init() {
	if os.Getenv("SSHKEYENUM_DEBUG") == "TRUE" {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(io.Discard)
	}
}

func main() {
	cli.Run()
}
