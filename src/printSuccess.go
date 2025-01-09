package main

import "fmt"

var depsCmds = map[string]string{
	"typescript": "npm install",
	"python":     "poetry env use 3.10 && poetry shell then poetry install",
}

var serviceCmds = map[string]string{
	"typescript": "npm run dev",
	"python":     "poetry run dev",
}

func (m model) printSuccess() {
	serviceCmd := serviceCmds[m.language]
	depsCmd := depsCmds[m.language]
	const (
		blue  = "\033[34m"
		reset = "\033[0m"
	)

	cdCmd := blue + "cd " + m.projectName + reset
	depsStr := blue + depsCmd + reset
	serviceStr := blue + serviceCmd + reset

	fmt.Printf(`
Project created successfully!

We suggest that you begin with following commands:

Navigate to the project, run: %s

Install dependencies, run: %s

Start the services, run: %s

And run workflows using the UI
`, cdCmd, depsStr, serviceStr)
}
