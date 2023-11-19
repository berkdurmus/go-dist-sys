package utils

import (
    "fmt"
    "log"
    "time"
)

// Logger for the application
type Logger struct {
    Prefix string
}

// NewLogger creates a new Logger instance
func NewLogger(prefix string) *Logger {
    return &Logger{Prefix: prefix}
}

// Info logs informational messages
func (l *Logger) Info(message string) {
    log.Printf("%s [INFO]: %s\n", l.Prefix, message)
}

// Error logs error messages
func (l *Logger) Error(err error) {
    log.Printf("%s [ERROR]: %s\n", l.Prefix, err.Error())
}

// TimeTrack is a utility function to track the duration of a function call
// Usage: defer TimeTrack(time.Now(), "functionName")
func TimeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    log.Printf("%s took %s", name, elapsed)
}

// Other utility functions can be added as needed
