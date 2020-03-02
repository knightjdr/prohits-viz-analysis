package export

import (
	"os"
	"regexp"

	"github.com/bouk/monkey"
	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Parse arguments", func() {
	It("should parse command line arguments", func() {
		os.Args = []string{
			"cmd",
			"--file", "file.json",
			"--font", "./font.ttf",
			"--format", "png",
			"--imageType", "heatmap",
		}

		expected := parameters{
			fontPath:  "./font.ttf",
			format:    "png",
			imageType: "heatmap",
			jsonFile:  "file.json",
		}
		Expect(parseArguments()).To(Equal(expected))
	})

	It("should set defaults when arguments not supplied", func() {
		os.Args = []string{
			"cmd",
			"--file", "file.json",
			"--imageType", "heatmap",
		}
		actual := parseArguments()
		Expect(actual.fontPath).To(Equal(""), "should set font path")
		Expect(actual.format).To(Equal("svg"), "should set output format")
	})

	It("should exit when JSON file is missing", func() {
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
			"--format", "png",
			"--imageType", "heatmap",
		}

		Expect(func() { parseArguments() }).To(Panic(), "should exit when missing JSON file to export")

		logfile, _ := afero.ReadFile(fs.Instance, "error.txt")
		matched, _ := regexp.MatchString("no JSON file specified", string(logfile))
		Expect(matched).To(BeTrue(), "should write error when missing file to export")
	})

	It("should exit when image type is missing", func() {
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
			"--file", "file.json",
			"--format", "png",
		}

		Expect(func() { parseArguments() }).To(Panic(), "should exit when missing image type")

		logfile, _ := afero.ReadFile(fs.Instance, "error.txt")
		matched, _ := regexp.MatchString("image type not specified", string(logfile))
		Expect(matched).To(BeTrue(), "should write error when missing image type")
	})

	It("should exit when output format is invalid", func() {
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
			"--file", "file.json",
			"--format", "pdf",
			"--imageType", "heatmap",
		}

		Expect(func() { parseArguments() }).To(Panic(), "should exit when the output format is invalid")

		logfile, _ := afero.ReadFile(fs.Instance, "error.txt")
		matched, _ := regexp.MatchString("invalid output format", string(logfile))
		Expect(matched).To(BeTrue(), "should write error when output format is invalid")
	})
})
