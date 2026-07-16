package main

import (
	"log"

	"github.com/ynde83/container-health/internal/audit"
)

func main() {

	auditor, err := audit.New()
	if err != nil {
		log.Fatal(err)
	}

	defer auditor.Close()

	if err := auditor.Run(); err != nil {
		log.Fatal(err)
	}
}
