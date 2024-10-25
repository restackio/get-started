package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/charmbracelet/huh"
)

type model struct {
	projectName  string
	currentDir   string
	packageRoot  string
	installDeps  bool
	startEngine  bool
	openStudio   bool
}

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	packageRoot, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		log.Fatal(err)
	}

	m := model{
		currentDir:  currentDir,
		packageRoot: packageRoot,
	}

	questions := []huh.Field{
		huh.NewInput().
			Title("Welcome to Restack Get Started!").
			Description("Enter project name:").
			Placeholder("restack-get-started").
			CharLimit(50).
			Value(&m.projectName),
		huh.NewConfirm().
			Title("Install dependencies?").
			Value(&m.installDeps),
		huh.NewConfirm().
			Title("Start Restack Engine?").
			Value(&m.startEngine),
		huh.NewConfirm().
			Title("Open Restack Engine Studio?").
			Value(&m.openStudio),
	}

	// Ask the first question (project name)
	err = huh.NewForm(huh.NewGroup(questions[0])).Run()
	if err != nil {
		log.Fatal(err)
	}

	// Copy files immediately after getting the project name
	if err := m.copyFiles(); err != nil {
		log.Fatal(err)
	}

	// Ask the remaining questions
	for _, question := range questions[1:] {
		err = huh.NewForm(huh.NewGroup(question)).Run()
		if err != nil {
			log.Fatal(err)
		}
	}

	if m.installDeps {
		if err := m.installDependencies(); err != nil {
			log.Fatal(err)
		}
	}

	if m.startEngine {
		if err := m.startRestackEngine(); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf(`
Project created successfully!

We suggest that you begin with following commands:

To navigate to the project, run: cd %s

To start the service, run: go run .

To schedule a workflow, run: go run scheduleWorkflow.go
`, m.projectName)
}

func (m model) copyFiles() error {
	targetDir := filepath.Join(m.currentDir, m.projectName)
	filesToCopy := []string{"src", "scheduleWorkflow.go", "go.mod"}

	fmt.Println("Copying files to", targetDir)

	for _, file := range filesToCopy {
		sourceFile := filepath.Join(m.packageRoot, file)
		targetFile := filepath.Join(targetDir, file)

		if err := copyFile(sourceFile, targetFile); err != nil {
			fmt.Println("Error copying file", file, ":", err)
			return err
		}
	}
	fmt.Println("Files copied successfully")

	return nil
}

func (m model) installDependencies() error {
	targetDir := filepath.Join(m.currentDir, m.projectName)
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = targetDir
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func (m model) startRestackEngine() error {
	cmd := exec.Command("docker", "rm", "-f", "studio")
	cmd.Run() // Ignore errors if container doesn't exist

	cmd = exec.Command("docker", "run", "-d", "--pull", "always", "--name", "studio", "-p", "5233:5233", "-p", "6233:6233", "-p", "7233:7233", "ghcr.io/restackio/engine:main")
	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Println("Restack Engine Studio started on http://localhost:5233")

	if m.openStudio {
		time.Sleep(5 * time.Second) // Wait for the server to start
		cmd = exec.Command("open", "http://localhost:5233")
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}

func copyFile(src, dst string) error {
	// Implement file copying logic
	// You can use os.MkdirAll, os.Create, and io.Copy to implement this
	return nil
}
