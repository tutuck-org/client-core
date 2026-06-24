package core

import (
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/crypto/ssh"
)

func getHomeDir() string {
	home, _ := os.UserHomeDir()
	return home
}

func ConnectSSH(server string) (io.ReadWriteCloser, error) {
	dir := filepath.Join(getHomeDir(), ".tutuck")
	privPath := filepath.Join(dir, "key")

	if _, err := os.Stat(privPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("No key found, run '%s generate' first", os.Args[0])
	}

	key, err := os.ReadFile(privPath)
	if err != nil {
		return nil, fmt.Errorf("Cannot read key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse key: %v", err)
	}

	config := &ssh.ClientConfig{
		User:            "any",
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error { return verifyHost(server, key) },
		Timeout:         5 * time.Second,
	}

	client, err := ssh.Dial("tcp", server, config)
	if err != nil {
		return nil, err
	}

	channel, _, err := client.OpenChannel("session", nil)
	if err != nil {
		return nil, err
	}

	return struct {
		io.Reader
		io.Writer
		io.Closer
	}{channel, channel, client}, nil
}
