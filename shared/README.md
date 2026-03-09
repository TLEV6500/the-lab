NOTE: Revise this after development. This looks so AI.

# Shared Contracts (FlatBuffers)

This directory contains the `.fbs` schemas. This is the single source of truth for data structures that cross language boundaries (C++ <-> Go <-> Wasm).

## 🔨 Workflow
If you modify a `.fbs` file, you MUST run the FlatBuffers compiler (`flatc`) to regenerate the C++ headers and Go structs before submitting a PR.