#pragma once

namespace Lab::Platform {
    // The abstraction layer. 
    // The core engine talks to this, never directly to Win32 or the DOM.
    class WindowManager {
    public:
        virtual ~WindowManager() = default;
        virtual void create_window(int width, int height, const char* title) = 0;
        virtual void poll_events() = 0; 
        virtual bool should_close() const = 0;
    };
}