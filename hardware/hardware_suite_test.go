package hardware_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestWS2801(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hardware Suite")
}
