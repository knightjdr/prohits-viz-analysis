package dotplot_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDotplot(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dotplot Suite")
}
