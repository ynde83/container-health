package main

import (
	"flag"
	"log"
	"os"

	"github.com/ynde83/container-health/internal/audit"
)

const (
	exitOK           = 0
	exitIssuesFound  = 1
	exitRuntimeError = 2
)

func main() {
	jsonOutput := flag.Bool("json", false, "print audit report as JSON")
	flag.Parse()

	auditor, err := audit.New()
	if err != nil {
		log.Print(err)
		os.Exit(exitRuntimeError)
	}

	hasIssues, runErr := auditor.Run(*jsonOutput)
	closeErr := auditor.Close()

	if runErr != nil {
		log.Print(runErr)
		os.Exit(exitRuntimeError)
	}

	if closeErr != nil {
		log.Print(closeErr)
		os.Exit(exitRuntimeError)
	}

	if hasIssues {
		os.Exit(exitIssuesFound)
	}

	os.Exit(exitOK)
}
