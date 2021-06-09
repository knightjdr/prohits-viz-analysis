package convert

import (
	"fmt"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
)

func checkFileType(filename string) {
	mtype, err := mimetype.DetectFile(filename)

	if err != nil || !strings.HasPrefix(mtype.String(), "text") {
		log.WriteAndExit(fmt.Sprintf("Could not convert %s", filename))
	}
}
