package main

import (
    "fmt"
    "os"
    "github.com/user/godistsys/node"
    "github.com/user/godistsys/scheduler"
    "github.com/user/godistsys/communication"
    "github.com/user/godistsys/utils"
)

func main() {
    // Initialize a logger
    logger := utils.NewLogger("GoDistSys")

    // Create a scheduler
    sch := scheduler.NewScheduler()

    // Add some nodes to the scheduler
    for i := 0; i < 5; i++ {
        n := node.NewNode(fmt.Sprintf("node-%d", i))
        sch.AddNode(n)
    }

    // Initialize the communication protocol (assuming it runs on its own goroutine or similar)
    protocol := communication.NewProtocol("localhost:8080")
    go protocol.ReceiveMessage()

    // Schedule some tasks
    for i := 0; i < 10; i++ {
        task := &scheduler.Task{
            ID:      fmt.Sprintf("task-%d", i),
            Payload: "Sample Task Payload",
        }
        sch.ScheduleTask(task)
    }

    // Example of sending a message (to the first node for simplicity)
    if len(sch.Nodes) > 0 {
        err := protocol.SendMessage(sch.Nodes[0].ID, "Hello Node")
        if err != nil {
            logger.Error(err)
        } else {
            logger.Info("Message sent successfully")
        }
    } else {
        logger.Error(fmt.Errorf("no nodes available for messaging"))
    }

    // Process tasks on each node
    for _, n := range sch.Nodes {
        n.ProcessTasks()
    }

    // Exiting the application
    logger.Info("Exiting GoDistSys")
    os.Exit(0)
}
