package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type RepoInfo struct {
	URL      string
	IsAtRoot bool
}

var repositoryInfos = map[string]map[string]RepoInfo{
	"typescript": {
		"agent-todo":    {URL: "https://github.com/restackio/examples-typescript.git", IsAtRoot: false},
		"agent-chat":    {URL: "https://github.com/restackio/examples-typescript.git", IsAtRoot: false},
		"agent-scaling": {URL: "https://github.com/restackio/autoscaling-agent.git", IsAtRoot: true},
	},
	"python": {
		"agent_todo": {URL: "https://github.com/restackio/examples-python.git", IsAtRoot: false},
		"agent_chat": {URL: "https://github.com/restackio/examples-python.git", IsAtRoot: false},
	},
}

func (m model) cloneBoilerplates() error {
	targetDir := filepath.Join(m.currentDir, m.applicationName)
	tempDir := filepath.Join(m.currentDir, "temp")

	exampleName := m.example[1:]
	repoInfo := repositoryInfos[m.language][exampleName]

	fmt.Printf("Repository URL: %s, Is at root: %v\n", repoInfo.URL, repoInfo.IsAtRoot)

	cmd := exec.Command("git", "clone", repoInfo.URL, tempDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("git clone failed: %v", err)
	}

	fmt.Printf("Moving files from %s to %s\n", tempDir, targetDir)

	var sourceDir string
	if repoInfo.IsAtRoot {
		// If the template is at the root of the repository, use the whole temp directory
		sourceDir = tempDir
	} else {
		// Otherwise, use the specified subdirectory
		sourceDir = tempDir + m.example
	}

	// Create target directory if it doesn't exist
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return fmt.Errorf("failed to create target directory: %v", err)
	}

	// Use cp -r instead of mv to handle both directory and root cases
	cmd = exec.Command("cp", "-r", sourceDir+"/.", targetDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("copy files failed: %v", err)
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
