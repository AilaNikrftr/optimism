package Log

import "errors"
//
// The LogEntry represents a single command or data entry in the Raft log.
//

// LogEntry represents a single entry in the Raft log.
type LogEntry struct {
    Term    int    // The term when the entry was received by the leader
    Index   int    // The log index of the entry
    Command string // The actual command/data to be executed
}

// Log represents the log storage for Raft.
type Log struct {
    entries       []LogEntry // List of log entries
    commitIndex   int        // Index of the highest committed entry
    lastApplied   int        // Index of the last applied entry
    nextEntryIndex int       // Index to be used for the next log entry
}

// NewLog initializes a new log.
func NewLog() *Log {
    return &Log{
        entries:       make([]LogEntry, 0),
        commitIndex:   -1,
        lastApplied:   -1,
        nextEntryIndex: 0,
    }
}

// AppendEntry appends a new log entry.
func (l *Log) AppendEntry(entry LogEntry) {
    l.entries = append(l.entries, entry)
    l.nextEntryIndex++
}

// GetEntries retrieves log entries starting from a specific index.
func (l *Log) GetEntries(startIndex int) ([]LogEntry, error) {
    if startIndex < 0 || startIndex >= len(l.entries) {
        return nil, errors.New("invalid start index")
    }
    return l.entries[startIndex:], nil
}

// CommitEntry marks entries as committed up to a specific index.
func (l *Log) CommitEntry(index int) error {
    if index < 0 || index >= l.nextEntryIndex {
        return errors.New("invalid index")
    }
    l.commitIndex = index
    return nil
}

// GetCommitIndex returns the current commit index.
func (l *Log) GetCommitIndex() int {
    return l.commitIndex
}

func (l *Log) AppendEntry(entry LogEntry) {
    l.entries = append(l.entries, entry)
    l.nextEntryIndex++
    fmt.Printf("Appended log entry: %+v\n", entry)
}
