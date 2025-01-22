package Log

import (
    "testing"
)

func TestLog(t *testing.T) {
    log := NewLog()

    // Test AppendEntry
    entry1 := LogEntry{Term: 1, Index: 0, Command: "set x=1"}
    log.AppendEntry(entry1)
    if len(log.entries) != 1 {
        t.Errorf("Expected 1 entry, got %d", len(log.entries))
    }

    // Test GetEntries
    entries, err := log.GetEntries(0)
    if err != nil || len(entries) != 1 || entries[0].Command != "set x=1" {
        t.Errorf("Failed to retrieve log entries correctly")
    }

    // Test CommitEntry
    err = log.CommitEntry(0)
    if err != nil {
        t.Errorf("Failed to commit log entry: %v", err)
    }
    if log.GetCommitIndex() != 0 {
        t.Errorf("Expected commit index 0, got %d", log.GetCommitIndex())
    }

    // Test invalid GetEntries
    _, err = log.GetEntries(5)
    if err == nil {
        t.Errorf("Expected error for invalid index, got nil")
    }
}
