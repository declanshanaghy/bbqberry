package backend_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"os"
	"testing"
)

func TestBackend(t *testing.T) {
	RegisterFailHandler(Fail)
	os.Setenv("DB_TIMEOUT_MILLIS", "1")
	RunSpecs(t, "Backend Suite")
}
