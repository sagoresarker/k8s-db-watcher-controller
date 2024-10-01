/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/sagoresarker/k8s-db-watcher-controller/pkg/controller"
	"github.com/sagoresarker/k8s-db-watcher-controller/pkg/docker"
	"github.com/sagoresarker/k8s-db-watcher-controller/pkg/postgres"
)

func main() {
	log.Println("Starting k8s-db-watcher-controller")

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	log.Println("Environment variables loaded successfully")

	// Get environment variables
	databaseURL := os.Getenv("DATABASE_URL")
	notifyChannel := os.Getenv("NOTIFY_CHANNEL")
	log.Printf("Database URL: %s, Notify Channel: %s", databaseURL, notifyChannel)

	// Initialize PostgreSQL listener
	pgListener, err := postgres.NewListener(databaseURL)
	if err != nil {
		log.Fatalf("Failed to create PostgreSQL listener: %v", err)
	}
	log.Println("PostgreSQL listener initialized successfully")

	// Initialize Docker launcher
	dockerLauncher, err := docker.NewLauncher()
	if err != nil {
		log.Fatalf("Failed to initialize Docker launcher: %v", err)
	}
	log.Println("Docker launcher initialized successfully")

	// Create and start the controller
	ctrl := controller.NewController(pgListener, dockerLauncher, notifyChannel)
	log.Println("Controller created, starting...")

	go ctrl.Run()
	log.Println("Controller is running")

	// Wait for termination signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	log.Println("Waiting for termination signal...")
	<-sigCh

	// Cleanup
	log.Println("Termination signal received, stopping controller...")
	ctrl.Stop()
	log.Println("Controller stopped")
	log.Println("k8s-db-watcher-controller shutdown complete")
}
