package go_testing_tools_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoTestingTools(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoTestingTools Suite")
}
