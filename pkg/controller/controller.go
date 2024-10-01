package controller

import (
	"context"
	"log"

	"github.com/sagoresarker/k8s-db-watcher-controller/pkg/docker"
	"github.com/sagoresarker/k8s-db-watcher-controller/pkg/postgres"
)

type Controller struct {
	pgListener     *postgres.Listener
	dockerLauncher *docker.Launcher
	stopCh         chan struct{}
}

func NewController(pgListener *postgres.Listener, dockerLauncher *docker.Launcher) *Controller {
	return &Controller{
		pgListener:     pgListener,
		dockerLauncher: dockerLauncher,
		stopCh:         make(chan struct{}),
	}
}

func (c *Controller) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	notifications, err := c.pgListener.Listen(ctx, "image_notify_channel")
	if err != nil {
		log.Fatalf("Failed to start listening: %v", err)
	}

	for {
		select {
		case imageName := <-notifications:
			log.Printf("Received notification for image: %s", imageName)
			err := c.dockerLauncher.LaunchContainer(ctx, imageName)
			if err != nil {
				log.Printf("Failed to launch container: %v", err)
			}
		case <-c.stopCh:
			return
		}
	}
}

func (c *Controller) Stop() {
	close(c.stopCh)
	c.pgListener.Close()
	c.dockerLauncher.Close()
}
