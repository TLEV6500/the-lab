NOTE: Revise this after development. This looks so AI.

# Core Engine & Renderer (C++)

This directory contains the foundational text buffer and GPU rendering pipeline. 

## 🤖 AI Copilot Rules
* **Language:** C++20.
* **Memory Allocation:** Be extremely conservative with heap allocations in the hot path (the render loop). Use arena allocators where possible.
* **Platform Independence:** Never call OS-specific APIs (Win32, POSIX) directly in `src/buffer` or `src/crdt`. Route all platform calls through the `src/platform` abstraction layer, as this code must compile to WebAssembly via Emscripten.

## 📁 Structure
* `src/buffer/`: The Piece Table implementation.
* `src/crdt/`: Mathematical conflict-free replicated data types.
* `src/render/`: WebGPU / native graphics API pipelines.