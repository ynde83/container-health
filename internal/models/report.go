package models

type Report struct {
	Container ContainerInfo
	Issues    []Issue
	Score     int
}
