package specificity

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	"github.com/spf13/afero"
)

var _ = Describe("Create folders", func() {
	It("should create default folders", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)

		settings := types.Settings{}

		createFolders(settings)
		exists, _ := afero.DirExists(fs.Instance, "interactive")
		Expect(exists).To(BeTrue(), "should create interactive folder")
		exists, _ = afero.DirExists(fs.Instance, "other")
		Expect(exists).To(BeTrue(), "should create other folder")
		exists, _ = afero.DirExists(fs.Instance, "svg")
		Expect(exists).To(BeTrue(), "should create svg folder")
	})

	It("should create optional folders", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)

		settings := types.Settings{
			Png: true,
		}

		createFolders(settings)
		exists, _ := afero.DirExists(fs.Instance, "png")
		Expect(exists).To(BeTrue(), "should create png folder")
	})
})
