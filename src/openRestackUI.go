package main

import (
	"os/exec"
	"time"
)

var uiUrl = "http://localhost:5233"

func (m model) openRestackUI() error {
	for {
		cmd := exec.Command("curl", "-s", "-o", "/dev/null", "-w", "%{http_code}", uiUrl)
		output, err := cmd.Output()
		if err == nil && string(output) == "200" {
			break
		}
		time.Sleep(time.Second)
	}
	cmd := exec.Command("open", uiUrl)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
