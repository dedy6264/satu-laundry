package utils

import (
	"fmt"
	"log"
)

// LoggMsg prints a message with service name and errors
// Parameters:
// - serviceName: name of the service
// - message: the message to log
// - err: error to include (can be nil)
func LoggMsg(serviceName, message string, err error) {
	if err != nil {
		fmt.Println("")
		log.Printf("[%s] %s - Error: %v", serviceName, message, err)
	} else {
		fmt.Println("")
		log.Printf("[%s] %s", serviceName, message)
	}
}
