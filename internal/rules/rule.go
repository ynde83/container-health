package rules

import "github.com/ynde83/container-health/internal/models"

type Rule interface {
	Check(info *models.ContainerInfo) *models.Issue
}
