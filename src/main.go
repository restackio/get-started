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
	startEngine  bool
	openUI       bool
}

func main() {

	language := validateLanguage()
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	m := model{
		currentDir: currentDir,
		language:   language,
		projectName: "restack-your-project",
	}

	questions := []huh.Field{
		huh.NewInput().
			Title("Welcome to Restack. Let's get you started.").
			Description("Enter project name:").
			Placeholder("restack-your-project").
			CharLimit(50).
			Value(&m.projectName),
		huh.NewConfirm().
			Title("Install dependencies?").
			Value(&m.installDeps),
		huh.NewConfirm().
			Title("Start Restack Developer UI?").
			Value(&m.startEngine),
		huh.NewConfirm().
			Title("Open Restack Developer UI?").
			Value(&m.openUI),
	}

	// Ask the first question (project name)
	err = huh.NewForm(huh.NewGroup(questions[0])).Run()
	if err != nil {
		log.Fatal(err)
	}

	// Copy files immediately after getting the project name
	if err := m.cloneExampleFolder(); err != nil {
		log.Fatal(err)
	}

	// Ask the second question (install dependencies)
	err = huh.NewForm(huh.NewGroup(questions[1])).Run()
	if err != nil {
		log.Fatal(err)
	}

	if m.installDeps {
		if err := m.installDependencies(); err != nil {
			log.Fatal(err)
		}
	}

	// Ask the third question (start engine)
	err = huh.NewForm(huh.NewGroup(questions[2])).Run()
	if err != nil {
		log.Fatal(err)
	}

	if m.startEngine {
		if err := m.startRestackEngine(); err != nil {
			log.Fatal(err)
		}
	}

	// Ask the fourth question (open developer UI)
	err = huh.NewForm(huh.NewGroup(questions[3])).Run()
	if err != nil {
		log.Fatal(err)
	}

	if m.openUI {
		if err := m.openRestackUI(); err != nil {
			log.Fatal(err)
		}
	}

	m.printSuccess()
}
