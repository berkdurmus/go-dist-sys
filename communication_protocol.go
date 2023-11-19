package communication

import (
    "fmt"
    "net"
)

// Protocol struct represents the communication protocol for node-to-node interactions
type Protocol struct {
    // Add relevant properties here, like connection details
    Address string // Address of the node
}

// NewProtocol creates a new Protocol instance
func NewProtocol(address string) *Protocol {
    return &Protocol{
        Address: address,
    }
}

// SendMessage sends a message to a specified address
func (p *Protocol) SendMessage(addr string, message string) error {
    conn, err := net.Dial("tcp", addr)
    if err != nil {
        return fmt.Errorf("error connecting to node: %v", err)
    }
    defer conn.Close()

    _, err = conn.Write([]byte(message))
    if err != nil {
        return fmt.Errorf("error sending message: %v", err)
    }

    return nil
}

// ReceiveMessage listens for incoming messages
func (p *Protocol) ReceiveMessage() {
    ln, err := net.Listen("tcp", p.Address)
    if err != nil {
        fmt.Printf("error setting up listener: %v\n", err)
        return
    }

    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Printf("error accepting connection: %v\n", err)
            continue
        }

        go p.handleConnection(conn)
    }
}

// handleConnection handles an individual connection from a node
func (p *Protocol) handleConnection(conn net.Conn) {
    defer conn.Close()

    buffer := make([]byte, 1024)
    len, err := conn.Read(buffer)
    if err != nil {
        fmt.Printf("error reading message: %v\n", err)
        return
    }

    fmt.Printf("Received message: %s\n", string(buffer[:len]))
}

