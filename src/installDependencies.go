package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

func (m model) installDependencies() error {
	targetDir := filepath.Join(m.currentDir, m.projectName)
	var cmd *exec.Cmd
	if m.language == "python" {
		cmd = exec.Command("poetry", "install")
	} else {
		cmd = exec.Command("npm", "install")
	}
	cmd.Dir = targetDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}