package rules

import (
	"strings"

	"github.com/ynde83/container-health/internal/models"
)

type RootRule struct{}

func (r RootRule) Check(container *models.ContainerInfo) *models.Issue {
	user := strings.ToLower(strings.TrimSpace(container.User))
	if user != "" && user != "0" && user != "root" {
		return nil
	}

	return &models.Issue{
		ID:         "RUNS_AS_ROOT",
		Severity:   models.Critical,
		Title:      "Container runs as root",
		Constraint: "Container must not run as root",
	}
}
