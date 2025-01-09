package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func (m model) copyBoilerplates() error {
	targetDir := filepath.Join(m.currentDir, m.projectName)
	sourceDir := filepath.Join("..", "boilerplates", m.language)

	fmt.Printf("Copying files from %s to %s\n", sourceDir, targetDir)

	cmd := exec.Command("cp", "-r", sourceDir, targetDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("copy files failed: %v", err)
	}

	fmt.Println("Files copied successfully")

	return nil
}
