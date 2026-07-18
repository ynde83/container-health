package rules

import "github.com/ynde83/container-health/internal/models"

type Rule interface {
	Check(container *models.ContainerInfo) *models.Issue
}
