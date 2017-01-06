package restapi_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRestapi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Restapi Suite")
}
