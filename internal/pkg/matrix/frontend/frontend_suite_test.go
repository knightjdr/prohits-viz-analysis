package frontend_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFrontend(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Frontend Suite")
}
