#pragma once

namespace Lab::Render {
    // Abstract interface for pushing pixels. 
    // Implementation will swap between wgpu/Vulkan (Native) and WebGPU (Wasm).
    class GPUPipeline {
    public:
        virtual ~GPUPipeline() = default;
        virtual void initialize() = 0;
        virtual void render_frame() = 0;
        
        // Updates the GPU texture atlas with new font glyphs
        virtual void cache_glyph(char c) = 0; 
    };
}