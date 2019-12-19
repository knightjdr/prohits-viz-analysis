package downsample_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDownsample(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Downsample Suite")
}
