package dockerprovider

import (
	"context"

	"github.com/alexcogojocaru/pie"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type provider struct {
	cli *client.Client
}

func NewProvider() (pie.Provider, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	return &provider{
		cli: cli,
	}, nil
}

func (p *provider) CreateTask(ctx context.Context, name, image string) (pie.Container, error) {
	spec, err := p.cli.ContainerCreate(ctx,
		&container.Config{
			Image: image,
		},
		&container.HostConfig{},
		&network.NetworkingConfig{},
		&v1.Platform{},
		name,
	)
	if err != nil {
		return nil, err
	}

	return Container{
		Spec: spec,
	}, nil
}

func (p *provider) StartTask(ctx context.Context, cont pie.Container) error {
	return p.cli.ContainerStart(ctx, cont.ID(), types.ContainerStartOptions{})
}

func (p *provider) DeleteTask(ctx context.Context, cont pie.Container) error {
	err := p.cli.ContainerStop(ctx, cont.ID(), container.StopOptions{})
	if err != nil {
		return err
	}

	return p.cli.ContainerRemove(ctx, cont.ID(), types.ContainerRemoveOptions{})
}

func (p *provider) ResetTask(ctx context.Context, cont pie.Container) error {
	return p.cli.ContainerKill(ctx, cont.ID(), "SIGUSR1")
}
