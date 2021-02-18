package settings

import (
	"os"

	"bou.ke/monkey"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

var _ = Describe("Setting validation", func() {
	It("should validate circheatmap", func() {
		analysis := &types.Analysis{
			Settings: types.Settings{
				Abundance:      "avgSpec",
				Condition:      "bait",
				Control:        "ctrl",
				Files:          []string{"file.txt"},
				OtherAbundance: []string{"column1", "column2"},
				Readout:        "prey",
				ReadoutLength:  "preyLength",
				Score:          "fdr",
				Type:           "circheatmap",
			},
		}

		expected := &types.Analysis{
			Columns: map[string]string{
				"abundance":     "avgSpec",
				"column1":       "column1",
				"column2":       "column2",
				"condition":     "bait",
				"control":       "ctrl",
				"readout":       "prey",
				"readoutLength": "preyLength",
				"score":         "fdr",
			},
			Settings: types.Settings{
				Abundance:      "avgSpec",
				Condition:      "bait",
				Control:        "ctrl",
				Files:          []string{"file.txt"},
				OtherAbundance: []string{"column1", "column2"},
				Readout:        "prey",
				ReadoutLength:  "preyLength",
				Score:          "fdr",
				Type:           "circheatmap",
			},
		}
		Validate(analysis)
		Expect(analysis).To(Equal(expected))
	})

	It("should validate dotplot", func() {
		analysis := types.Analysis{
			Settings: types.Settings{
				Abundance:     "avgSpec",
				Condition:     "bait",
				Control:       "ctrl",
				Files:         []string{"file.txt"},
				Readout:       "prey",
				ReadoutLength: "preyLength",
				Score:         "fdr",
				Type:          "dotplot",
			},
		}

		expected := types.Analysis{
			Columns: map[string]string{
				"abundance":     "avgSpec",
				"condition":     "bait",
				"control":       "ctrl",
				"readout":       "prey",
				"readoutLength": "preyLength",
				"score":         "fdr",
			},
			Settings: types.Settings{
				Abundance:     "avgSpec",
				Condition:     "bait",
				Control:       "ctrl",
				Files:         []string{"file.txt"},
				Readout:       "prey",
				ReadoutLength: "preyLength",
				Score:         "fdr",
				Type:          "dotplot",
			},
		}
		Validate(&analysis)
		Expect(analysis).To(Equal(expected))
	})

	It("should exit when missing required file settings", func() {
		// Mock filesystem for logging errors.
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644)

		analysis := types.Analysis{
			Settings: types.Settings{
				Type: "unknown",
			},
		}

		// Mock exit.
		fakeExit := func(int) {
			panic("os.Exit called")
		}
		exitPatch := monkey.Patch(os.Exit, fakeExit)
		defer exitPatch.Unpatch()

		Expect(func() { Validate(&analysis) }).To((Panic()), "should exit when missing required settings")
	})
})
