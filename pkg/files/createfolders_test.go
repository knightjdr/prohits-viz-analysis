package files_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"

	. "github.com/knightjdr/prohits-viz-analysis/pkg/files"
	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
)

var _ = Describe("Create folders", func() {
	It("should create folders in list", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)

		CreateFolders([]string{"test/folder1", "test/folder2"})

		exists, _ := afero.DirExists(fs.Instance, "test/folder1")
		Expect(exists).To(BeTrue())
		exists, _ = afero.DirExists(fs.Instance, "test/folder2")
		Expect(exists).To(BeTrue())
	})
})
