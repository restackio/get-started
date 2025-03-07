package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var repositoryUrls = map[string]string{
	"typescript": "https://github.com/restackio/examples-typescript.git",
	"python":     "https://github.com/restackio/examples-python.git",
}

func (m model) cloneBoilerplates() error {
	targetDir := filepath.Join(m.currentDir, m.applicationName)
	tempDir := filepath.Join(m.currentDir, "temp")

	repoName := repositoryUrls[m.language]

	cmd := exec.Command("git", "clone", repoName, tempDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("git clone failed: %v", err)
	}

	fmt.Printf("Moving files from %s to %s\n", tempDir, targetDir)

	cmd = exec.Command("mv", tempDir+m.example, targetDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("move files failed: %v", err)
	}

	cmd = exec.Command("rm", "-rf", tempDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error removing temp directory: %v\n", err)
	}

	fmt.Println("Repository cloned successfully")

	return nil
}
