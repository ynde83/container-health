package rules

import "github.com/ynde83/container-health/internal/models"

type Runner struct {
	rules []Rule
}

func NewRunner(rules []Rule) *Runner {
	return &Runner{
		rules: rules,
	}
}

func (r *Runner) Run(container *models.ContainerInfo) []models.Issue {
	var issues []models.Issue

	for _, rule := range r.rules {
		issue := rule.Check(container)

		if issue != nil {
			issues = append(issues, *issue)
		}
	}

	return issues
}
