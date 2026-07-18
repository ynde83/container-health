package audit

import (
	"fmt"
	"log"

	"github.com/ynde83/container-health/internal/docker"
	"github.com/ynde83/container-health/internal/models"
	"github.com/ynde83/container-health/internal/rules"
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

func (a *Auditor) Run() error {
	containers, err := a.docker.GetContainers()
	if err != nil {
		return err
	}

	runner := rules.NewRunner([]rules.Rule{
		rules.HealthcheckRule{},
	})

	fmt.Printf("Containers found: %d\n\n", len(containers))

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
		}

		fmt.Printf("%+v\n\n", report)
	}

	return nil
}
