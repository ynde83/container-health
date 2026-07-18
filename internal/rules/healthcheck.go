package rules

import "github.com/ynde83/container-health/internal/models"

type HealthcheckRule struct{}

func (r HealthcheckRule) Check(container *models.ContainerInfo) *models.Issue {
	if container.HasHealthcheck {
		return nil
	}

	return &models.Issue{
		ID:       "NO_HEALTHCHECK",
		Severity: models.Warning,
		Title:    "Container has no healthcheck",
	}
}
