package float_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFloat(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Float Suite")
}
