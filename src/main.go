package main

import (
	"log"
	"os"

	"github.com/charmbracelet/huh"
)

type model struct {
	language     string
	projectName  string
	currentDir   string
	installDeps  bool
	startRestack bool
	openUI       bool
}

func main() {

	language := validateLanguage()
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	m := model{
		currentDir:   currentDir,
		language:     language,
		projectName:  "restack-your-project",
		startRestack: true,
		openUI:       true,
	}

	questions := []huh.Field{
		huh.NewInput().
			Title("Welcome to Restack. Let's get you started.").
			Description("Enter project name:").
			Placeholder("restack-your-project").
			CharLimit(50).
			Value(&m.projectName),
		huh.NewConfirm().
			Title("Start Restack Developer UI? (recommended)").
			Value(&m.startRestack),
	}

	// Ask the first question (project name)
	err = huh.NewForm(huh.NewGroup(questions[0])).Run()
	if err != nil {
		log.Fatal(err)
	}

	// Copy files immediately after getting the project name
	if err := m.copyBoilerplates(); err != nil {
		log.Fatal(err)
	}

	// Ask to start Restack
	err = huh.NewForm(huh.NewGroup(questions[2])).Run()
	if err != nil {
		log.Fatal(err)
	}

	if m.startRestack {
		if err := m.startRestackEngine(); err != nil {
			log.Fatal(err)
		}
	}

	m.printSuccess()
}
