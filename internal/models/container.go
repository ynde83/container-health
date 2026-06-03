package models

type ContainerInfo struct {
	ID   string
	Name string

	Image string
	State string

	User string

	RestartCount int

	RestartPolicy string

	Privileged bool

	HasHealthcheck bool
}