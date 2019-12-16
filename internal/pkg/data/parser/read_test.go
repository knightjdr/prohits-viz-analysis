package parser

import (
	"os"

	"github.com/bouk/monkey"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	"github.com/spf13/afero"
)

var _ = Describe("Read file", func() {
	It("should parse file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/testfile.txt",
			[]byte("column1\tcolumn2\tcolumn3\na\tb\tc\nd\te\tf\n"),
			0444,
		)

		analysis := &types.Analysis{
			Columns: map[string]string{
				"condition": "column1",
				"readout":   "column3",
			},
			Settings: types.Settings{
				Files: []string{"test/testfile.txt"},
			},
		}

		expected := &types.Analysis{
			Columns: map[string]string{
				"condition": "column1",
				"readout":   "column3",
			},
			Data: []map[string]string{
				{"condition": "a", "readout": "c"},
				{"condition": "d", "readout": "f"},
			},
			Settings: types.Settings{
				Files: []string{"test/testfile.txt"},
			},
		}
		Read(analysis, false)
		Expect(analysis).To(Equal(expected))
	})

	It("should exit when there are no parsed results", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/empty.txt",
			[]byte("column1,column2,column3\n"),
			0444,
		)

		// Mock exit.
		fakeExit := func(int) {
			panic("os.Exit called")
		}
		exitPatch := monkey.Patch(os.Exit, fakeExit)
		defer exitPatch.Unpatch()

		analysis := &types.Analysis{
			Columns: map[string]string{
				"condition": "column1",
				"readout":   "column3",
			},
			Settings: types.Settings{
				Files: []string{"test/empty.txt"},
			},
		}

		Expect(func() { Read(analysis, false) }).To(Panic())
	})
})
