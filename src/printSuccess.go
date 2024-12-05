package main

import "fmt"

var serviceCmds = map[string]string{
	"typescript": "npm run dev",
	"python":     "poetry run services",
}

var scheduleCmds = map[string]string{
	"typescript": "npm run schedule",
	"python":     "poetry run schedule",
}

func (m model) printSuccess() {
	serviceCmd := serviceCmds[m.language]
	scheduleCmd := scheduleCmds[m.language]

	const (
		blue = "\033[34m"
		reset = "\033[0m"
	)

	cdCmd := blue + "cd " + m.projectName + reset
	serviceStr := blue + serviceCmd + reset
	scheduleStr := blue + scheduleCmd + reset

	fmt.Printf(`
Project created successfully!

We suggest that you begin with following commands:

To navigate to the project, run: %s

To start the service, run: %s

To schedule a workflow, run: %s
`, cdCmd, serviceStr, scheduleStr)
}
