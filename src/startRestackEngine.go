package main

import (
	"fmt"
	"os"
	"os/exec"
)

func (m model) startRestackEngine() error {
	// Check if Docker is installed
	if _, err := exec.LookPath("docker"); err != nil {
		fmt.Fprintln(os.Stderr, "Docker is not installed. Please install Docker to run Restack Engine. See https://docs.restack.io/troubleshooting/#docker-not-installed for troubleshooting.")
		return fmt.Errorf("docker not installed: %w", err)
	}

	cmd := exec.Command("docker", "rm", "-f", "restack")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	cmd = exec.Command("docker", "run", "-d", "--pull", "always", "--name", "restack", "-p", "5233:5233", "-p", "6233:6233", "-p", "7233:7233", "-p", "9233:9233", "ghcr.io/restackio/restack:main")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Println("Restack Developer UI started on http://localhost:5233")

	return nil
}
