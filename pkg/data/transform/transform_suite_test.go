package transform_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTransform(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Transform Suite")
}
