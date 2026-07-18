package audit

import (
	"fmt"
	"log"

	"github.com/ynde83/container-health/internal/docker"
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

	fmt.Printf("Containers found: %d\n\n", len(containers))

	for _, c := range containers {

		info, err := a.docker.InspectContainer(c.ID)
		if err != nil {
			log.Printf("inspect %s failed: %v\n", c.ID, err)
			continue
		}

		fmt.Printf("Container: %s\n", info.Name)
		fmt.Printf("Image: %s\n", info.Image)
		fmt.Printf("State: %s\n", info.State)

		if info.User == "" {
			fmt.Println("User: root (default)")
		} else {
			fmt.Printf("User: %s\n", info.User)
		}

		fmt.Printf("Restart Count: %d\n", info.RestartCount)

		if info.RestartPolicy == "" {
			fmt.Println("Restart Policy: none")
		} else {
			fmt.Printf("Restart Policy: %s\n", info.RestartPolicy)
		}

		fmt.Printf("Privileged: %t\n", info.Privileged)
		fmt.Printf("Healthcheck: %t\n", info.HasHealthcheck)

		fmt.Println("----------------------------------------")
	}

	return nil
}
