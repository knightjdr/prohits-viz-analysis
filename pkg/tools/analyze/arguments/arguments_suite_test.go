package arguments_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestArguments(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Arguments Suite")
}
