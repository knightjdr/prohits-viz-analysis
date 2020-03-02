package log

import "os"

// WriteAndExit logs a message and exits.
func WriteAndExit(message string) {
	Write(message)
	os.Exit(1)
}
