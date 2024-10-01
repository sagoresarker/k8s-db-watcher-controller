package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type Launcher struct {
	client *client.Client
}

// NewLauncher creates a new Launcher instance
func NewLauncher() (*Launcher, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, fmt.Errorf("failed to create Docker client: %w", err)
	}
	return &Launcher{client: cli}, nil
}

// LaunchContainer creates and starts a new container
func (l *Launcher) LaunchContainer(ctx context.Context, imageName string) error {
	resp, err := l.client.ContainerCreate(ctx, &container.Config{
		Image: imageName,
	}, nil, nil, nil, "")
	if err != nil {
		return fmt.Errorf("unable to create container: %w", err)
	}

	if err := l.client.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		return fmt.Errorf("unable to start container: %w", err)
	}

	fmt.Printf("Container started: %s\n", resp.ID)
	return nil
}

// Close closes the Docker client
func (l *Launcher) Close() error {
	return l.client.Close()
}
