package dotplot

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	"github.com/spf13/afero"
)

var _ = Describe("Create folders", func() {
	It("create default folders", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)

		settings := types.Settings{}

		createFolders(settings)
		exists, _ := afero.DirExists(fs.Instance, "cytoscape")
		Expect(exists).To(BeTrue(), "should create cytoscape folder")
		exists, _ = afero.DirExists(fs.Instance, "interactive")
		Expect(exists).To(BeTrue(), "should create interactive folder")
		exists, _ = afero.DirExists(fs.Instance, "minimap")
		Expect(exists).To(BeTrue(), "should create minimap folder")
		exists, _ = afero.DirExists(fs.Instance, "other")
		Expect(exists).To(BeTrue(), "should create other folder")
		exists, _ = afero.DirExists(fs.Instance, "svg")
		Expect(exists).To(BeTrue(), "should create svg folder")
		exists, _ = afero.DirExists(fs.Instance, "treeview")
		Expect(exists).To(BeTrue(), "should create treeview folder")
	})

	It("create optional folders", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)

		settings := types.Settings{
			Clustering: "biclustering",
			Png:        true,
		}

		createFolders(settings)
		exists, _ := afero.DirExists(fs.Instance, "biclustering")
		Expect(exists).To(BeTrue(), "should create biclustering folder")
		exists, _ = afero.DirExists(fs.Instance, "png")
		Expect(exists).To(BeTrue(), "should create png folder")
	})
})
