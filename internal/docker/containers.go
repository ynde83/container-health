package docker

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
)

func (d *DockerClient) GetContainers() ([]types.Container, error) {

	return d.Client.ContainerList(
		d.Ctx,
		container.ListOptions{
			All: true,
		},
	)
}