package rules

import "github.com/ynde83/container-health/internal/models"

type RestartCountRule struct{}

func (r RestartCountRule) Check(container *models.ContainerInfo) *models.Issue {
	if container.RestartCount <= 3 {
		return nil
	}

	return &models.Issue{
		ID:         "HIGH_RESTART_COUNT",
		Severity:   models.Warning,
		Title:      "Container has restarted more than 3 times",
		Constraint: "Container restart count must not be greater than 3",
	}
}
