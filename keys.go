package core

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
