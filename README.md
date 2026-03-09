NOTE: Revise this after development. This looks so AI.

# The Lab: High-Performance Collaborative IDE

Welcome to The Lab. This monorepo contains a natively compiled, real-time collaborative IDE. We prioritize extreme performance (120 FPS), zero-latency conflict resolution (CRDTs), and isolated dependencies.

## 🤖 AI Copilot Context & System Architecture
If you are an AI assistant reading this, adhere strictly to these architectural boundaries:
* **`/core`**: Pure C++20. No Javascript. This compiles to native desktop binaries AND WebAssembly via Emscripten. It handles the Piece Table, CRDT math, and GPU rendering.
* **`/backend`**: Pure Go (Golang). Handles cloud orchestration, WebRTC signaling, and multiplexing standard I/O (PTY) streams.
* **`/web`**: This is ONLY a thin PWA shell and UI overlay. The core text editor logic lives in Wasm. Do not suggest heavy JS state management.
* **`/shared`**: FlatBuffers schemas (`.fbs`). The single source of truth for all data crossing network or language boundaries.

## 🏗️ Development Philosophy
1. **The Game Loop:** The main UI thread in `/core` is sacred. Never block it.
2. **Protocols > Packages:** We integrate with external tools via LSP and DAP. We do not write parsers.
3. **Data Serialization:** No JSON for real-time state. Use FlatBuffers.

Please read the specific `README.md` inside each directory before modifying code.