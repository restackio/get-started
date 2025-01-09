package main

import (
	"fmt"
	"os"
)

func validateLanguage() string {
	var lang string
	found := false

	for i, arg := range os.Args[1:] {
		if arg == "--lang" && i+1 < len(os.Args[1:]) {
			lang = os.Args[i+2]
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Error: Please specify language with --lang ('python' or 'typescript')")
		os.Exit(1)
	}

	if lang != "python" && lang != "typescript" {
		fmt.Println("Error: Language must be either 'python' or 'typescript'")
		os.Exit(1)
	}

	return lang
}
