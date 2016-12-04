package influx_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestInflux(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Influx Suite")
}
