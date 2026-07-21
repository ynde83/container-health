package audit

import (
	"log"

	"github.com/ynde83/container-health/internal/docker"
	"github.com/ynde83/container-health/internal/models"
	"github.com/ynde83/container-health/internal/output"
	"github.com/ynde83/container-health/internal/rules"
	"github.com/ynde83/container-health/internal/score"
)

type Auditor struct {
	docker *docker.Client
}

func New() (*Auditor, error) {
	dockerClient, err := docker.New()
	if err != nil {
		return nil, err
	}

	return &Auditor{
		docker: dockerClient,
	}, nil
}

func (a *Auditor) Close() error {
	return a.docker.Close()
}

func (a *Auditor) Run(jsonOutput bool) error {
	containers, err := a.docker.GetContainers()
	if err != nil {
		return err
	}

	runner := rules.NewRunner(rules.Default())

	reports := make([]models.Report, 0, len(containers))

	for _, c := range containers {

		info, err := a.docker.InspectContainer(c.ID)
		if err != nil {
			log.Printf("inspect %s failed: %v\n", c.ID, err)
			continue
		}

		issues := runner.Run(info)

		report := models.Report{
			Container: *info,
			Issues:    issues,
			Score:     score.Calculate(issues),
		}

		reports = append(reports, report)
	}

	if jsonOutput {
		return output.PrintJSON(reports)
	}

	output.Print(reports)

	return nil
}
