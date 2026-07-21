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
	image := ""
	user := ""

	if inspect.Config != nil {
		hasHealthcheck = inspect.Config.Healthcheck != nil
		image = inspect.Config.Image
		user = inspect.Config.User
	}

	restartPolicy := ""
	privileged := false

	if inspect.HostConfig != nil {
		restartPolicy = string(inspect.HostConfig.RestartPolicy.Name)
		privileged = inspect.HostConfig.Privileged
	}

	state := ""
	if inspect.State != nil {
		state = inspect.State.Status
	}

	info := &models.ContainerInfo{
		ID: id,

		Name: strings.TrimPrefix(
			inspect.Name,
			"/",
		),

		Image: image,

		State: state,

		User: user,

		RestartCount: inspect.RestartCount,

		RestartPolicy: restartPolicy,

		Privileged: privileged,

		HasHealthcheck: hasHealthcheck,
	}

	return info, nil
}
