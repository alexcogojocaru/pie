package pie

type Container interface {
	ID() string
}

type ContainerOption func()
