package core

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/ssh"
)

func GenerateKeys() error {
	dir := filepath.Join(getHomeDir(), ".tutuck")
	privPath := filepath.Join(dir, "key")

	if _, err := os.Stat(privPath); err == nil {
		return fmt.Errorf("Keys already exist in %s", dir)
	}

	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("Cannot create %s: %v", dir, err)
	}

	cmd := exec.Command("ssh-keygen", "-t", "ed25519", "-f", privPath, "-N", "")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ssh-keygen failed: %v", err)
	}

	fmt.Println("Keys generated in ~/.tutuck/")
	fmt.Println("Tutuck is ready to use!")
	return nil
}

func knownHostsPath() string {
	return filepath.Join(getHomeDir(), ".tutuck", "known_hosts")
}

func loadKnownHosts() (map[string]string, error) {
	hosts := make(map[string]string)

	data, err := os.ReadFile(knownHostsPath())
	if os.IsNotExist(err) {
		return hosts, nil
	}
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Fields(line)

		if len(parts) != 2 {
			continue
		}

		hosts[parts[0]] = parts[1]
	}

	return hosts, nil
}

func saveKnownHost(host string, fingerprint string) error {
	hosts, err := loadKnownHosts()
	if err != nil {
		return err
	}

	hosts[host] = fingerprint
	var lines []string

	for host, fp := range hosts {
		lines = append(lines, host+" "+fp)
	}

	content := strings.Join(lines, "\n")

	return os.WriteFile(knownHostsPath(), []byte(content), 0600)

}

func verifyHost(host string, key ssh.PublicKey) error {
	fp := ssh.FingerprintSHA256(key)

	hosts, err := loadKnownHosts()
	if err != nil {
		return err
	}

	knownFP, exists := hosts[host]

	if exists {
		if knownFP == fp {
			return nil
		}

		return fmt.Errorf("WARNING: server fingerprint changed\n expected: %s\n got: %s", knownFP, fp)
	} else {
		fmt.Printf("\nUnknown server\n\n Host: %s\n Fingerprint:\n  %s\n\n", host, fp)
		fmt.Printf("\nTrust this server? [y/N]: ")
		var answer string
		fmt.Scanln(&answer)
		answer = strings.ToLower(strings.TrimSpace(answer))
		if answer != "y" && answer != "yes" {
			return fmt.Errorf("Server not trusted")
		}

		if err := saveKnownHost(host, fp); err != nil {
			return err
		}

		return nil

	}
}
