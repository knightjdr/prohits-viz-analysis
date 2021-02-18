package settings

import (
	"os"
	"regexp"

	"bou.ke/monkey"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

var _ = Describe("File validation", func() {
	It("should validate file settings and return column map", func() {
		settings := types.Settings{
			Abundance:     "avgSpec",
			Condition:     "bait",
			Control:       "ctrl",
			Files:         []string{"file.txt"},
			Readout:       "prey",
			ReadoutLength: "preyLength",
			Score:         "fdr",
		}

		expected := map[string]string{
			"abundance":     "avgSpec",
			"condition":     "bait",
			"control":       "ctrl",
			"readout":       "prey",
			"readoutLength": "preyLength",
			"score":         "fdr",
		}
		Expect(validateFileSettings(settings)).To(Equal(expected))
	})
})

var _ = Describe("Validate columns", func() {
	It("should validate required file settings", func() {
		// Mock filesystem for logging errors.
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644)

		settings := types.Settings{
			Abundance: "avgSpec",
			Condition: "bait",
			Files:     []string{"file.txt"},
			Readout:   "prey",
			Score:     "fdr",
		}

		// Mock exit.
		fakeExit := func(int) {
			panic("os.Exit called")
		}
		exitPatch := monkey.Patch(os.Exit, fakeExit)
		defer exitPatch.Unpatch()

		Expect(func() { validateColumns(settings) }).To(Not(Panic()))
	})

	It("should exit when missing required file settings", func() {
		// Mock filesystem for logging errors.
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644)

		settings := types.Settings{}

		// Mock exit.
		fakeExit := func(int) {
			panic("os.Exit called")
		}
		exitPatch := monkey.Patch(os.Exit, fakeExit)
		defer exitPatch.Unpatch()

		Expect(func() { validateColumns(settings) }).To((Panic()), "should exit when missing required settings")

		logfile, _ := afero.ReadFile(fs.Instance, "error.txt")
		matched, _ := regexp.MatchString("No abundance column specified", string(logfile))
		Expect(matched).To(BeTrue(), "should write error when missing abundance column")
		matched, _ = regexp.MatchString("No condition column specified", string(logfile))
		Expect(matched).To(BeTrue(), "should write error when missing condition column")
		matched, _ = regexp.MatchString("No input file specified", string(logfile))
		Expect(matched).To(BeTrue(), "should write error when missing input file")
		matched, _ = regexp.MatchString("No readout column specified", string(logfile))
		Expect(matched).To(BeTrue(), "should write error when missing readout column")
		matched, _ = regexp.MatchString("No score column specified", string(logfile))
		Expect(matched).To(BeTrue(), "should write error when missing score column")
	})
})

var _ = Describe("Map columns", func() {
	It("should map columns to fields", func() {
		settings := types.Settings{
			Abundance:     "avgSpec",
			Condition:     "bait",
			Control:       "ctrl",
			Readout:       "prey",
			ReadoutLength: "preyLength",
			Score:         "fdr",
		}

		expected := map[string]string{
			"abundance":     "avgSpec",
			"condition":     "bait",
			"control":       "ctrl",
			"readout":       "prey",
			"readoutLength": "preyLength",
			"score":         "fdr",
		}
		Expect(createColumnMap(settings)).To(Equal(expected))
	})
})
