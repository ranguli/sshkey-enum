package cli

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ranguli/sshkey-enum/internal/pkg/file"
	"github.com/ranguli/sshkey-enum/internal/pkg/flagchecker"
)

func usage() {
	fmt.Printf("Usage: sshkey-enum -u/-U username(.txt) -h/-H host(.txt) -p PORT \n")
	flag.PrintDefaults()
	os.Exit(1)
}

func newOptions(host string, hostIsWordlist bool, username string, usernameIsWordlist bool, key string, port int) *options {
	options := options{host: host, hostIsWordlist: hostIsWordlist, username: username, usernameIsWordlist: usernameIsWordlist, key: key, port: port}
	return &options
}

type options struct {
	host               string
	hostIsWordlist     bool
	username           string
	usernameIsWordlist bool
	key                string
	port               int
}

type credentials struct {
	host     string
	username string
	key      string
	port     int
}

func Run() {
	username := flag.String("u", "", "Username")
	usernameFile := flag.String("U", "", "Wordlist of usernames")
	host := flag.String("h", "", "IP Address")
	hostFile := flag.String("H", "", "Wordlist of IPs for multiple SSH servers")
	key := flag.String("i", "", "SSH public key.")
	port := flag.Int("p", 22, "Port for the SSH server")

	flag.Parse()
	flag.Usage = usage

	if flag.NFlag() == 0 {
		usage()
	}

	options, err := parseOptions(*host, *hostFile, *username, *usernameFile, *key, *port)

	if err != nil {
		fmt.Println(err)
		usage()
	}

	fmt.Println(options)
}

func parseOptions(host string, hostFile string, username string, usernameFile string, key string, port int) (*options, error) {
	hostIsWordlist := false
	usernameIsWordlist := false

	log.Printf("Input received:\n host: %s\t username: %s\t key: %s\t", host, username, key)

	if flagchecker.CheckDualFlag(username, usernameFile) {
		return nil, errors.New(badOptionMessage("-u", "-U"))
	}

	if flagchecker.CheckDualFlag(host, hostFile) {
		return nil, errors.New(badOptionMessage("-h", "-H"))
	}

	if len(host) == 0 {
		hostIsWordlist = true

		if len(hostFile) == 0 {
			return nil, errors.New(file.FileErrorMessage(hostFile))
		}
	}

	if len(username) == 0 {
		usernameIsWordlist = true
		if len(usernameFile) == 0 {
			return nil, errors.New(file.FileErrorMessage(usernameFile))
		}
	}

	if len(key) == 0 {
		return nil, errors.New("Please provide an SSH public key with -i")
	} else if !file.FileExists(key) {
		return nil, errors.New(file.FileErrorMessage(key))
	}

	options := newOptions(host, hostIsWordlist, username, usernameIsWordlist, key, port)

	log.Println(options)
	return options, nil
}

func badOptionMessage(option1 string, option2 string) string {
	return fmt.Sprintf("Please provide exactly one of %s or %s.\n", option1, option2)
}
