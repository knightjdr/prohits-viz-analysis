// Package settings logs analysis settings.
package settings

import (
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	"github.com/spf13/afero"
)

// Log analysis settings.
func Log(settings types.Settings) {
	var messages strings.Builder

	logSharedSettings(&messages, settings)

	switch settings.Type {
	case "correlation":
		logCorrelationSettings(&messages, settings)
	case "dotplot":
		logDotplotSettings(&messages, settings)
	}

	afero.WriteFile(fs.Instance, "log.txt", []byte(messages.String()), 0644)
}
