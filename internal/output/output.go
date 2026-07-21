package output

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ynde83/container-health/internal/models"
)

func Print(reports []models.Report) {
	fmt.Printf("Containers found: %d\n\n", len(reports))

	for _, report := range reports {
		printReport(report)
		fmt.Println()
	}
}

func PrintJSON(reports []models.Report) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(reports)
}

func printReport(report models.Report) {
	fmt.Printf("Container: %s\n", report.Container.Name)
	fmt.Println()

	fmt.Println("Recommendation:")
	fmt.Printf("  ID: %s\n", report.Container.ID)
	fmt.Printf("  Name: %s\n", report.Container.Name)
	fmt.Printf("  Image: %s\n", report.Container.Image)
	fmt.Printf("  State: %s\n", report.Container.State)
	fmt.Println()

	fmt.Printf("Score: %d\n", report.Score)
	fmt.Println()

	fmt.Println("Issues:")
	if len(report.Issues) == 0 {
		fmt.Println("  None")
		return
	}

	for _, issue := range report.Issues {
		fmt.Printf("  - [%s] %s (%s)\n", issue.Severity, issue.Title, issue.ID)
		fmt.Printf("    Constraint: %s\n", issue.Constraint)
	}
}
