package main

import (
	"fmt"
	"log"

	"github.com/ynde83/container-health/internal/docker"
)

func main() {

	dockerClient, err := docker.New()
	if err != nil {
		log.Fatal(err)
	}

	defer dockerClient.Close()

	containers, err := dockerClient.GetContainers()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Containers found: %d\n\n", len(containers))

	for _, c := range containers {

		info, err := dockerClient.InspectContainer(
			c.ID,
		)

		if err != nil {
			log.Printf(
				"inspect %s failed: %v\n",
				c.ID,
				err,
			)
			continue
		}

		fmt.Printf("Container: %s\n", info.Name)
		fmt.Printf("Image: %s\n", info.Image)
		fmt.Printf("State: %s\n", info.State)

		if info.User == "" {
			fmt.Printf("User: root (default)\n")
		} else {
			fmt.Printf("User: %s\n", info.User)
		}

		fmt.Printf(
			"Restart Count: %d\n",
			info.RestartCount,
		)

		if info.RestartPolicy == "" {
			fmt.Printf(
				"Restart Policy: none\n",
			)
		} else {
			fmt.Printf(
				"Restart Policy: %s\n",
				info.RestartPolicy,
			)
		}

		fmt.Printf(
			"Privileged: %t\n",
			info.Privileged,
		)

		fmt.Printf(
			"Healthcheck: %t\n",
			info.HasHealthcheck,
		)

		fmt.Println(
			"----------------------------------------",
		)
	}
}