package minimap_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMinimap(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Minimap Suite")
}
