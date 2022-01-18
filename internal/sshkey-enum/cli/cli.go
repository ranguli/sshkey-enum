package cli

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ranguli/sshkey-enum/pkg/sshclient"
)

func usage() {
	fmt.Printf("Usage: sshkey-enum -u/-U username(.txt) -h/-H host(.txt) -p PORT \n")
	flag.PrintDefaults()
	os.Exit(1)
}

func newOptions(host string, hostIsWordlist bool, username string, usernameIsWordlist bool, keyFile string, port int) *Options {
	cli_options := Options{host: host, hostIsWordlist: hostIsWordlist, username: username, usernameIsWordlist: usernameIsWordlist, keyFile: keyFile, port: port}
	return &cli_options
}

type Options struct {
	host               string
	hostIsWordlist     bool
	username           string
	usernameIsWordlist bool
	keyFile            string
	port               int
}

func Run() {
	username := flag.String("u", "", "Username")
	usernameFile := flag.String("U", "", "Wordlist of usernames")
	host := flag.String("h", "", "IP Address")
	hostFile := flag.String("H", "", "Wordlist of IPs for multiple SSH servers")
	keyFile := flag.String("i", "", "SSH private key.")
	port := flag.Int("p", 22, "Port for the SSH server")

	flag.Parse()
	flag.Usage = usage

	if flag.NFlag() == 0 {
		usage()
	}

	cli_options, err := parseOptions(*host, *hostFile, *username, *usernameFile, *keyFile, *port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	hosts, err := parseHosts(*cli_options)
	if err != nil {
		fmt.Println(err)
	}

	usernames, err := parseUsernames(*cli_options)
	if err != nil {
		fmt.Println(err)
	}

	for i := range hosts {
		for j := range usernames {
			log.Printf("Host: %s\t Username: %s\n", hosts[i], usernames[j])

			credentials := sshclient.Credentials{Host: hosts[i], Username: usernames[j], Keyfile: cli_options.keyFile, Port: cli_options.port}
			success, _ := sshclient.Connect(credentials)

			if success {
				fmt.Println(successMessage(credentials))
			} else {
				fmt.Println(failureMessage(credentials))
			}
		}
	}
}

func successMessage(creds sshclient.Credentials) string {
	return fmt.Sprintf("[SUCCESS] %s@%s:%d", creds.Username, creds.Host, creds.Port)
}

func failureMessage(creds sshclient.Credentials) string {
	return fmt.Sprintf("[FAIL] %s@%s:%d", creds.Username, creds.Host, creds.Port)
}

func badOptionMessage(option1 string, option2 string) string {
	return fmt.Sprintf("Please provide exactly one of %s or %s.\n", option1, option2)
}
