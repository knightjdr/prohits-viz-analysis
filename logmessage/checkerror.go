package logmessage

// CheckError checks if there is an error, logs it and panics if requested.
func CheckError(err error, shouldPanic bool) {
	if err != nil {
		message := err.Error()
		Write(message)
		if shouldPanic {
			panic(message)
		}
	}
}
