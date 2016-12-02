package framework_test

import (
	. "github.com/onsi/ginkgo"
	"fmt"
)

type GinkgoTestReporter struct {}

func (g GinkgoTestReporter) Errorf(format string, args ...interface{}) {
	Fail(fmt.Sprintf(format, args))
}

func (g GinkgoTestReporter) Fatalf(format string, args ...interface{}) {
	Fail(fmt.Sprintf(format, args))
}
