package Raft
// Package raft provides the implementation of the Raft consensus algorithm.
// It handles state management, log replication, leader election, and coordination
// between Raft nodes for maintaining a consistent state across the cluster.
//
// This file is the main entry point for your Raft implementation.
// define the overall behavior of Raft, initialize components, and coordinate the roles (leader, follower, etc.)
//

import (
	"errors"
	"sync"
	"time"
)

// LogEntry represents a single log entry in the Raft protocol.
type LogEntry struct {
	Term    int         // The term number for the log entry
	Index   int         // The index of the log entry
	Command interface{} // The command stored in this log entry
}

// Config holds the configuration for a Raft node.
type Config struct {
	NodeID            string        // Unique ID for the node
	ElectionTimeout   time.Duration // Time to wait before starting an election
	HeartbeatInterval time.Duration // Time interval for leader heartbeats
	LogDirectory      string        // Directory for storing logs
}

// Node represents a single Raft node.
type Node struct {
	ID         string        // Unique identifier for the node
	Role       string        // Current role of the node: Leader, Follower, or Candidate
	CurrentTerm int          // Latest term the node has seen
	VotedFor    string       // ID of the candidate this node voted for
	Log         []LogEntry   // Log entries for this node
	CommitIndex int          // Index of the highest log entry known to be committed
	LastApplied int          // Index of the highest log entry applied to state machine
	Cluster     []string     // List of other nodes in the cluster
	Config      Config       // Node configuration
	mu          sync.Mutex   // Mutex for thread-safe access
	CommitChan  chan []byte  // Channel to notify committed log entries
	StopChan    chan bool    // Channel to signal node shutdown
}


// NewNode initializes a new Raft node with the given configuration and cluster.
func NewNode(config Config, cluster []string) (*Node, error) {
	fmt.Println("Initializing a new Raft node...")

	// Validate configuration
	if config.NodeID == "" {
		return nil, errors.New("NodeID cannot be empty")
	}
	fmt.Printf("Node ID: %s, Cluster: %v\n", config.NodeID, cluster)

	// Initialize the node
	node := &Node{
		ID:         config.NodeID,
		Role:       "Follower", // Start as a follower
		CurrentTerm: 0,
		VotedFor:    "",
		Log:         []LogEntry{},
		CommitIndex: 0,
		LastApplied: 0,
		Cluster:     cluster,
		Config:      config,
		CommitChan:  make(chan []byte, 100),
		StopChan:    make(chan bool),
	}

	fmt.Printf("Raft node %s successfully created with role %s.\n", node.ID, node.Role)
	return node, nil
}

// Start begins the Raft node's operations.
func (n *Node) Start() {
	fmt.Printf("Node %s is starting as a %s...\n", n.ID, n.Role)

	// Placeholder: Add logic for follower behavior
	go n.runFollower()

	// Monitor stop signal
	go func() {
		<-n.StopChan
		fmt.Printf("Node %s is stopping.\n", n.ID)
	}()
}

// Stop gracefully shuts down the Raft node.
func (n *Node) Stop() {
	fmt.Printf("Stopping node %s...\n", n.ID)
	n.StopChan <- true
}

// runFollower represents the behavior of a node in the Follower state.
func (n *Node) runFollower() {
	fmt.Printf("Node %s is running as a Follower.\n", n.ID)
	// Placeholder: Add logic for follower actions
}






