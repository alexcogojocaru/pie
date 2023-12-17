package pie

import "context"

type Provider interface {
	CreateTask(ctx context.Context, name string, image string) (Container, error)
	StartTask(context.Context, Container) error
	DeleteTask(context.Context, Container) error
	ResetTask(context.Context, Container) error
}
