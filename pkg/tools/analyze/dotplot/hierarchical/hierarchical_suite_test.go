package hierarchical_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHierarchical(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hierarchical Suite")
}
