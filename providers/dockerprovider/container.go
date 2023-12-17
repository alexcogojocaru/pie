package dockerprovider

import "github.com/docker/docker/api/types/container"

type Container struct {
	Spec container.CreateResponse
}

func (c Container) ID() string {
	return c.Spec.ID
}
