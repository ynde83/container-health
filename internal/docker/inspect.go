package docker

import (
	"strings"

	"github.com/ynde83/container-health/internal/models"
)

func (d *Client) InspectContainer(
	id string,
) (*models.ContainerInfo, error) {

	inspect, err := d.Client.ContainerInspect(
		d.Ctx,
		id,
	)

	if err != nil {
		return nil, err
	}

	hasHealthcheck := false

	if inspect.Config != nil {
		hasHealthcheck = inspect.Config.Healthcheck != nil
	}

	info := &models.ContainerInfo{
		ID: id,

		Name: strings.TrimPrefix(
			inspect.Name,
			"/",
		),

		Image: inspect.Config.Image,

		State: inspect.State.Status,

		User: inspect.Config.User,

		RestartCount: inspect.RestartCount,

		RestartPolicy: string(
			inspect.HostConfig.RestartPolicy.Name,
		),

		Privileged: inspect.HostConfig.Privileged,

		HasHealthcheck: hasHealthcheck,
	}

	return info, nil
}
