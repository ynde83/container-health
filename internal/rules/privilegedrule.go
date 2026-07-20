package rules

import "github.com/ynde83/container-health/internal/models"

type PrivilegedRule struct{}

func (r PrivilegedRule) Check(container *models.ContainerInfo) *models.Issue {
	if !container.Privileged {
		return nil
	}

	return &models.Issue{
		ID:         "PRIVILEGED_CONTAINER",
		Severity:   models.Critical,
		Title:      "Container runs in privileged mode",
		Constraint: "Container must not run in privileged mode",
	}
}
