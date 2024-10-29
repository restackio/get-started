package main

import (
	"fmt"
	"os"
	"os/exec"
)

func (m model) startRestackEngine() error {
	cmd := exec.Command("docker", "rm", "-f", "studio")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	cmd = exec.Command("docker", "run", "-d", "--pull", "always", "--name", "studio", "-p", "5233:5233", "-p", "6233:6233", "-p", "7233:7233", "ghcr.io/restackio/engine:main")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Println("Restack Engine Studio started on http://localhost:5233")

	return nil
}
