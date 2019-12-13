package log

import "os"

// CheckError checks if there is an error, logs it and exits if requested.
func CheckError(err error, shouldExit bool) {
	if err != nil {
		message := err.Error()
		Write(message)
		if shouldExit {
			os.Exit(1)
		}
	}
}
