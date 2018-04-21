// Package logmessage writes a message to a log file (if it exists)
package logmessage

import (
	"log"
	"os"

	"github.com/knightjdr/prohits-viz-analysis/pvfilereader/fs"
)

// Write writes a message to a log file or console if no log specified
func Write(file string, message string) {
	// exit and print to console if no log file specified
	if file == "" {
		log.Println(message)
		os.Exit(1)
	}

	// open log file (create if it doesn't exist)
	f, err := fs.Instance.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}
	defer f.Close()

	// write message
	log.SetOutput(f)
	log.Println(message)
}
