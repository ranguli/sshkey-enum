package sshclient

import (
	"errors"
	"fmt"
	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

type Credentials struct {
	Host     string
	Username string
	Keyfile  string
	Port     int
}

func getKeySigner(keyFile string) (ssh.Signer, error) {
	key, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to read private key: %v", err))
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to parse private key: %v", err))
	}

	return signer, nil
}

func getClientConfig(user string, signer ssh.Signer) *ssh.ClientConfig {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // This is not ideal security-wise
	}

	return config
}

func createHostString(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}

func Connect(creds Credentials) (bool, error) {
	signer, err := getKeySigner(creds.Keyfile)
	if err != nil {
		return false, err
	}

	config := getClientConfig(creds.Username, signer)

	hostString := createHostString(creds.Host, creds.Port)

	// Connect to the remote server and perform the SSH handshake.
	client, err := ssh.Dial("tcp", hostString, config)
	if err != nil {
		return false, errors.New(fmt.Sprintf("Unable to connect: %v", err))
	}
	defer client.Close()

	return true, nil
}
