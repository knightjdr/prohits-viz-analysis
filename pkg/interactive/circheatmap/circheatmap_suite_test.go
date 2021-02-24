package circheatmap_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCircheatmap(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Circheatmap Suite")
}
