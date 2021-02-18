package arguments

import (
	"os"
	"regexp"

	"bou.ke/monkey"
	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Parse arguments", func() {
	It("should parse arguments", func() {
		// Mock filesystem.
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "test/settings.json", []byte(jsonText), 0644)

		os.Args = []string{
			"cmd",
			"--settings", "test/settings.json",
		}

		expected := &types.Analysis{
			Columns: map[string]string{
				"abundance":     "avgspec",
				"condition":     "bait",
				"control":       "ctrl",
				"readout":       "prey",
				"readoutLength": "preyLength",
				"score":         "fdr",
			},
			Settings: types.Settings{
				Abundance:          "avgspec",
				AbundanceCap:       50,
				BiclusteringApprox: true,
				Clustering:         "hierarchical",
				Condition:          "bait",
				Control:            "ctrl",
				Files:              []string{"file1.txt", "file2.txt"},
				Readout:            "prey",
				ReadoutLength:      "preyLength",
				Score:              "fdr",
				Type:               "dotplot",
			},
		}
		Expect(Parse()).To(Equal(expected))
	})
})

var _ = Describe("Read arguments", func() {
	It("should read arguments", func() {
		// Mock filesystem.
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "test/settings.json", []byte(jsonText), 0644)

		os.Args = []string{
			"cmd",
			"--settings", "test/settings.json",
		}

		expected := &types.Analysis{
			Settings: types.Settings{
				Abundance:          "avgspec",
				AbundanceCap:       50,
				BiclusteringApprox: true,
				Clustering:         "hierarchical",
				Condition:          "bait",
				Control:            "ctrl",
				Files:              []string{"file1.txt", "file2.txt"},
				Readout:            "prey",
				ReadoutLength:      "preyLength",
				Score:              "fdr",
				Type:               "dotplot",
			},
		}
		Expect(readArguments()).To(Equal(expected))
	})

	It("should exit when missing required arguments", func() {
		// Mock filesystem.
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644)

		os.Args = []string{
			"cmd",
		}

		// Mock exit.
		fakeExit := func(int) {
			panic("os.Exit called")
		}
		exitPatch := monkey.Patch(os.Exit, fakeExit)
		defer exitPatch.Unpatch()

		Expect(func() { readArguments() }).To((Panic()), "should exit when missing required arguments")

		logfile, _ := afero.ReadFile(fs.Instance, "error.txt")
		matched, _ := regexp.MatchString("no settings file specified", string(logfile))
		Expect(matched).To(BeTrue(), "should write error when missing settings file")
	})
})
