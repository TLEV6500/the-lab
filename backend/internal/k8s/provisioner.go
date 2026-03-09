package k8s

import (
	"fmt"
	// "k8s.io/client-go/kubernetes" -> Will be required later
)

// Provisioner is responsible for spinning up isolated DevContainers.
type Provisioner struct {
	// clientset *kubernetes.Clientset
}

// Requests the cluster to spin up a new Pod for a collaborative session.
func (p *Provisioner) CreateDevWorkspace(sessionID string) error {
	fmt.Printf("Instructing Kubernetes to spin up Workspace Pod for Session: %s\n", sessionID)
	// TODO: Define Pod spec utilizing infra/devcontainers/Dockerfile.base
	// TODO: Expose the Pod's internal PTY stream back to the Go signaling hub
	return nil
}