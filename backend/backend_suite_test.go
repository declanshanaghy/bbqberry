package backend_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"os"
	"testing"
	"github.com/Polarishq/middleware/framework/log"
)

func TestBackend(t *testing.T) {
	log.SetDebug(true)

	RegisterFailHandler(Fail)
	os.Setenv("DB_TIMEOUT_MILLIS", "1")
	RunSpecs(t, "Backend Suite")
}
