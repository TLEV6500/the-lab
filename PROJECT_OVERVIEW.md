# Executive Summary: The "Lab"

**Document Type:** Project Overview & Technical Synthesis
**Status:** Active

## 1. Project Vision

The "Lab" is a natively compiled, zero-latency, synchronized virtual war room for engineering teams. It rejects the heavy, bloated, browser-based architecture of modern collaborative editors (like Electron). Instead, it provides a 120+ FPS, natively rendered environment that takes a team from architectural whiteboarding on an infinite canvas all the way through to live, multiplayer debugging in a shared cloud environment.

## 2. Core Engineering Philosophy

* **The "Handmade" Principle:** Core data structures (text buffers, conflict resolution math) and the rendering pipeline are built in-house in C++ to guarantee memory efficiency and sub-16ms render times.
* **Protocols Over Packages:** We do not build language parsers or syntax highlighters. The Lab acts as a client that relies strictly on standardized protocols like the Language Server Protocol (LSP) and Debug Adapter Protocol (DAP).
* **Dependency Isolation:** Third-party libraries are strictly vendored, locked, and hidden behind internal wrapper interfaces.
* **Out-of-Process Extensibility:** External tools (project management, issue trackers) run as separate background processes or cloud services communicating via IPC/API, ensuring they can never crash the main UI thread.

## 3. System Architecture & Tech Stack

The system is divided into four highly specialized, decoupled domains operating within a single monorepo.

### A. Core Engine & UI (C++)

The foundation of the application. It operates similarly to a real-time game engine, decoupling input, state mutation, and rendering into a continuous loop.

* **Text State:** A concurrent Piece Table for $O(1)$ memory-efficient text manipulation.
* **Graphics:** Bypasses OS UI toolkits, rendering directly to the GPU via Vulkan/Metal (native) or WebGPU (web).
* **Portability:** Compiles to both native desktop executables and WebAssembly (Wasm) via Emscripten.

### B. Netcode & Synchronization (C++)

The multiplayer layer responsible for peer-to-peer state merging without a centralized sequencing server.

* **State Merging:** Conflict-Free Replicated Data Types (CRDTs) ensure all connected peers converge on the exact same document state mathematically.
* **Transport:** Direct WebRTC data channels for zero-latency peer-to-peer keystroke transmission.
* **Serialization:** FlatBuffers for network packets, allowing the engine to read binary data with zero parsing overhead.

### C. Cloud Orchestration & Execution (Go)

The backend infrastructure that manages the shared execution environments and heavy network routing.

* **Multiplexing:** Utilizes Go's concurrency (Goroutines/Channels) to safely multiplex standard I/O (PTY) streams, allowing multiple developers to type into the same terminal simultaneously.
* **Provisioning:** Orchestrates Docker and Kubernetes to spin up identical, ephemeral "DevContainers" for every collaborative session.
* **Signaling:** Acts as the WebRTC matchmaker, exchanging connection data so C++ clients can connect peer-to-peer.

### D. The Web Shell & PWA Delivery (Framework Agnostic)

The web wrapper that allows the native C++ engine to be installed and run in the browser as a Progressive Web App (PWA).

* **The Bridge:** Mounts a `<canvas>` element and passes the WebGL/WebGPU context into the Wasm memory space.
* **Caching:** Utilizes Service Workers to aggressively cache the heavy `.wasm` binaries so the IDE loads instantly offline.
* **Peripheral UI:** Handles floating menus, settings, and extension panes using lightweight, non-blocking web technologies layered on top of the Wasm canvas.

## 4. Core Features

* **Multi-Cursor Editing:** Zero-latency collaborative coding backed by P2P CRDTs.
* **Shared Ephemeral Workspaces:** Instant, containerized cloud environments ensuring identical execution contexts for the whole team.
* **Multiplayer PTY (Synchronized Terminals):** Real-time command-line collaboration with Read/Write Role-Based Access Control.
* **Collaborative Debugging:** Synchronized DAP state allowing teams to inspect call stacks and step through code execution simultaneously.
* **Synchronized Infinite Canvas:** A 2D vector drawing space for real-time architecture mapping, with bidirectional links between graphical nodes and repository files.
* **Native Voice Presence:** WebRTC-powered spatial audio channels natively integrated into project rooms.

## 5. System Boundaries & Data Flow

To maintain velocity, strict interface contracts define how these subsystems interact:

* **UI to External Tools:** Communicates via standard JSON-RPC (LSP/DAP).
* **Language to Language (C++ to Go):** Communicates exclusively via shared FlatBuffer schemas (`.fbs`).
* **Client to Cloud:** Communicates via WebSockets (for PTY streams) and WebRTC signaling.
