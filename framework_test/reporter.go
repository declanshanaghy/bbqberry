package framework_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
)

type ginkgoTestReporter struct{}


// Fail causes the test to fail with the given formatted string reported as the failure reason
func (g ginkgoTestReporter) Errorf(format string, args ...interface{}) {
	Fail(fmt.Sprintf(format, args))
}

// Fatalf causes the test to fail with the given formatted string reported as the failure reason
func (g ginkgoTestReporter) Fatalf(format string, args ...interface{}) {
	Fail(fmt.Sprintf(format, args))
}
