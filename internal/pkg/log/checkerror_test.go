package log_test

import (
	"errors"
	"os"
	"regexp"

	"github.com/bouk/monkey"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/log"
	"github.com/spf13/afero"
)

var _ = Describe("Check for error to log", func() {
	It("should log message to file", func() {
		// Mock filesystem.
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644)

		err := errors.New("test message")
		log.CheckError(err, false)

		logfile, _ := afero.ReadFile(fs.Instance, "error.txt")
		expected := "test message"
		matched, _ := regexp.MatchString(expected, string(logfile))
		Expect(matched).To(BeTrue())
	})

	It("should exit", func() {
		// Mock filesystem.
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644)

		// Mock exit.
		fakeExit := func(int) {
			panic("os.Exit called")
		}
		exitPatch := monkey.Patch(os.Exit, fakeExit)
		defer exitPatch.Unpatch()

		err := errors.New("test message")
		Expect(func() { log.CheckError(err, true) }).To(Panic())
	})
})
