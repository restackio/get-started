package main

import (
	"fmt"
	"os"
)

func validateLanguage() string {
	var lang string
	found := false

	// Iterate over os.Args starting from index 1 (skip program name)
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "--lang" {
			if i+1 < len(os.Args) {
				lang = os.Args[i+1]
				found = true
			} else {
				// Flag provided without an accompanying value
				fmt.Println("Error: --lang flag provided without a value")
				os.Exit(1)
			}
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
