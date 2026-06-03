package models

type Severity string

const (
	Critical Severity = "critical"
	Warning  Severity = "warning"
)

type Issue struct {
	ID string
	Severity Severity
	Constraint string
	Title string
}

type Report struct {
	Score int
	Issues []Issue
}
