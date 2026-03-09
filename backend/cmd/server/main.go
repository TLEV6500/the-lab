package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Lab Backend Orchestrator starting...")
	
	// TODO: Init WebRTC Signaling Server
	// TODO: Init K8s Client
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}