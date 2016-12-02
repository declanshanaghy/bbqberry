package backend_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	
	"testing"
)

func TestTemperatures(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "WS2801 Suite")
}

