package rules

import (
	"testing"

	"github.com/ynde83/container-health/internal/models"
)

func TestHealthcheckRule_Check(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		container *models.ContainerInfo
		wantIssue *models.Issue
	}{
		{
			name: "returns nil when container has healthcheck",
			container: &models.ContainerInfo{
				HasHealthcheck: true,
			},
			wantIssue: nil,
		},
		{
			name: "returns warning issue when container has no healthcheck",
			container: &models.ContainerInfo{
				HasHealthcheck: false,
			},
			wantIssue: &models.Issue{
				ID:       "NO_HEALTHCHECK",
				Severity: models.Warning,
				Title:    "Container has no healthcheck",
			},
		},
	}

	rule := HealthcheckRule{}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotIssue := rule.Check(tt.container)

			if tt.wantIssue == nil {
				if gotIssue != nil {
					t.Fatalf("Check() = %#v, want nil", gotIssue)
				}

				return
			}

			if gotIssue == nil {
				t.Fatal("Check() = nil, want issue")
			}

			if gotIssue.ID != tt.wantIssue.ID {
				t.Errorf("Check().ID = %q, want %q", gotIssue.ID, tt.wantIssue.ID)
			}

			if gotIssue.Severity != tt.wantIssue.Severity {
				t.Errorf("Check().Severity = %q, want %q", gotIssue.Severity, tt.wantIssue.Severity)
			}

			if gotIssue.Title != tt.wantIssue.Title {
				t.Errorf("Check().Title = %q, want %q", gotIssue.Title, tt.wantIssue.Title)
			}

			if gotIssue.Constraint != tt.wantIssue.Constraint {
				t.Errorf("Check().Constraint = %q, want %q", gotIssue.Constraint, tt.wantIssue.Constraint)
			}
		})
	}
}