# Web Shell & PWA Delivery

This directory contains the web client delivery mechanism. It acts as a thin, highly optimized host environment for the compiled C++ WebAssembly (Wasm) core engine. 

This layer is strictly a wrapper and a peripheral UI host. It does not contain any core IDE business logic, text buffer management, or CRDT synchronization math.

## 🎯 Core System Responsibilities

1. **The Native Engine Mount Point:**
   * Provide the full-screen `<canvas>` element.
   * Initialize and load the Emscripten-compiled `.wasm` binary and its JavaScript glue code.
   * Pass the WebGL/WebGPU context from the browser down to the C++ render pipeline.

2. **Input & Event Routing (The Browser Bridge):**
   * Capture standard browser events (`keydown`, `pointermove`, `resize`).
   * Intercept these events, prevent their default browser behaviors (like scrolling or standard browser shortcuts), and route them directly into the Wasm memory space for the C++ engine to process.

3. **Progressive Web App (PWA) Lifecycle:**
   * Serve the `manifest.json` to allow the web client to be installed as a standalone desktop application.
   * Manage the Service Worker to aggressively cache the heavy `.wasm` binaries and UI assets, ensuring the IDE loads instantly on subsequent visits without hitting the network.

4. **Peripheral UI & Out-of-Process Extensions:**
   * Render floating, non-editor UI elements (e.g., settings menus, context-aware extension panes, project management dashboards) using standard DOM elements layered *on top* of the Wasm canvas.
   * Communicate with the underlying C++ engine via strict JavaScript Interop/IPC, ensuring the DOM never blocks the WebGPU render loop.

## 🤖 AI Copilot Rules & System Boundaries
* **Framework Agnostic:** This directory may utilize any lightweight web technology (vanilla JS, Web Components, or minimal frameworks), provided it ships minimal JavaScript to the client.
* **The Prime Directive:** The main thread must be reserved for the Wasm engine. Any DOM manipulation or UI rendering here must be strictly optimized to prevent garbage collection pauses from stuttering the C++ 120 FPS game loop.
* **No Core Logic:** Do not attempt to read, write, or synchronize the text buffer from JavaScript. All text manipulation belongs in `/core`.