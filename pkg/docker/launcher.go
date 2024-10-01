package docker

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
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
	// Check if the image exists locally
	_, _, err := l.client.ImageInspectWithRaw(ctx, imageName)
	if err != nil {
		if client.IsErrNotFound(err) {
			// Image not found locally, attempt to pull it
			fmt.Printf("Image %s not found locally. Attempting to pull...\n", imageName)
			if err := l.pullImage(ctx, imageName); err != nil {
				return fmt.Errorf("failed to pull image %s: %w", imageName, err)
			}
		} else {
			return fmt.Errorf("error inspecting image %s: %w", imageName, err)
		}
	}

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

// pullImage pulls a Docker image from the registry
func (l *Launcher) pullImage(ctx context.Context, imageName string) error {
	out, err := l.client.ImagePull(ctx, imageName, image.PullOptions{})
	if err != nil {
		return err
	}
	defer out.Close()

	// Print pull progress
	_, err = io.Copy(os.Stdout, out)
	return err
}

// Close closes the Docker client
func (l *Launcher) Close() error {
	return l.client.Close()
}
