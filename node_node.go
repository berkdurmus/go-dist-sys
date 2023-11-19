package node

import (
    "sync"
)

// NodeStatus represents the status of a node
type NodeStatus int

const (
    Idle NodeStatus = iota
    Busy
)

// Task represents a basic task structure
type Task struct {
    ID      string
    Payload string // Payload can be any data associated with the task
}

// Node represents a node in the distributed system
type Node struct {
    ID     string
    Status NodeStatus
    Queue  []*Task
    mu     sync.Mutex // Mutex to handle concurrent access to the Queue
}

// NewNode creates a new Node instance
func NewNode(id string) *Node {
    return &Node{
        ID:     id,
        Status: Idle,
        Queue:  make([]*Task, 0),
    }
}

// AddTask adds a new task to the node's queue
func (n *Node) AddTask(task *Task) {
    n.mu.Lock()
    defer n.mu.Unlock()

    n.Queue = append(n.Queue, task)
    n.updateStatus()
}

// ProcessTasks processes all tasks in the queue
func (n *Node) ProcessTasks() {
    for _, task := range n.Queue {
        n.processTask(task)
    }
    n.clearQueue()
}

// processTask processes an individual task (implementation can vary)
func (n *Node) processTask(task *Task) {
    // Process the task (the actual implementation depends on the task type and payload)
    // For example, it could involve computation, data storage, etc.
}

// clearQueue clears the task queue and updates the node's status
func (n *Node) clearQueue() {
    n.mu.Lock()
    defer n.mu.Unlock()

    n.Queue = n.Queue[:0] // Clear the queue
    n.updateStatus()
}

// updateStatus updates the node's status based on the queue length
func (n *Node) updateStatus() {
    if len(n.Queue) == 0 {
        n.Status = Idle
    } else {
        n.Status = Busy
    }
}
