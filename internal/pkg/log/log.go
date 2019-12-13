// Package log writes a message to a log file or console.
package log

import (
	goLog "log"
	"os"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
)

// Write writes a message to a log file or console if no log specified.
func Write(message string) {
	// Open log file (create if it doesn't exist).
	f, err := fs.Instance.OpenFile("error.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	CheckError(err, true)
	defer f.Close()

	goLog.SetOutput(f)
	goLog.Println(message)
}
