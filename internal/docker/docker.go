package docker

import (
	"context"
	"fmt"
	
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func InspectContainer(id string)(types.ContainerJSON, error){
	cli, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return types.ContainerJSON{}, fmt.Errorf("failed to create client: %v", err)
	}
	defer cli.Close()

	inspect, err := cli.ContainerInspect(context.Background(), id)
	if err != nil {
		fmt.Printf("Error inspecting container %s: %v\n", id, err)
		return types.ContainerJSON{}, fmt.Errorf("failed to inspect container: %v", err)
	}

	return inspect, nil
}

func GetContainers() ([]types.Container, error) {
	cli, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}
	defer cli.Close()

	containers, err := cli.ContainerList(
		context.Background(),
		container.ListOptions{All: true},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to list containers: %v", err)
	}

	return containers, nil
}