package main

import (
	"log"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/posthog/posthog-go"
)

type model struct {
	language        string
	applicationName string
	currentDir      string
	installDeps     bool
	startRestack    bool
	openUI          bool
}

func main() {
	// Initialize PostHog client
	client := posthog.NewWithConfig(
		"phc_QAChHsfb5cq65wolzsxiJ6cZk1V9IcfGqCidBWhaLgK",
		posthog.Config{
			Endpoint: "https://us.i.posthog.com",
		},
	)
	defer client.Close()

	language := validateLanguage()
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	m := model{
		currentDir:      currentDir,
		language:        language,
		applicationName: "restack-app",
		startRestack:    true,
		openUI:          true,
	}

	questions := []huh.Field{
		huh.NewInput().
			Title("Welcome to Restack. Let's get you started.").
			Description("Enter application name:").
			Placeholder("restack-app").
			CharLimit(50).
			Value(&m.applicationName),
		huh.NewConfirm().
			Title("Start Restack in Docker? (recommended)").
			Value(&m.startRestack),
	}

	// Ask the first question (application name)
	err = huh.NewForm(huh.NewGroup(questions[0])).Run()
	if err != nil {
		log.Fatal(err)
	}

	// Trigger an event after getting the project name
	client.Enqueue(posthog.Capture{
		Event: "get_started",
		Properties: posthog.NewProperties().
			Set("application_name", m.applicationName).
			Set("language", m.language),
	})

	// Copy files immediately after getting the project name
	if err := m.cloneBoilerplates(); err != nil {
		log.Fatal(err)
	}

	// Ask to start Restack
	err = huh.NewForm(huh.NewGroup(questions[1])).Run()
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
