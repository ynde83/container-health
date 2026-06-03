package docker

import (
	"context"

	"github.com/docker/docker/client"
)

type DockerClient struct {
	Client *client.Client
	Ctx    context.Context
}

func New() (*DockerClient, error) {
	cli, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)

	if err != nil {
		return nil, err
	}

	return &DockerClient{
		Client: cli,
		Ctx:    context.Background(),
	}, nil
}

func (d *DockerClient) Close() error {
	return d.Client.Close()
}