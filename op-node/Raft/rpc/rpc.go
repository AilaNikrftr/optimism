package RPC

import (
	"fmt"
	"sync"
)

// AppendEntriesRequest represents the request structure for the AppendEntries RPC.
type AppendEntriesRequest struct {
	Term         int    // Leader's term
	LeaderID     string // Leader's ID
	PrevLogIndex int    // Index of the log entry immediately preceding the new ones
	PrevLogTerm  int    // Term of the PrevLogIndex entry
	Entries      []LogEntry // Log entries to store (empty for heartbeat)
	LeaderCommit int    // Leader's commit index
}

// AppendEntriesResponse represents the response structure for the AppendEntries RPC.
type AppendEntriesResponse struct {
	Term    int  // Current term, for leader to update itself
	Success bool // True if follower contained entry matching PrevLogIndex and PrevLogTerm
}

// LogEntry represents a single entry in the replicated log.
type LogEntry struct {
	Term    int    // Term when entry was received by leader
	Command string // Command for the state machine
}

// RPCClient represents the client that sends RPC calls to other nodes.
type RPCClient struct {
	NodeID string // The ID of the client node
}

// SendAppendEntries sends an AppendEntries RPC to a target node.
func (c *RPCClient) SendAppendEntries(targetNode string, req *AppendEntriesRequest) (*AppendEntriesResponse, error) {
	fmt.Printf("Client %s sending AppendEntries to node %s: %+v\n", c.NodeID, targetNode, req)

	// Simulate sending the request and receiving a response (e.g., via HTTP or gRPC)
	response := &AppendEntriesResponse{
		Term:    req.Term,
		Success: true, // Assume success for now
	}
	return response, nil
}

// RPCServer represents the server that handles incoming RPC calls.
type RPCServer struct {
	NodeID string       // The ID of the server node
	Log    []LogEntry   // The log entries stored on this server
	Mutex  sync.Mutex   // Mutex to protect shared data
}

// HandleAppendEntries handles an incoming AppendEntries RPC call.
func (s *RPCServer) HandleAppendEntries(req *AppendEntriesRequest) (*AppendEntriesResponse, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	fmt.Printf("Server %s received AppendEntries: %+v\n", s.NodeID, req)

	// Validate the PrevLogIndex and PrevLogTerm
	if req.PrevLogIndex >= 0 && (len(s.Log) <= req.PrevLogIndex || s.Log[req.PrevLogIndex].Term != req.PrevLogTerm) {
		fmt.Printf("Server %s: Log inconsistency detected\n", s.NodeID)
		return &AppendEntriesResponse{
			Term:    s.currentTerm(),
			Success: false,
		}, nil
	}

	// Append new entries to the log
	for i, entry := range req.Entries {
		if req.PrevLogIndex+1+i < len(s.Log) {
			s.Log[req.PrevLogIndex+1+i] = entry
		} else {
			s.Log = append(s.Log, entry)
		}
	}

	// Update commit index if needed
	if req.LeaderCommit > s.commitIndex() {
		commitIndex := min(req.LeaderCommit, len(s.Log)-1)
		fmt.Printf("Server %s: Updating commit index to %d\n", s.NodeID, commitIndex)
	}

	return &AppendEntriesResponse{
		Term:    s.currentTerm(),
		Success: true,
	}, nil
}

// currentTerm returns the current term of the server.
func (s *RPCServer) currentTerm() int {
	// Placeholder for the current term logic
	return 1
}

// commitIndex returns the commit index of the server.
func (s *RPCServer) commitIndex() int {
	// Placeholder for the commit index logic
	return len(s.Log) - 1
}

// Helper function to get the minimum of two integers.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
