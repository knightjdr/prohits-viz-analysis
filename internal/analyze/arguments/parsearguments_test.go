package arguments

import (
	"os"
	"regexp"

	"github.com/bouk/monkey"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
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
		afero.WriteFile(fs.Instance, "test/settings.json", []byte(settingsText), 0644)

		os.Args = []string{
			"cmd",
			"-analysisType=dotplot",
			"--settings", "test/settings.json",
		}

		expected := types.Analysis{
			Columns: map[string]string{
				"abundance":     "avgspec",
				"condition":     "bait",
				"control":       "ctrl",
				"readout":       "prey",
				"readoutLength": "preyLength",
				"score":         "fdr",
			},
			Settings: &types.Dotplot{
				File: types.File{
					Abundance:     "avgspec",
					Condition:     "bait",
					Control:       "ctrl",
					Files:         []string{"file1.txt", "file2.txt"},
					Readout:       "prey",
					ReadoutLength: "preyLength",
					Score:         "fdr",
				},
				AbundanceCap:       50,
				BiclusteringApprox: true,
				Clustering:         "hierarchical",
			},
			Type: "dotplot",
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
		afero.WriteFile(fs.Instance, "test/settings.json", []byte(settingsText), 0644)

		os.Args = []string{
			"cmd",
			"-analysisType=dotplot",
			"--settings", "test/settings.json",
		}

		expectedSettings := &types.Dotplot{
			File: types.File{
				Abundance:     "avgspec",
				Condition:     "bait",
				Control:       "ctrl",
				Files:         []string{"file1.txt", "file2.txt"},
				Readout:       "prey",
				ReadoutLength: "preyLength",
				Score:         "fdr",
			},
			AbundanceCap:       50,
			BiclusteringApprox: true,
			Clustering:         "hierarchical",
		}
		actualType, actualSettings := readArguments()
		Expect(actualType).To(Equal("dotplot"))
		Expect(actualSettings).To(Equal(expectedSettings))
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
		matched, _ := regexp.MatchString("No analysis type specified", string(logfile))
		Expect(matched).To(BeTrue(), "should write error when missing analysis type")
		matched, _ = regexp.MatchString("No settings file specified", string(logfile))
		Expect(matched).To(BeTrue(), "should write error when missing settings file")
	})
})
