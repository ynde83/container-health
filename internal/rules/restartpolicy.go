package rules

import "github.com/ynde83/container-health/internal/models"

type RestartPolicyRule struct{}

func (r RestartPolicyRule) Check(container *models.ContainerInfo) *models.Issue {
	if container.RestartPolicy != "" {
		return nil
	}

	return &models.Issue{
		ID:         "NO_RESTART_POLICY",
		Severity:   models.Warning,
		Title:      "Container has no restart policy",
		Constraint: "Container must define a restart policy",
	}
}