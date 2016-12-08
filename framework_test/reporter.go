package framework_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
)

type GinkgoTestReporter struct{}


// Fail causes the test to fail with the given formatted string reported as the failure reason
func (g GinkgoTestReporter) Errorf(format string, args ...interface{}) {
	Fail(fmt.Sprintf(format, args))
}

// Fatalf causes the test to fail with the given formatted string reported as the failure reason
func (g GinkgoTestReporter) Fatalf(format string, args ...interface{}) {
	Fail(fmt.Sprintf(format, args))
}
