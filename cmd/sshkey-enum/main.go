package main

import (
	"io"
	"log"
	"os"

	"github.com/ranguli/sshkey-enum/internal/sshkey-enum/cli"
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
