package treeview_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTreeview(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Treeview Suite")
}
