package scheduler

import (
    "github.com/user/godistsys/node"
    "sync"
)

// Task represents a basic task structure
// Assuming this is defined in a common package or duplicated here for simplicity
type Task struct {
    ID      string
    Payload string
}

// Scheduler manages and distributes tasks to nodes
type Scheduler struct {
    Nodes []*node.Node
    mu    sync.Mutex // Mutex to handle concurrent access to the Nodes slice
}

// NewScheduler creates a new Scheduler instance
func NewScheduler() *Scheduler {
    return &Scheduler{
        Nodes: make([]*node.Node, 0),
    }
}

// AddNode adds a new node to the scheduler
func (s *Scheduler) AddNode(n *node.Node) {
    s.mu.Lock()
    defer s.mu.Unlock()

    s.Nodes = append(s.Nodes, n)
}

// RemoveNode removes a node from the scheduler
func (s *Scheduler) RemoveNode(nodeID string) {
    s.mu.Lock()
    defer s.mu.Unlock()

    for i, n := range s.Nodes {
        if n.ID == nodeID {
            s.Nodes = append(s.Nodes[:i], s.Nodes[i+1:]...)
            break
        }
    }
}

// ScheduleTask schedules a task to an appropriate node
func (s *Scheduler) ScheduleTask(task *Task) {
    s.mu.Lock()
    defer s.mu.Unlock()

    // Select a node for the task (simple round-robin for this example)
    if len(s.Nodes) == 0 {
        // Handle the case where there are no nodes
        return
    }

    node := s.Nodes[0] // This is a simplistic approach
    s.Nodes = append(s.Nodes[1:], node) // Move the used node to the end of the slice

    node.AddTask(task) // Assuming AddTask is a method of the node.Node struct
}

// Other methods can be added as needed, such as updating node status, handling failed tasks, etc.
// HandleFailedTask handles a task that has failed on a given node
func (s *Scheduler) HandleFailedTask(taskID string, failedNodeID string) {
    s.mu.Lock()
    defer s.mu.Unlock()

    // Find the task and the node on which it failed
    var failedTask *Task
    var failedTaskIndex int
    var failedNode *node.Node
    for _, n := range s.Nodes {
        if n.ID == failedNodeID {
            failedNode = n
            for i, t := range n.Tasks {
                if t.ID == taskID {
                    failedTask = t
                    failedTaskIndex = i
                    break
                }
            }
            break
        }
    }

    if failedTask == nil {
        // Task or Node not found
        return
    }

    // Remove the failed task from the failed node
    failedNode.Tasks = append(failedNode.Tasks[:failedTaskIndex], failedNode.Tasks[failedTaskIndex+1:]...)

    // Optionally, re-schedule the failed task to another node
    // This can be as simple as calling ScheduleTask again,
    // or involve more complex logic based on the failure reason, task type, etc.
    s.ScheduleTask(failedTask)
}
