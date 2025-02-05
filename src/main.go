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
	example         string
	installDeps     bool
	startRestack    bool
	openUI          bool
}

func main() {
	// Initialize PostHog client
	client, err := posthog.NewWithConfig(
		"phc_QAChHsfb5cq65wolzsxiJ6cZk1V9IcfGqCidBWhaLgK",
		posthog.Config{
			Endpoint: "https://us.i.posthog.com",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
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
			Title("Welcome to Restack Setup! Let's get started.").
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

	// --- New dropdown for selecting the example template ---
	var exampleQuestion huh.Field
	if m.language == "typescript" {
		exampleQuestion = huh.NewSelect[string]().
			Title("Select template").
			Description("Choose from our TypeScript examples").
			Options([]huh.Option[string]{
				{Value: "/agent-todo", Key: "Default agent (recommended)"},
				{Value: "/agent-chat", Key: "Empty agent"},
			}...).
			Value(&m.example)
	} else if m.language == "python" {
		exampleQuestion = huh.NewSelect[string]().
			Title("Select template").
			Description("Choose from our Python examples").
			Options([]huh.Option[string]{
				{Value: "/agent_todo", Key: "Default agent (recommended)"},
				{Value: "/agent_chat", Key: "Empty agent"},
			}...).
			Value(&m.example)
	}

	// Ask the dropdown question for example selection
	err = huh.NewForm(huh.NewGroup(exampleQuestion)).Run()
	if err != nil {
		log.Fatal(err)
	}
	// --- End dropdown section ---

	// Copy files immediately after getting the project name and example selection
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
