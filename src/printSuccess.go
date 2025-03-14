package main

import "fmt"

var depsCmds = map[string]string{
	"typescript": "npm install",
	"python":     "uv sync",
}

var serviceCmds = map[string]string{
	"typescript": "npm run dev",
	"python":     "uv run dev",
}

func (m model) printSuccess() {
	serviceCmd := serviceCmds[m.language]
	depsCmd := depsCmds[m.language]
	const (
		blue  = "\033[34m"
		reset = "\033[0m"
	)

	cdCmd := blue + "cd " + m.applicationName + reset
	depsStr := blue + depsCmd + reset
	serviceStr := blue + serviceCmd + reset

	pythonCmd := ""
	if m.language == "python" {
		pythonCmd = blue + "Start Python shell, run: uv venv && source .venv/bin/activate" + reset
	}

	fmt.Printf(`
Project created successfully!

We suggest that you begin with following commands:

Navigate to the project, run: %s

%s

Install dependencies, run: %s

Start the services, run: %s

And run workflows using the UI
`, cdCmd, pythonCmd, depsStr, serviceStr)
}
