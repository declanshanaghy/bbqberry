package ws2801_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestWS2801(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "WS2801 Suite")
}
