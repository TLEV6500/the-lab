package signaling

// The Signaling Hub acts as the matchmaker. It exchanges SDP offers and ICE candidates
// so the C++ clients can establish a direct Peer-to-Peer connection.
// Once P2P is established, this hub steps out of the way of the text editing.

type Hub struct {
	// TODO: Map of active sessions and connected WebSockets
}

func (h *Hub) HandleICECandidate(sessionID string, candidateData []byte) {
	// Route candidate data to the other peers in the room
}