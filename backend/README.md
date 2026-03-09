NOTE: Revise this after development. This looks so AI.

# Backend Orchestration (Go)

This service manages the "Cloud Lab" infrastructure. It acts as the WebRTC signaling server, provisions Kubernetes containers, and multiplexes shared terminal (PTY) streams.

## 🤖 AI Copilot Rules
* **Language:** Go 1.21+.
* **Concurrency:** Use Goroutines and Channels (CSP) for multiplexing standard I/O streams. Prevent race conditions. Do not rely heavily on global Mutexes.
* **Data Contracts:** All CRDT and network packets must be serialized using the FlatBuffers schemas defined in `../shared/schemas`. Do not parse JSON for real-time events.

## 📁 Structure
* `cmd/server/`: Main application entry point.
* `internal/pty/`: Multiplexer for shared terminal sessions.
* `internal/k8s/`: DevContainer provisioning logic.