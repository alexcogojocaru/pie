package dockerprovider_test

import (
	"context"
	"testing"

	"github.com/alexcogojocaru/pie/providers/dockerprovider"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestContainerCreate(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()

	p, err := dockerprovider.NewProvider()
	assert.Nil(t, err, "Failed to create provider")

	container, err := p.CreateTask(ctx, id.String(), "demo-library:latest")
	assert.Nil(t, err, "Failed to create task")

	err = p.DeleteTask(ctx, container)
	assert.Nil(t, err, "Failed to delete task")
}

func TestContainerFlow(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()

	p, err := dockerprovider.NewProvider()
	assert.Nil(t, err, "Failed to create provider")

	container, err := p.CreateTask(ctx, id.String(), "demo-library:latest")
	assert.Nil(t, err, "Failed to create task")

	err = p.StartTask(ctx, container)
	assert.Nil(t, err, "Failed to start task")

	err = p.ResetTask(ctx, container)
	assert.Nil(t, err, "Failed to reset task")

	err = p.DeleteTask(ctx, container)
	assert.Nil(t, err, "Failed to delete task")
}
