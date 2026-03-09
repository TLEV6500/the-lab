package pty

import (
	"fmt"
	"sync"
)

// Multiplexer handles broadcasting stdout from a single PTY to multiple connected clients,
// and safely funneling stdin from multiple clients to the single PTY.
type Multiplexer struct {
	mu          sync.RWMutex
	clients     map[string]chan []byte // Channels mapped to client IDs
	ptyStdout   chan []byte            // Data coming from the container's shell
}

func NewMultiplexer() *Multiplexer {
	return &Multiplexer{
		clients:   make(map[string]chan []byte),
		ptyStdout: make(chan []byte, 1024),
	}
}

// Broadcasts output from the shared terminal to all connected UI clients.
func (m *Multiplexer) BroadcastLoop() {
	for data := range m.ptyStdout {
		m.mu.RLock()
		for _, clientChan := range m.clients {
			// Non-blocking send to avoid one slow client freezing the terminal for everyone
			select {
			case clientChan <- data:
			default:
				fmt.Println("Warning: Dropped PTY packet for slow client")
			}
		}
		m.mu.RUnlock()
	}
}