package Raft

import (
	"fmt"
	"testing"
	"time"
)

fmt.Println("=== Starting ===")
func TestNewNode(t *testing.T) {
	fmt.Println("=== TEST: Initializing a Raft Node ===")

	config := Config{
		NodeID:            "node1",
		ElectionTimeout:   5 * time.Second,
		HeartbeatInterval: 1 * time.Second,
		LogDirectory:      "./logs",
	}
	cluster := []string{"node1", "node2", "node3", "node4", "node5"}

	node, err := NewNode(config, cluster)
	if err != nil {
		t.Fatalf("Failed to initialize Raft node: %v", err)
	}

	fmt.Printf("Node initialized: %+v\n", node)
	node.Start()
	time.Sleep(2 * time.Second)
	node.Stop()
}
