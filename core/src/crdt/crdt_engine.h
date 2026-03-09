#pragma once
#include <cstdint>
// Assuming FlatBuffers generated headers will be included here

namespace Lab::CRDT {
    // Handles the mathematical merging of peer-to-peer operations.
    class Engine {
    public:
        Engine(uint64_t client_id);
        
        // Translates a local keystroke into a network-ready mathematical operation
        void apply_local_insert(char character, uint64_t position);
        
        // Integrates an operation received from a peer via WebRTC
        void merge_remote_operation(/* TODO: Pass FlatBuffer CrdtOperation struct */);
        
    private:
        uint64_t local_client_id;
        uint64_t lamport_clock;
    };
}