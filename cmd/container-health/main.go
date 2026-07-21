package main

import (
	"flag"
	"log"

	"github.com/ynde83/container-health/internal/audit"
)

func main() {
	jsonOutput := flag.Bool("json", false, "print audit report as JSON")
	flag.Parse()

	auditor, err := audit.New()
	if err != nil {
		log.Fatal(err)
	}

	defer auditor.Close()

	if err := auditor.Run(*jsonOutput); err != nil {
		log.Fatal(err)
	}
}
