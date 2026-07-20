package score

import "github.com/ynde83/container-health/internal/models"

const (
	maxScore          = 100
	minScore          = 0
	warningDeduction  = 10
	criticalDeduction = 30
)

func Calculate(issues []models.Issue) int {
	score := maxScore

	for _, issue := range issues {
		score -= deduction(issue.Severity)
	}

	return clamp(score)
}

func deduction(severity models.Severity) int {
	switch severity {
	case models.Warning:
		return warningDeduction
	case models.Critical:
		return criticalDeduction
	default:
		return 0
	}
}

func clamp(score int) int {
	if score < minScore {
		return minScore
	}

	if score > maxScore {
		return maxScore
	}

	return score
}
