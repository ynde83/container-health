package rules

import (
	"strings"

	"github.com/ynde83/container-health/internal/models"
)

type RestartPolicyRule struct{}

func (r RestartPolicyRule) Check(container *models.ContainerInfo) *models.Issue {
	restartPolicy := strings.ToLower(strings.TrimSpace(container.RestartPolicy))
	if restartPolicy != "" && restartPolicy != "no" {
		return nil
	}

	return &models.Issue{
		ID:         "NO_RESTART_POLICY",
		Severity:   models.Warning,
		Title:      "Container has no restart policy",
		Constraint: "Container must define a restart policy",
	}
}
