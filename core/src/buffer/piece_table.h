#pragma once
#include <string>
#include <vector>

namespace Lab::Buffer {
    // The fundamental data structure for text. O(1) appends, O(log N) inserts.
    // Designed to never block the main render loop, even with 500MB log files.
    class PieceTable {
    public:
        PieceTable();
        ~PieceTable() = default;

        void insert(size_t index, const std::string& text);
        void remove(size_t index, size_t length);
        
        // Polled by the Render Engine at 120Hz. Must be extremely fast.
        std::string get_visible_text(size_t start_index, size_t length) const;
        
    private:
        std::string original_buffer;
        std::string add_buffer;
        // TODO: Define the Piece struct (pointer, length, buffer_origin)
    };
}