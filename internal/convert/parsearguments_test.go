package convert

import (
	"os"
	"regexp"

	"github.com/bouk/monkey"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Parse arguments", func() {
	It("should parse command line arguments", func() {
		os.Args = []string{
			"cmd",
			"--file", "file.txt",
		}
		Expect(parseArguments()).To(Equal("file.txt"))
	})

	It("should exit when file is missing", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644)

		fakeExit := func(int) {
			panic("os.Exit called")
		}
		exitPatch := monkey.Patch(os.Exit, fakeExit)
		defer exitPatch.Unpatch()

		os.Args = []string{
			"cmd",
		}

		Expect(func() { parseArguments() }).To(Panic(), "should exit when missing file to convert")

		logfile, _ := afero.ReadFile(fs.Instance, "error.txt")
		matched, _ := regexp.MatchString("no file specified", string(logfile))
		Expect(matched).To(BeTrue(), "should write error when missing file to convert")
	})
})
