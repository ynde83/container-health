package docker

import (
	"context"

	"github.com/docker/docker/client"
)

type Client struct {
	Client *client.Client
	Ctx    context.Context
}

func New() (*Client, error) {
	cli, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)

	if err != nil {
		return nil, err
	}

	return &Client{
		Client: cli,
		Ctx:    context.Background(),
	}, nil
}

func (d *Client) Close() error {
	return d.Client.Close()
}
